package mq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
)

func event_handler_req_charging_prepare(tid uint64, serial uint32, param []*Report.Param) {
	pkg := protocol.ParseNsqChargingPrepare(tid, serial, param)
	connection := conn.NewConns().GetConn(tid)
	if connection != nil {
		connection.SendToTerm(pkg)
	}
}
