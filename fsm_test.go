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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFSM(t *testing.T) {
	// fsm.GetTargetStatus("fsm1", "status1", "save", []string{"admin_group"})

	Convey("failed get target status", t, func() {
		Convey("when system not in conf", func() {
			Convey("will return system not found", func() {
				_, err := GetTargetStatus("fsm3", "status1", "event1", []string{"group1"})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "system not found")
			})
		})
		Convey("when status not in system", func() {
			Convey("will return transaction not found", func() {
				_, err := GetTargetStatus("fsm1", "status2", "event1", []string{"group1"})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "transaction not found")
			})
		})
		Convey("when event not in system", func() {
			Convey("will return transaction not found", func() {
				_, err := GetTargetStatus("fsm1", "status1", "event", []string{"group1"})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "transaction not found")
			})
		})
		Convey("when group not in system", func() {
			Convey("will return transaction not found", func() {
				_, err := GetTargetStatus("fsm1", "status1", "event1", []string{"group3"})
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "transaction not found")
			})
		})
	})
	Convey("success get target status", t, func() {
		Convey("get test fsm", func() {
			Convey("will return transaction", func() {
				tran, err := GetTargetStatus("fsm1", "status1", "event1", []string{"group1"})
				So(err, ShouldBeNil)
				So(tran.TargetStatus, ShouldEqual, "status2")
				tran, err = GetTargetStatus("fsm1", "status2", "event2", []string{"group2", "group3"})
				So(err, ShouldBeNil)
				So(tran.TargetStatus, ShouldEqual, "status3")

				tran, err = GetTargetStatus("fsm2", "status1", "event1", []string{"group1"})
				So(err, ShouldBeNil)
				So(tran.TargetStatus, ShouldEqual, "status2")
				tran, err = GetTargetStatus("fsm2", "status2", "event2", []string{"group1"})
				So(err, ShouldBeNil)
				So(tran.TargetStatus, ShouldEqual, "status3")
			})
		})
	})
}
