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
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var fsmTrans map[string]*transaction
var fsmSystems map[string]bool

type fsm struct {
	SystemFSM []systemFSM `json:"fsm"`
}

type systemFSM struct {
	Name    string `json:"name"`
	FSMSets sets   `json:"sets"`
}

type sets struct {
	Status             []string            `json:"status"`
	Events             []string            `json:"events"`
	Transactions       []transaction       `json:"transactions"`
	TransactionsGroups []transactionsGroup `json:"transactions_group"`
}

type transaction struct {
	Id            int32  `json:"id"`
	CurrentStatus string `json:"current"`
	Event         string `json:"event"`
	TargetStatus  string `json:"target"`
}

type transactionsGroup struct {
	GroupName      string  `json:"name"`
	TransactionIds []int32 `json:"transaction_ids"`
}

func readConfig(filename string) (fsmConfigs *fsm) {
	var err error
	jsonData := []byte{}
	if jsonData, err = ioutil.ReadFile(filename); err != nil {
		panic(err)
	}

	fsmConfigs = new(fsm)
	if err = json.Unmarshal(jsonData, fsmConfigs); err != nil {
		panic(err)
	}
	fsmSystems = make(map[string]bool)
	fsmTrans = make(map[string]*transaction)
	for _, v := range fsmConfigs.SystemFSM {
		if fsmSystems[v.Name] {
			panic(fmt.Errorf("fsm name(%s) is already have, check your configure", v.Name))
		}
		fsmSystems[v.Name] = true
		events := make(map[string]bool)
		for _, event := range v.FSMSets.Events {
			k := fmt.Sprintf("%s:%s", v.Name, event)
			if events[k] {
				panic(fmt.Errorf("fsm repeat name(%s): status(%s), check your configure", v.Name, event))
			}
			events[k] = true
		}
		status := make(map[string]bool)
		for _, s := range v.FSMSets.Status {
			k := fmt.Sprintf("%s:%s", v.Name, s)
			if status[k] {
				panic(fmt.Errorf("fsm repeat name(%s): event(%s), check your configure", v.Name, s))
			}
			status[k] = true
		}

		checkTrans := make(map[string]bool)
		trans := make(map[int32]*transaction)
		for _, t := range v.FSMSets.Transactions {
			fmtCstatus := fmt.Sprintf("%s:%s", v.Name, t.CurrentStatus)
			fmtTstatus := fmt.Sprintf("%s:%s", v.Name, t.TargetStatus)
			fmtEvent := fmt.Sprintf("%s:%s", v.Name, t.Event)
			tran := t
			if trans[tran.Id] != nil {
				panic(fmt.Errorf("fsm name(%s): repeat transaction id (%d), check your configure", v.Name, t.Id))
			}
			if !status[fmtCstatus] {
				panic(fmt.Errorf("fsm name(%s): current status(%s) not in array status, check your configure", v.Name, t.CurrentStatus))
			}
			if !status[fmtTstatus] {
				panic(fmt.Errorf("fsm name(%s): target status(%s) not in array status, check your configure", v.Name, t.TargetStatus))
			}
			if !events[fmtEvent] {
				panic(fmt.Errorf("fsm name(%s): event(%s) not in array events, check your configure", v.Name, t.Event))
			}

			k := fmt.Sprintf("%s:%s:%s", v.Name, t.CurrentStatus, t.Event)
			if checkTrans[k] {
				panic(fmt.Errorf("fsm transaction is already exist: %s", k))
			}
			checkTrans[k] = true
			trans[t.Id] = &tran
		}
		checkOrgGroup := make(map[string]bool)
		for _, ot := range v.FSMSets.TransactionsGroups {
			if checkOrgGroup[ot.GroupName] {
				panic(fmt.Errorf("fsm name(%s): repeat group name(%s), check your configure", v.Name, ot.GroupName))
			}
			checkOrgGroup[ot.GroupName] = true
			for _, id := range ot.TransactionIds {
				t := trans[id]
				if t == nil {
					panic(fmt.Errorf("fsm name(%s): transaction id(%d) not in array transaction, check your configure", v.Name, id))
				}
				k := generateTransKey(v.Name, ot.GroupName, t.CurrentStatus, t.Event)
				fsmTrans[k] = t
			}
		}
	}
	return
}

func generateTransKey(system, group, curStatus, event string) string {
	return fmt.Sprintf("%s:%s:%s:%s", system, group, curStatus, event)
}
