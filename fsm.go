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

// check system whether it exists in fsm config,
// if not return false, then true
func CheckSystem(system string) bool {
	return fsmSystems[system]
}

// Get target status by current status and event,
// Current status now is business status,
// event is a trigger to find target status,
// group means several transation paths could combine together.
func GetTargetStatus(system, curStatus, event string, groups ...string) (tran Transaction, err error) {
	if !CheckSystem(system) {
		err = fmt.Errorf("system not found")
		return
	}

	if len(groups) == 0 {
		t := fsmTrans[generateTransKey(system, defaultGroupName, curStatus, event)]
		if t != nil {
			return *t, nil
		}

	} else {
		for _, group := range groups {
			key := generateTransKey(system, group, curStatus, event)
			if t := fsmTrans[key]; t != nil {
				return *t, nil
			}
		}
	}

	err = fmt.Errorf("transaction not found")
	return
}
