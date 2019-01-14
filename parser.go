package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/haproxytech/config-parser/helpers"
	"github.com/haproxytech/config-parser/parsers/extra"
)

//Parser reads and writes configuration on given file
type Parser struct {
	Comments  ParserTypes
	Default   ParserTypes
	Global    ParserTypes
	Frontends map[string]ParserTypes
	Backends  map[string]ParserTypes
	Listen    map[string]ParserTypes
	Resolvers map[string]ParserTypes
	UserLists map[string]ParserTypes
}

func (p *Parser) get(data map[string]ParserTypes, key string, attribute string) (ParserType, error) {
	for _, parser := range data[key].parsers {
		if parser.GetParserName() == attribute && parser.Valid() {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("attribute not found")
}

//GetDefaultsAttr get attribute from defaults section
func (p *Parser) GetDefaultsAttr(attribute string) (ParserType, error) {
	return p.Default.Get(attribute)
}

//GetGlobalAttr get attribute from global section
func (p *Parser) GetGlobalAttr(attribute string) (ParserType, error) {
	return p.Global.Get(attribute)
}

//NewGlobalAttr adds attribute to global section, if exists, replaces it
func (p *Parser) NewGlobalAttr(parser ParserType) {
	p.Global.Set(parser)
}

//GetUserlistAttr get attribute from listen sections
func (p *Parser) GetUserlistAttr(section string, attribute string) (ParserType, error) {
	return p.get(p.UserLists, section, attribute)
}

//GetFrontendAttr get attribute from frontend sections
func (p *Parser) GetFrontendAttr(frontendName string, attribute string) (ParserType, error) {
	return p.get(p.Frontends, frontendName, attribute)
}

//GetBackendAttr get attribute from backend sections
func (p *Parser) GetBackendAttr(backendName string, attribute string) (ParserType, error) {
	return p.get(p.Backends, backendName, attribute)
}

//GetListenAttr get attribute from listen sections
func (p *Parser) GetListenAttr(section string, attribute string) (ParserType, error) {
	return p.get(p.Listen, section, attribute)
}
func (p *Parser) writeParsers(parsers []ParserType, result *strings.Builder) {
	for _, parser := range parsers {
		if !parser.Valid() {
			continue
		}
		for _, line := range parser.Result(true) {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}
}

//String returns configuration in writable form
func (p *Parser) String() string {
	var result strings.Builder

	p.writeParsers(p.Comments.parsers, &result)

	result.WriteString("\ndefaults ")
	result.WriteString("\n")
	p.writeParsers(p.Default.parsers, &result)

	result.WriteString("\nglobal ")
	result.WriteString("\n")
	p.writeParsers(p.Global.parsers, &result)

	for sectionName, section := range p.UserLists {
		result.WriteString("\nuserlist ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for sectionName, section := range p.Resolvers {
		result.WriteString("\nresolvers ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for sectionName, section := range p.Frontends {
		result.WriteString("\nfrontend ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}

	for sectionName, section := range p.Backends {
		result.WriteString("\nbackend ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}
	for sectionName, section := range p.Listen {
		result.WriteString("\nlisten ")
		result.WriteString(sectionName)
		result.WriteString("\n")
		p.writeParsers(section.parsers, &result)
	}
	return result.String()
}

func (p *Parser) Save(filename string) error {
	d1 := []byte(p.String())
	err := ioutil.WriteFile(filename, d1, 0644)
	if err != nil {
		return err
	}
	return nil
}

//ProcessLine parses line plus determines if we need to change state
func (p *Parser) ProcessLine(state string, activeParser ParserTypes, line string, parts, previousParts []string, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist ParserTypes) (newState string, newParserFrontend, newParserBackend, newParserListen, newParserResolver, newParserUserlist ParserTypes) {
	for _, parser := range activeParser.parsers {
		if newState, err := parser.Parse(line, parts, previousParts); err == nil {
			//should we have an option to remove it when found?
			if newState != "" {
				//log.Printf("change state from %s to %s\n", state, newState)
				state = newState
				if state == "frontend" {
					sectionName := parser.(*extra.SectionName)
					parserFrontend = getFrontendParser()
					p.Frontends[sectionName.SectionName] = parserFrontend
				}
				if state == "backend" {
					sectionName := parser.(*extra.SectionName)
					parserBackend = getBackendParser()
					p.Backends[sectionName.SectionName] = parserBackend
				}
				if state == "listen" {
					sectionName := parser.(*extra.SectionName)
					parserListen = getListenParser()
					p.Listen[sectionName.SectionName] = parserListen
				}
				if state == "resolvers" {
					sectionName := parser.(*extra.SectionName)
					parserResolver = getResolverParser()
					p.Resolvers[sectionName.SectionName] = parserResolver
				}
				if state == "userlist" {
					sectionName := parser.(*extra.SectionName)
					parserUserlist = getUserlistParser()
					p.UserLists[sectionName.SectionName] = parserUserlist
				}
			}
			break
		}
	}
	return state, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist
}

func (p *Parser) LoadData(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	p.Comments = getStartParser()
	p.Default = getDefaultParser()
	p.Global = getGlobalParser()

	p.Frontends = map[string]ParserTypes{}
	p.Backends = map[string]ParserTypes{}
	p.Listen = map[string]ParserTypes{}
	p.Resolvers = map[string]ParserTypes{}
	p.UserLists = map[string]ParserTypes{}

	var parserFrontend ParserTypes
	var parserBackend ParserTypes
	var parserListen ParserTypes
	var parserResolver ParserTypes
	var parserUserlist ParserTypes

	lines := helpers.StringSplitIgnoreEmpty(string(dat), '\n')
	state := ""
	previousLine := []string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := helpers.StringSplitIgnoreEmpty(line, ' ')
		if len(parts) == 0 {
			continue
		}
		switch state {
		case "":
			state, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist =
				p.ProcessLine(state, p.Comments, line, parts, previousLine, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist)
			previousLine = parts
		case "defaults":
			state, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist =
				p.ProcessLine(state, p.Default, line, parts, previousLine, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist)
			previousLine = parts
		case "global":
			state, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist =
				p.ProcessLine(state, p.Global, line, parts, previousLine, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist)
			previousLine = parts
		case "frontend":
			state, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist =
				p.ProcessLine(state, parserFrontend, line, parts, previousLine, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist)
			previousLine = parts
		case "backend":
			state, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist =
				p.ProcessLine(state, parserBackend, line, parts, previousLine, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist)
			previousLine = parts
		case "listen":
			state, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist =
				p.ProcessLine(state, parserListen, line, parts, previousLine, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist)
			previousLine = parts
		case "resolvers":
			state, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist =
				p.ProcessLine(state, parserResolver, line, parts, previousLine, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist)
			previousLine = parts
		case "userlist":
			state, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist =
				p.ProcessLine(state, parserUserlist, line, parts, previousLine, parserFrontend, parserBackend, parserListen, parserResolver, parserUserlist)
			previousLine = parts
		}
	}
	return nil
}
