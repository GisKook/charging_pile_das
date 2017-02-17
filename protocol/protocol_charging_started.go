package protocol

import (
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
)

type ChargingStartedPacket struct {
	Uuid              string
	Tid               uint64
	StartMeterReading uint32
	UserID            string
	StartTime         uint32
	PinCode           string
	TransactionID     string
	Timestamp         uint64

	DBID      uint32
	StationID uint32
}

func (p *ChargingStartedPacket) Serialize() []byte {
	status := &Report.ChargingPileStatus{
		DasUuid:   p.Uuid,
		Cpid:      p.Tid,
		Status:    uint32(PROTOCOL_CHARGING_PILE_CHARGING),
		Timestamp: p.Timestamp,
		Id:        p.DBID,
		StationId: p.StationID,
	}

	data, _ := proto.Marshal(status)

	return data
}

func ParseChargingStarted(buffer []byte, station_id uint32, id uint32) *ChargingStartedPacket {
	reader, _, _, tid := ParseHeader(buffer)
	start_meter_reading := base.ReadDWord(reader)
	userid := base.ReadString(reader, PROTOCOL_USERID_LEN)
	start_time := base.ReadDWord(reader)
	pin_code := base.ReadString(reader, PROTOCOL_PINCODE_LEN)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	time_stamp := base.ReadBcdTime(reader)

	return &ChargingStartedPacket{
		Uuid:              conf.GetConf().Uuid,
		Tid:               tid,
		StartMeterReading: start_meter_reading,
		UserID:            userid,
		StartTime:         start_time,
		PinCode:           pin_code,
		TransactionID:     transaction_id,
		Timestamp:         time_stamp,
		DBID:              id,
		StationID:         station_id,
	}

}
