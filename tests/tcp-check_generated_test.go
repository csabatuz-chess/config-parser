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

	"github.com/haproxytech/config-parser/v4/parsers/tcp"
)

func TestCheckstcp(t *testing.T) {
	tests := map[string]bool{
		"tcp-check comment testcomment":                                     true,
		"tcp-check connect":                                                 true,
		"tcp-check connect port 443 ssl":                                    true,
		"tcp-check connect port 110 linger":                                 true,
		"tcp-check connect port 143":                                        true,
		"tcp-check expect string +PONG":                                     true,
		"tcp-check expect string role:master":                               true,
		"tcp-check expect string +OK":                                       true,
		"tcp-check send-lf testfmt":                                         true,
		"tcp-check send-lf testfmt comment testcomment":                     true,
		"tcp-check send-binary testhexstring":                               true,
		"tcp-check send-binary testhexstring comment testcomment":           true,
		"tcp-check send-binary-lf testhexfmt":                               true,
		"tcp-check send-binary-lf testhexfmt comment testcomment":           true,
		"tcp-check set-var(check.port) int(1234)":                           true,
		`tcp-check expect string +OK\ POP3\ ready`:                          true,
		`tcp-check expect string *\ OK\ IMAP4\ ready`:                       true,
		`tcp-check send PING\r\n`:                                           true,
		`tcp-check send PING\r\n comment testcomment`:                       true,
		`tcp-check send QUIT\r\n`:                                           true,
		`tcp-check send QUIT\r\n comment testcomment`:                       true,
		`tcp-check send info\ replication\r\n`:                              true,
		`tcp-check set-var-fmt(check.name) "%H"`:                            true,
		`tcp-check set-var-fmt(txn.from) "addr=%[src]:%[src_port]"`:         true,
		`tcp-check unset-var(txn.from)`:                                     true,
		"tcp-check set-var(check.port)":                                     false,
		"tcp-check set-var(check.port) int(1234) if x":                      false,
		"tcp-check unset-var(txn.from) if x":                                false,
		"---":                                                               false,
		"--- ---":                                                           false,
		`tcp-check set-var-fmt(txn.from) "addr=%[src]:%[src_port] if TRUE"`: true,
	}
	parser := &tcp.Checks{}
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
