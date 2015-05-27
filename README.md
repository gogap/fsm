# fsm

Finite-state machine in go

* [![Build Status](https://travis-ci.org/go-rut/fsm.png)](https://travis-ci.org/go-rut/fsm)
* [![GoDoc](http://godoc.org/github.com/go-rut/fsm?status.svg)](http://godoc.org/github.com/go-rut/fsm)


* [点击进入中文相关说明](http://zh.wikipedia.org/wiki/%E6%9C%89%E9%99%90%E7%8A%B6%E6%80%81%E6%9C%BA)
* [Click to article in English](http://en.wikipedia.org/wiki/Finite-state_machine)

## Sample

**Config**

```
{
  "fsm": [
    {
      "name": "fsm1",
      "sets": {
        "status": ["status1", "status2", "status3"],
        "events": ["event1", "event2"],
        "transactions": [
          {"id": 1, "current": "status1", "event": "event1", "target": "status2"},
          {"id": 2, "current": "status2", "event": "event2", "target": "status3"}
        ],
        "transactions_group": [
          {"name": "group1", "transaction_ids": [1]},
          {"name": "group2", "transaction_ids": [2]}
        ]
      }
    },
    {
      "name": "fsm2",
      "sets": {
        "status": ["status1", "status2", "status3"],
        "events": ["event1", "event2"],
        "transactions": [
          {"id": 1, "current": "status1", "event": "event1", "target": "status2"},
          {"id": 2, "current": "status2", "event": "event2", "target": "status3"}
        ],
        "transactions_group": [
          {"name": "group1", "transaction_ids": [1, 2]}
        ]
      }
    }
  ]
}
```

**TEST**

fsm_test.go

```
func TestFSM(t *testing.T) {

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

```
