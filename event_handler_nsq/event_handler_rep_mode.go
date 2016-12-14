package event_handler_nsq

import (
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/giskook/charging_pile_das/protocol"
	"log"
)

func event_handler_rep_mode(tid uint64, serial uint32, param []*Report.Param) {
	log.Println(tid)
	pkg := protocol.ParseNsqMode(tid, param)
	connection := conn.NewConns().GetConn(tid)
	if connection != nil {
		connection.SendToTerm(pkg)
	}
}