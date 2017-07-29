/*
Copyright 2017 Google Inc.

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

package sqltypes

import (
	"strings"

	querypb "github.com/morphar/sqlparsers/pkg/mysql/query"
)

// MakeTestBindVar makes a *querypb.BindVariable from
// an interface{}.It panics on invalid input.
// This function should only be used for testing.
func MakeTestBindVar(v interface{}) *querypb.BindVariable {
	if v == nil {
		return NullBindVariable
	}
	bv, err := BuildBindVariable(v)
	if err != nil {
		panic(err)
	}
	return bv
}

func split(str string) []string {
	splits := strings.Split(str, "|")
	for i, v := range splits {
		splits[i] = strings.TrimSpace(v)
	}
	return splits
}
