package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/pb"
)

type ChargingNsqPacket struct {
	Tid              uint64
	Serial           uint32
	Userid           string
	PinCode          string
	TransactionID    string
	TranscationValue uint32
}

func (p *ChargingNsqPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REQ_CHARGING, p.Tid)
	base.WriteString(&writer, p.Userid)
	base.WriteString(&writer, p.PinCode)
	base.WriteBcdString(&writer, p.TransactionID)
	base.WriteDWord(&writer, p.TranscationValue)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseNsqCharging(cpid uint64, serial uint32, param []*Report.Param) *ChargingNsqPacket {
	return &ChargingNsqPacket{
		Tid:              cpid,
		Serial:           serial,
		Userid:           param[0].Strpara,
		PinCode:          param[1].Strpara,
		TransactionID:    param[2].Strpara,
		TranscationValue: uint32(param[3].Npara),
	}
}
