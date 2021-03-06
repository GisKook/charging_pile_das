package event_handler

import (
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/conn"
	"github.com/giskook/charging_pile_das/pkg"
	"github.com/giskook/charging_pile_das/protocol"
	"github.com/giskook/charging_pile_das/server"
	"github.com/giskook/gotcp"
)

func event_handler_login(c *gotcp.Conn, p *pkg.Charging_Pile_Packet) {
	connection := c.GetExtraData().(*conn.Conn)
	login_pkg := p.Packet.(*protocol.LoginPacket)
	connection.ID = login_pkg.Tid
	conn.NewConns().SetID(login_pkg.Tid, connection)

	connection.Charging_Pile.Status = login_pkg.Status
	connection.Charging_Pile.Timestamp = login_pkg.Timestamp
	if login_pkg.Status == 1 {
		connection.Charging_Pile.EndMeterReading = login_pkg.EndMeterReading
		connection.Charging_Pile.StopTime = login_pkg.StopTime
		connection.Charging_Pile.TransactionID = login_pkg.TransactionID
	}

	server.GetServer().MQ.Send(conf.GetConf().Nsq.Producer.TopicAuth, p.Serialize())
}
