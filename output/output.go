/*
   Copyright 2014 Outbrain Inc.

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

// output provides with controlled printing functions
package output

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Printer interface {
	PrintString(data []byte)
	PrintStringArray(array []string)
}

type TxtPrinter struct {
	OmitTrailingNL bool
}

func (p TxtPrinter) PrintString(data []byte) {
	s := fmt.Sprintf("%s", data)
	if p.OmitTrailingNL {
		fmt.Print(s)
	} else {
		fmt.Println(s)
	}
}

func (p TxtPrinter) PrintStringArray(stringArray []string) {
	s := fmt.Sprintf("%s", strings.Join(stringArray, "\n"))
	if p.OmitTrailingNL {
		fmt.Print(s)
	} else {
		fmt.Println(s)
	}
}

type JSONPrinter struct{}

func (_ JSONPrinter) PrintString(data []byte) {
	s, _ := json.Marshal(string(data))
	fmt.Println(string(s))
}

func (_ JSONPrinter) PrintStringArray(stringArray []string) {
	s, _ := json.Marshal(stringArray)
	fmt.Println(string(s))
}