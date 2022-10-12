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

package integration_test

import (
	"bytes"
	"testing"

	parser "github.com/haproxytech/config-parser/v4"
	"github.com/haproxytech/config-parser/v4/options"
)

func TestWholeConfigsSectionsProgram(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Name, Config string
	}{
		{"program_commandechoHelloWorld", program_commandechoHelloWorld},
		{"program_commandspoamirrorruntime0mirroru", program_commandspoamirrorruntime0mirroru},
	}
	for _, config := range tests {
		t.Run(config.Name, func(t *testing.T) {
			t.Parallel()
			var buffer bytes.Buffer
			buffer.WriteString(config.Config)
			p, err := parser.New(options.Reader(&buffer))
			if err != nil {
				t.Fatalf(err.Error())
			}
			result := p.String()
			if result != config.Config {
				compare(t, config.Config, result)
				t.Error("======== ORIGINAL =========")
				t.Error(config.Config)
				t.Error("======== RESULT ===========")
				t.Error(result)
				t.Error("===========================")
				t.Fatalf("configurations does not match")
			}
		})
	}
}
