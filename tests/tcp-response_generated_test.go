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

	"github.com/haproxytech/config-parser/v5/parsers/tcp"
)

func TestResponsestcp(t *testing.T) {
	tests := map[string]bool{
		"tcp-response content lua.foo":                                          true,
		"tcp-response content lua.foo param if !HTTP":                           true,
		"tcp-response content lua.foo param param1":                             true,
		"tcp-response content set-dst dest":                                     true,
		"tcp-response content unset-var(sess.my_var)":                           true,
		"tcp-response content set-bandwidth-limit my-limit":                     true,
		"tcp-response content set-bandwidth-limit my-limit limit 1m period 10s": true,
		"tcp-response content set-bandwidth-limit my-limit period 10s":          true,
		"tcp-response content set-bandwidth-limit my-limit limit 1m":            true,
		"tcp-response content set-log-level silent":                             true,
		"tcp-response content set-log-level silent if FALSE":                    true,
		"tcp-response content set-mark 20":                                      true,
		"tcp-response content set-mark 0x1Ab if FALSE":                          true,
		"tcp-response content set-tos 0 if FALSE":                               true,
		"tcp-response content set-tos 0":                                        true,
		"tcp-response content set-nice 0 if FALSE":                              true,
		"tcp-response content set-nice 0":                                       true,
		"tcp-response content close":                                            true,
		"tcp-response content close if !HTTP":                                   true,
		"tcp-response content sc-inc-gpc(1,2)":                                  true,
		"tcp-response content sc-inc-gpc(1,2) if is-error":                      true,
		"tcp-response content sc-inc-gpc0(2)":                                   true,
		"tcp-response content sc-inc-gpc0(2) if is-error":                       true,
		"tcp-response content sc-inc-gpc1(2)":                                   true,
		"tcp-response content sc-inc-gpc1(2) if is-error":                       true,
		"tcp-response":                                                       false,
		"tcp-response content lua.":                                          false,
		"tcp-response content lua. param":                                    false,
		"tcp-response content set-priority-class":                            false,
		"tcp-response content do-resolve":                                    false,
		"tcp-response content set-priority-offset":                           false,
		"tcp-response content set-dst":                                       false,
		"tcp-response content capture":                                       false,
		"tcp-response content set-bandwidth-limit my-limit limit":            false,
		"tcp-response content set-bandwidth-limit my-limit period":           false,
		"tcp-response content set-bandwidth-limit my-limit 10s":              false,
		"tcp-response content set-bandwidth-limit my-limit period 10s limit": false,
		"tcp-response content set-bandwidth-limit my-limit limit period 10s": false,
		"tcp-response content set-log-level":                                 false,
		"tcp-response content set-mark":                                      false,
		"tcp-response content set-tos":                                       false,
		"tcp-response content set-nice":                                      false,
		"tcp-response content sc-inc-gpc":                                    false,
		"---":                                                                false,
		"--- ---":                                                            false,
	}
	parser := &tcp.Responses{}
	for command, shouldPass := range tests {
		t.Run(command, func(t *testing.T) {
			line := strings.TrimSpace(command)
			lines := strings.SplitN(line, "\n", -1)
			var err error
			parser.Init()
			if len(lines) > 1 {
				for _, line = range lines {
					line = strings.TrimSpace(line)
					if err = ProcessLine(line, parser); err != nil {
						break
					}
				}
			} else {
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
