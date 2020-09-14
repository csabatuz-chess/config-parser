// Code generated by go generate; DO NOT EDIT.
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

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/v3/parsers"
)

func TestBalance(t *testing.T) {
	tests := map[string]bool{
		"balance roundrobin": true,
		"balance uri": true,
		"balance uri whole": true,
		"balance uri len 12": true,
		"balance uri depth 8": true,
		"balance uri depth 8 whole": true,
		"balance uri depth 8 len 12 whole": true,
		"balance url_param": true,
		"balance url_param session_id": true,
		"balance url_param check_post 10": true,
		"balance url_param check_post 10 max_wait 20": true,
		"balance url_param session_id check_post 10 max_wait 20": true,
		"balance hdr(hdrName)": true,
		"balance hdr(hdrName) use_domain_only": true,
		"balance random": true,
		"balance random(15)": true,
		"balance rdp-cookie": true,
		"balance rdp-cookie(something)": true,
		"balance something": false,
		"balance": false,
		"balance uri len notInteger": false,
		"balance uri depth notInteger": false,
		"balance url_param check_post notInteger": false,
		"---": false,
		"--- ---": false,
	}
	parser := &parsers.Balance{}
	for command, shouldPass := range tests {
		t.Run(command, func(t *testing.T) {
		line :=strings.TrimSpace(command)
		lines := strings.SplitN(line,"\n", -1)
		var err error
		parser.Init()
		if len(lines)> 1{
			for _,line = range(lines){
			  line = strings.TrimSpace(line)
				if err=ProcessLine(line, parser);err!=nil{
					break
				}
			}
		}else{
			err = ProcessLine(line, parser)
		}
			if shouldPass {
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				result, err := parser.Result()
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				var returnLine string
				if result[0].Comment == "" {
					returnLine = result[0].Data
				} else {
					returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
				}
				if command != returnLine {
					t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, command))
				}
			} else {
				if err == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
				}
				_, parseErr := parser.Result()
				if parseErr == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
				}
			}
		})
	}
}
