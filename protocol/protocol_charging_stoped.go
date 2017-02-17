package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type ChargingStoppedPacket struct {
	Uuid            string
	Tid             uint64
	StopReason      uint8
	EndMeterReading uint32
	UserID          string
	StopTime        uint32
	TransactionID   string
	Timestamp       uint64

	DBID      uint32
	StationID uint32
}

func (p *ChargingStoppedPacket) Serialize() []byte {
	status := &Report.ChargingPileStatus{
		DasUuid:   p.Uuid,
		Cpid:      p.Tid,
		Status:    uint32(PROTOCOL_CHARGING_PILE_IDLE),
		Timestamp: p.Timestamp,
		Id:        p.DBID,
		StationId: p.StationID,
	}

	data, _ := proto.Marshal(status)

	return data
}

func ParseChargingStopped(buffer []byte, station_id uint32, id uint32) *ChargingStoppedPacket {
	reader, _, _, tid := ParseHeader(buffer)
	stop_reason, _ := reader.ReadByte()
	end_meter_readging := base.ReadDWord(reader)
	userid := base.ReadString(reader, PROTOCOL_USERID_LEN)
	stop_time := base.ReadDWord(reader)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	time_stamp := base.ReadBcdTime(reader)

	return &ChargingStoppedPacket{
		Uuid:            conf.GetConf().Uuid,
		Tid:             tid,
		StopReason:      stop_reason,
		EndMeterReading: end_meter_readging,
		UserID:          userid,
		StopTime:        stop_time,
		TransactionID:   transaction_id,
		Timestamp:       time_stamp,

		DBID:      id,
		StationID: station_id,
	}
}