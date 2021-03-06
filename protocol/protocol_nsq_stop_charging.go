package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
)

type StopChargingNsqPacket struct {
	Tid uint64
}

func (p *StopChargingNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_STOP_CHARGING, p.Tid)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqStopCharging(cpid uint64) *StopChargingNsqPacket {
	return &StopChargingNsqPacket{
		Tid: cpid,
	}
}
