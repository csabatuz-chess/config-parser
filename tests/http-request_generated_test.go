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

	"github.com/haproxytech/config-parser/v5/parsers/http"
)

func TestRequestshttp(t *testing.T) {
	tests := map[string]bool{
		`http-request capture req.cook_cnt(FirstVisit),bool len 10`:                            true,
		"http-request set-map(map.lst) %[src] %[req.hdr(X-Value)] if value":                    true,
		"http-request set-map(map.lst) %[src] %[req.hdr(X-Value)]":                             true,
		"http-request add-acl(map.lst) [src]":                                                  true,
		"http-request add-header X-value value":                                                true,
		"http-request cache-use cache-name":                                                    true,
		"http-request cache-use cache-name if FALSE":                                           true,
		"http-request del-acl(map.lst) [src]":                                                  true,
		"http-request allow":                                                                   true,
		"http-request auth":                                                                    true,
		"http-request del-header X-value":                                                      true,
		"http-request del-header X-value if TRUE":                                              true,
		"http-request del-header X-value -m str if TRUE":                                       true,
		"http-request del-map(map.lst) %[src] if ! value":                                      true,
		"http-request del-map(map.lst) %[src]":                                                 true,
		"http-request deny":                                                                    true,
		"http-request deny deny_status 400":                                                    true,
		"http-request deny if TRUE":                                                            true,
		"http-request deny deny_status 400 if TRUE":                                            true,
		"http-request deny deny_status 400 content-type application/json if TRUE":              true,
		"http-request deny deny_status 400 content-type application/json":                      true,
		"http-request deny deny_status 400 content-type application/json default-errorfiles":   true,
		"http-request deny deny_status 400 content-type application/json errorfile errors":     true,
		"http-request deny deny_status 400 content-type application/json string error if TRUE": true,
		"http-request deny deny_status 400 content-type application/json lf-string error hdr host google.com if TRUE":              true,
		"http-request deny deny_status 400 content-type application/json file /var/errors.file":                                    true,
		"http-request deny deny_status 400 content-type application/json lf-file /var/errors.file":                                 true,
		"http-request deny deny_status 400 content-type application/json string error hdr host google.com if TRUE":                 true,
		"http-request deny deny_status 400 content-type application/json string error hdr host google.com hdr x-value bla if TRUE": true,
		"http-request deny deny_status 400 content-type application/json string error hdr host google.com hdr x-value bla":         true,
		"http-request disable-l7-retry":                                                          true,
		"http-request disable-l7-retry if FALSE":                                                 true,
		"http-request early-hint hint %[src]":                                                    true,
		"http-request early-hint hint %[src] if FALSE":                                           true,
		"http-request early-hint if FALSE":                                                       true,
		"http-request lua.foo":                                                                   true,
		"http-request lua.foo if FALSE":                                                          true,
		"http-request lua.foo param":                                                             true,
		"http-request lua.foo param param2":                                                      true,
		"http-request normalize-uri fragment-encode":                                             true,
		"http-request normalize-uri fragment-encode if TRUE":                                     true,
		"http-request normalize-uri fragment-strip":                                              true,
		"http-request normalize-uri fragment-strip if TRUE":                                      true,
		"http-request normalize-uri path-merge-slashes":                                          true,
		"http-request normalize-uri path-merge-slashes if TRUE":                                  true,
		"http-request normalize-uri path-strip-dot":                                              true,
		"http-request normalize-uri path-strip-dot if TRUE":                                      true,
		"http-request normalize-uri path-strip-dotdot":                                           true,
		"http-request normalize-uri path-strip-dotdot full":                                      true,
		"http-request normalize-uri path-strip-dotdot if TRUE":                                   true,
		"http-request normalize-uri path-strip-dotdot full if TRUE":                              true,
		"http-request normalize-uri percent-decode-unreserved":                                   true,
		"http-request normalize-uri percent-decode-unreserved if TRUE":                           true,
		"http-request normalize-uri percent-decode-unreserved strict":                            true,
		"http-request normalize-uri percent-decode-unreserved strict if TRUE":                    true,
		"http-request normalize-uri percent-to-uppercase":                                        true,
		"http-request normalize-uri percent-to-uppercase if TRUE":                                true,
		"http-request normalize-uri percent-to-uppercase strict":                                 true,
		"http-request normalize-uri percent-to-uppercase strict if TRUE":                         true,
		"http-request normalize-uri query-sort-by-name":                                          true,
		"http-request normalize-uri query-sort-by-name if TRUE":                                  true,
		"http-request redirect prefix https://mysite.com":                                        true,
		"http-request reject":                                                                    true,
		"http-request replace-header User-agent curl foo":                                        true,
		"http-request replace-path (.*) /foo":                                                    true,
		"http-request replace-path (.*) /foo if TRUE":                                            true,
		"http-request replace-pathq (.*) /foo":                                                   true,
		"http-request replace-pathq (.*) /foo if TRUE":                                           true,
		"http-request replace-uri ^http://(.*) https://1":                                        true,
		"http-request replace-uri ^http://(.*) https://1 if FALSE":                               true,
		"http-request replace-value X-Forwarded-For ^192.168.(.*)$ 172.16.1":                     true,
		"http-request sc-add-gpc(1,2) 1":                                                         true,
		"http-request sc-add-gpc(1,2) 1 if is-error":                                             true,
		"http-request sc-inc-gpc(1,2)":                                                           true,
		"http-request sc-inc-gpc(1,2) if FALSE":                                                  true,
		"http-request sc-inc-gpc0(1)":                                                            true,
		"http-request sc-inc-gpc0(1) if FALSE":                                                   true,
		"http-request sc-inc-gpc1(1)":                                                            true,
		"http-request sc-inc-gpc1(1) if FALSE":                                                   true,
		"http-request sc-set-gpt0(1) hdr(Host),lower":                                            true,
		"http-request sc-set-gpt0(1) 10":                                                         true,
		"http-request sc-set-gpt0(1) hdr(Host),lower if FALSE":                                   true,
		"http-request send-spoe-group engine group":                                              true,
		"http-request set-header X-value value":                                                  true,
		"http-request set-log-level silent":                                                      true,
		"http-request set-mark 20":                                                               true,
		"http-request set-mark 0x1Ab":                                                            true,
		"http-request set-nice 0":                                                                true,
		"http-request set-nice 0 if FALSE":                                                       true,
		"http-request set-method POST":                                                           true,
		"http-request set-method POST if FALSE":                                                  true,
		"http-request set-path /%[hdr(host)]%[path]":                                             true,
		"http-request set-pathq /%[hdr(host)]%[path]":                                            true,
		"http-request set-priority-class req.hdr(priority)":                                      true,
		"http-request set-priority-class req.hdr(priority) if FALSE":                             true,
		"http-request set-priority-offset req.hdr(offset)":                                       true,
		"http-request set-priority-offset req.hdr(offset) if FALSE":                              true,
		"http-request set-query %[query,regsub(%3D,=,g)]":                                        true,
		"http-request set-src hdr(src)":                                                          true,
		"http-request set-src hdr(src) if FALSE":                                                 true,
		"http-request set-src-port hdr(port)":                                                    true,
		"http-request set-src-port hdr(port) if FALSE":                                           true,
		"http-request set-timeout server 20":                                                     true,
		"http-request set-timeout tunnel 20":                                                     true,
		"http-request set-timeout tunnel 20s if TRUE":                                            true,
		"http-request set-timeout server 20s if TRUE":                                            true,
		"http-request set-tos 0 if FALSE":                                                        true,
		"http-request set-tos 0":                                                                 true,
		"http-request set-uri /%[hdr(host)]%[path]":                                              true,
		"http-request set-var(req.my_var) req.fhdr(user-agent),lower":                            true,
		"http-request set-var-fmt(req.my_var) req.fhdr(user-agent),lower":                        true,
		"http-request silent-drop":                                                               true,
		"http-request silent-drop if FALSE":                                                      true,
		"http-request strict-mode on":                                                            true,
		"http-request strict-mode on if FALSE":                                                   true,
		"http-request tarpit":                                                                    true,
		"http-request tarpit deny_status 400":                                                    true,
		"http-request tarpit if TRUE":                                                            true,
		"http-request tarpit deny_status 400 if TRUE":                                            true,
		"http-request tarpit deny_status 400 content-type application/json if TRUE":              true,
		"http-request tarpit deny_status 400 content-type application/json":                      true,
		"http-request tarpit deny_status 400 content-type application/json default-errorfiles":   true,
		"http-request tarpit deny_status 400 content-type application/json errorfile errors":     true,
		"http-request tarpit deny_status 400 content-type application/json string error if TRUE": true,
		"http-request tarpit deny_status 400 content-type application/json lf-string error hdr host google.com if TRUE":              true,
		"http-request tarpit deny_status 400 content-type application/json file /var/errors.file":                                    true,
		"http-request tarpit deny_status 400 content-type application/json lf-file /var/errors.file":                                 true,
		"http-request tarpit deny_status 400 content-type application/json string error hdr host google.com if TRUE":                 true,
		"http-request tarpit deny_status 400 content-type application/json string error hdr host google.com hdr x-value bla if TRUE": true,
		"http-request tarpit deny_status 400 content-type application/json string error hdr host google.com hdr x-value bla":         true,
		"http-request track-sc0 src":                                                                                         true,
		"http-request track-sc1 src":                                                                                         true,
		"http-request track-sc2 src":                                                                                         true,
		"http-request track-sc5 src":                                                                                         true,
		"http-request track-sc5 src table a_table":                                                                           true,
		"http-request track-sc5 src table a_table if some_cond":                                                              true,
		"http-request track-sc5 src if some_cond":                                                                            true,
		"http-request unset-var(req.my_var)":                                                                                 true,
		"http-request unset-var(req.my_var) if FALSE":                                                                        true,
		"http-request wait-for-body time 20s":                                                                                true,
		"http-request wait-for-body time 20s if TRUE":                                                                        true,
		"http-request wait-for-body time 20s at-least 100k":                                                                  true,
		"http-request wait-for-body time 20s at-least 100k if TRUE":                                                          true,
		"http-request wait-for-handshake":                                                                                    true,
		"http-request wait-for-handshake if FALSE":                                                                           true,
		"http-request do-resolve(txn.myip,mydns) hdr(Host),lower":                                                            true,
		"http-request do-resolve(txn.myip,mydns) hdr(Host),lower if { var(txn.myip) -m found }":                              true,
		"http-request do-resolve(txn.myip,mydns) hdr(Host),lower unless { var(txn.myip) -m found }":                          true,
		"http-request do-resolve(txn.myip,mydns,ipv4) hdr(Host),lower":                                                       true,
		"http-request do-resolve(txn.myip,mydns,ipv6) hdr(Host),lower":                                                       true,
		"http-request set-dst var(txn.myip)":                                                                                 true,
		"http-request set-dst var(txn.myip) if { var(txn.myip) -m found }":                                                   true,
		"http-request set-dst var(txn.myip) unless { var(txn.myip) -m found }":                                               true,
		"http-request set-dst-port hdr(x-port)":                                                                              true,
		"http-request set-dst-port hdr(x-port) if { var(txn.myip) -m found }":                                                true,
		"http-request set-dst-port hdr(x-port) unless { var(txn.myip) -m found }":                                            true,
		"http-request set-dst-port int(4000)":                                                                                true,
		"http-request return status 400 default-errorfiles if { var(txn.myip) -m found }":                                    true,
		"http-request return status 400 errorfile /my/fancy/errorfile if { var(txn.myip) -m found }":                         true,
		"http-request return status 400 errorfiles myerror if { var(txn.myip) -m found }":                                    true,
		"http-request redirect location /file.html if { var(txn.routecookie) -m found } !{ var(txn.pod),nbsrv -m found }:1]": true,
		"http-request set-bandwidth-limit my-limit":                                                                          true,
		"http-request set-bandwidth-limit my-limit limit 1m period 10s":                                                      true,
		"http-request set-bandwidth-limit my-limit period 10s":                                                               true,
		"http-request set-bandwidth-limit my-limit limit 1m":                                                                 true,
		`http-request add-header Authorization Basic\ eC1oYXByb3h5LXJlY3J1aXRzOlBlb3BsZSB3aG8gZGVjb2RlIG1lc3NhZ2VzIG9mdGVuIGxvdmUgd29ya2luZyBhdCBIQVByb3h5LiBEbyBub3QgYmUgc2h5LCBjb250YWN0IHVz`:  true,
		`http-request add-header Authorisation "Basic eC1oYXByb3h5LXJlY3J1aXRzOlBlb3BsZSB3aG8gZGVjb2RlIG1lc3NhZ2VzIG9mdGVuIGxvdmUgd29ya2luZyBhdCBIQVByb3h5LiBEbyBub3QgYmUgc2h5LCBjb250YWN0IHVz"`: true,
		`http-request return status 200 content-type "text/plain" string "My content" if { var(txn.myip) -m found }`:                                                                             true,
		`http-request return status 200 content-type "text/plain" string "My content" unless { var(txn.myip) -m found }`:                                                                         true,
		`http-request return content-type "text/plain" string "My content" if { var(txn.myip) -m found }`:                                                                                        true,
		`http-request return content-type 'text/plain' string 'My content' if { var(txn.myip) -m found }`:                                                                                        true,
		`http-request return content-type "text/plain" lf-string "Hello, you are: %[src]" if { var(txn.myip) -m found }`:                                                                         true,
		`http-request return content-type "text/plain" file /my/fancy/response/file if { var(txn.myip) -m found }`:                                                                               true,
		`http-request return content-type "text/plain" lf-file /my/fancy/lof/format/response/file if { var(txn.myip) -m found }`:                                                                 true,
		`http-request return content-type "text/plain" string "My content" hdr X-value value if { var(txn.myip) -m found }`:                                                                      true,
		`http-request return content-type "text/plain" string "My content" hdr X-value x-value hdr Y-value y-value if { var(txn.myip) -m found }`:                                                true,
		`http-request return content-type "text/plain" lf-string "Hello, you are: %[src]"`:                                                                                                       true,
		`http-request redirect location /file.html if { var(txn.routecookie) "ROUTEMP" }:1`:                                                                                                      true,
		"http-request": false,
		"http-request capture req.cook_cnt(FirstVisit),bool strlen 10":            false,
		"http-request set-map(map.lst) %[src]":                                    false,
		"http-request add-acl(map.lst)":                                           false,
		"http-request add-header X-value":                                         false,
		"http-request cache-use":                                                  false,
		"http-request cache-use if FALSE":                                         false,
		"http-request del-acl(map.lst)":                                           false,
		"http-request del-header":                                                 false,
		"http-request del-header X-value -m if TRUE":                              false,
		"http-request del-header X-value bla":                                     false,
		"http-request del-map(map.lst)":                                           false,
		"http-request deny test test":                                             false,
		"http-request early-hint hint":                                            false,
		"http-request early-hint hint if FALSE":                                   false,
		"http-request lua.":                                                       false,
		"http-request lua. if FALSE":                                              false,
		"http-request lua. param":                                                 false,
		"http-request normalize-uri bla":                                          false,
		"http-request normalize-uri path-strip-dot strict":                        false,
		"http-request normalize-uri path-strip-dot full":                          false,
		"http-request normalize-uri if TRUE":                                      false,
		"http-request normalize-uri":                                              false,
		"http-request redirect prefix":                                            false,
		"http-request replace-header User-agent curl":                             false,
		"http-request replace-path (.*)":                                          false,
		"http-request replace-path (.*) if TRUE":                                  false,
		"http-request replace-pathq (.*)":                                         false,
		"http-request replace-pathq (.*) if TRUE":                                 false,
		"http-request replace-uri ^http://(.*)":                                   false,
		"http-request replace-uri":                                                false,
		"http-request replace-uri ^http://(.*) if FALSE":                          false,
		"http-request replace-value X-Forwarded-For ^192.168.(.*)$":               false,
		"http-request sc-add-gpc":                                                 false,
		"http-request sc-inc-gpc":                                                 false,
		"http-request sc-inc-gpc0":                                                false,
		"http-request sc-inc-gpc1":                                                false,
		"http-request sc-set-gpt0(1)":                                             false,
		"http-request sc-set-gpt0":                                                false,
		"http-request sc-set-gpt0(1) if FALSE":                                    false,
		"http-request send-spoe-group engine":                                     false,
		"http-request set-header X-value":                                         false,
		"http-request set-log-level":                                              false,
		"http-request set-mark":                                                   false,
		"http-request set-nice":                                                   false,
		"http-request set-method":                                                 false,
		"http-request set-path":                                                   false,
		"http-request set-pathq":                                                  false,
		"http-request set-priority-class":                                         false,
		"http-request set-priority-offset":                                        false,
		"http-request set-query":                                                  false,
		"http-request set-src":                                                    false,
		"http-request set-src-port":                                               false,
		"http-request set-timeout client 20":                                      false,
		"http-request set-tos":                                                    false,
		"http-request set-uri":                                                    false,
		"http-request set-var(req.my_var)":                                        false,
		"http-request set-var-fmt(req.my_var)":                                    false,
		"http-request strict-mode":                                                false,
		"http-request strict-mode if FALSE":                                       false,
		"http-request tarpit test test":                                           false,
		"http-request track-sc0":                                                  false,
		"http-request track-sc1":                                                  false,
		"http-request track-sc2":                                                  false,
		"http-request track-sc":                                                   false,
		"http-request track-sc5":                                                  false,
		"http-request track-sc5 src table":                                        false,
		"http-request track-sc5 src if":                                           false,
		"http-request track-sc src if some_cond":                                  false,
		"http-request track-sc src table a_table if some_cond":                    false,
		"http-request unset-var(req.)":                                            false,
		"http-request unset-var(req)":                                             false,
		"http-request wait-for-body 20s at-least 100k":                            false,
		"http-request wait-for-body time 2000 test":                               false,
		"http-request do-resolve(txn.myip)":                                       false,
		"http-request do-resolve(txn.myip,mydns)":                                 false,
		"http-request do-resolve(txn.myip,mydns,ipv4)":                            false,
		"http-request set-dst":                                                    false,
		"http-request set-dst-port":                                               false,
		"http-request return 8 t hdr":                                             false,
		"http-request return hdr":                                                 false,
		"http-request return hdr one":                                             false,
		"http-request return errorfile":                                           false,
		"http-request return 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 file":              false,
		"http-request return 0 hdr":                                               false,
		"http-request return 0 0 hdr 0":                                           false,
		"http-request return e r s n s c m	t e r  s c t e s t e r s c v e hdr ï": false,
		"http-request redirect location if { var(txn.routecookie) -m found } !{ var(txn.pod),nbsrv -m found }:1]":                 false,
		"http-request redirect location /file.html code if { var(txn.routecookie) -m found } !{ var(txn.pod),nbsrv -m found }:1]": false,
		"http-request set-bandwidth-limit my-limit limit":                                                                         false,
		"http-request set-bandwidth-limit my-limit period":                                                                        false,
		"http-request set-bandwidth-limit my-limit 10s":                                                                           false,
		"http-request set-bandwidth-limit my-limit period 10s limit":                                                              false,
		"http-request set-bandwidth-limit my-limit limit period 10s":                                                              false,
		"---":     false,
		"--- ---": false,
	}
	parser := &http.Requests{}
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
