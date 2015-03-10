// Copyright [2015] [name of copyright hhh@rutcode.com]

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fsm

import (
	"fmt"
)

func init() {
	readConfig("./conf/fsm.json")
}

func CheckSystem(system string) bool {
	return fsmSystems[system]
}

func GetTargetStatus(system, curStatus, event string, groups []string) (tran transaction, err error) {
	if !CheckSystem(system) {
		err = fmt.Errorf("system not found")
		return
	}
	for _, group := range groups {
		key := generateTransKey(system, group, curStatus, event)
		if t := fsmTrans[key]; t != nil {
			return *t, nil
		}
	}
	err = fmt.Errorf("transaction not found")
	return
}
