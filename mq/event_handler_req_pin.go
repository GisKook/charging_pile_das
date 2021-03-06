package mq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/protocol"
	"log"
)

func event_handler_req_pin(tid uint64, serial uint32) {
	log.Println("event_handler_req_pin")
	pkg := protocol.ParseNsqReqPinCode(tid, serial)
	connection := conn.NewConns().GetConn(tid)
	if connection != nil {
		connection.SendToTerm(pkg)
	}
}
