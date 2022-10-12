/*
Copyright 2022 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"

	"github.com/google/renameio"
)

const license = `// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
`

type TemplateType uint8

const (
	TemplateTypeNormal TemplateType = 1
	TemplateTypeOther  TemplateType = 2
	TemplateTypeTest   TemplateType = 10
)

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		log.Println("File " + filePath + " already exists")
	}
	return !os.IsNotExist(err)
}

func cleanFileName(filename string) string {
	return strings.ReplaceAll(filename, " ", "-")
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//go:embed normal.tmpl
var templateNormal []byte

//go:embed others.tmpl
var templateOthers []byte

//go:embed test.tmpl
var templateTest []byte

func executeTemplate(tmplType TemplateType, data *Data, filePath string) {
	var tmpl *template.Template
	switch tmplType {
	case TemplateTypeNormal:
		tmpl = template.Must(template.New("").Parse(license + "\n" + string(templateNormal)))
	case TemplateTypeOther:
		tmpl = template.Must(template.New("").Parse(license + "\n" + string(templateOthers)))
	case TemplateTypeTest:
		tmpl = template.Must(template.New("").Parse(license + "\n" + string(templateTest)))
	}

	log.Println(filePath)
	var tpl bytes.Buffer
	err := tmpl.Execute(&tpl, data)
	CheckErr(err)
	res, errFmt := GoFmt(tpl.Bytes())
	if errFmt != nil {
		res = tpl.Bytes()
	}
	err = renameio.WriteFile(filePath, res, 0o666)
	CheckErr(errFmt)
	CheckErr(err)
}

func GoFmt(input []byte) ([]byte, error) {
	res, err := format.Source(input)
	if err != nil {
		return nil, fmt.Errorf("cmd.Run() failed with %w", err)
	}
	return res, nil
}

func getNiceName(str string) string {
	var res strings.Builder
	for _, c := range str {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			res.WriteRune(c)
		}
	}
	strRes := res.String()
	size := len(strRes)
	if size > 32 {
		size = 32
	}
	strRes = strRes[0:size]
	return strRes
}
