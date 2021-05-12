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

package actions

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/v4/common"
)

type UseService struct {
	Name     string
	Cond     string
	CondTest string
	Comment  string
}

// Parse parses http-request user-service <name> [ { if | unless } <condition> ]
func (us *UseService) Parse(parts []string, comment string) error {
	if comment != "" {
		us.Comment = comment
	}

	if len(parts) >= 3 {
		_, condition := common.SplitRequest(parts[3:])
		us.Name = parts[2]
		if len(condition) > 1 {
			us.Cond = condition[0]
			us.CondTest = strings.Join(condition[1:], " ")
		}
		return nil
	}
	return fmt.Errorf("not enough params")
}

func (us *UseService) String() string {
	condition := ""
	if us.Cond != "" {
		condition = fmt.Sprintf(" %s %s", us.Cond, us.CondTest)
	}
	return fmt.Sprintf("use-service %s%s", us.Name, condition)
}

func (us *UseService) GetComment() string {
	return us.Comment
}
