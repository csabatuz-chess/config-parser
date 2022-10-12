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

package parsers

import (
	"fmt"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/types"
)

type ErrorLoc303 struct {
	data        *types.ErrorLoc303
	preComments []string // comments that appear before the the actual line
}

func (l *ErrorLoc303) Parse(line string, parts []string, comment string) (string, error) {
	if len(parts) < 3 {
		return "", &errors.ParseError{Parser: "ErrorLoc303", Line: line}
	}
	errorLoc := &types.ErrorLoc303{
		URL:     parts[2],
		Comment: comment,
	}
	code := parts[1]
	if _, ok := errorFileAllowedCode[code]; !ok {
		return "", &errors.ParseError{Parser: "ErrorLoc303", Line: line}
	}
	errorLoc.Code = code
	l.data = errorLoc
	return "", nil
}

func (l *ErrorLoc303) Result() ([]common.ReturnResultLine, error) {
	if l.data == nil {
		return nil, errors.ErrFetch
	}
	return []common.ReturnResultLine{{Data: fmt.Sprintf("errorloc303 %s %s", l.data.Code, l.data.URL), Comment: l.data.Comment}}, nil
}
