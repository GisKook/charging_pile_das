package protocol

import (
	"bytes"
	"github.com/giskook/charging_pile_das/base"
	"github.com/giskook/charging_pile_das/conf"
	"github.com/giskook/charging_pile_das/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

type ChargingStartedPacket struct {
	Uuid              string
	Tid               uint64
	StartMeterReading uint32
	UserID            string
	StartTime         uint32
	PinCode           string
	TransactionID     string
	PreCharge         uint32
	Timestamp         uint64

	DBID      uint32
	StationID uint32
}

func (p *ChargingStartedPacket) SerializeTss() []byte {
	status := &Report.ChargingPileStatus{
		DasUuid:            p.Uuid,
		Cpid:               p.Tid,
		Status:             uint32(PROTOCOL_CHARGE_PILE_STATUS_STARTED),
		StartMeterReading:  float32(p.StartMeterReading) / 10.0, // protocol in charge pile is 0.1 degree
		StartTime:          uint64(p.StartTime),
		Timestamp:          p.Timestamp,
		ChargingDuration:   0,
		ChargingCapacity:   0.0,
		ChargingCost:       0.0,
		PreCharge:          float32(p.PreCharge) / 100.0,
		CurrentOrderNumber: p.TransactionID,
		Id:                 p.DBID,
		StationId:          p.StationID,
	}

	data, _ := proto.Marshal(status)

	return data
}

func (p *ChargingStartedPacket) SerializeWeChat() []byte {
	log.Println(p.TransactionID)
	command := &Report.Command{
		Type: Report.Command_CMT_REP_CHARGING_STARTED,
		Uuid: p.Uuid,
		Tid:  p.Tid,
		Paras: []*Report.Param{
			&Report.Param{
				Type:    Report.Param_STRING,
				Strpara: p.UserID,
			},
			&Report.Param{
				Type:    Report.Param_STRING,
				Strpara: p.TransactionID,
			},
		},
	}

	data, _ := proto.Marshal(command)

	return data

}

func (p *ChargingStartedPacket) Serialize() []byte {
	var writer bytes.Buffer
	WriteHeader(&writer, 0,
		PROTOCOL_REP_CHARGING_STARTED_FEEDBACK, p.Tid)
	base.WriteLength(&writer)
	base.WriteWord(&writer, CalcCRC(writer.Bytes()[1:], uint16(writer.Len()-1)))
	writer.WriteByte(PROTOCOL_END_FLAG)

	return writer.Bytes()
}

func ParseChargingStarted(buffer []byte, station_id uint32, id uint32) *ChargingStartedPacket {
	reader, _, _, tid := ParseHeader(buffer)
	start_meter_reading := base.ReadDWord(reader)
	userid := base.ReadString(reader, PROTOCOL_USERID_LEN)
	start_time := base.ReadDWord(reader)
	pin_code := base.ReadString(reader, PROTOCOL_PINCODE_LEN)
	transaction_id := base.ReadBcdString(reader, PROTOCOL_TRANSACTION_BCD_LEN)
	pre_charge := base.ReadDWord(reader)
	time_stamp := base.ReadBcdTime(reader)

	return &ChargingStartedPacket{
		Uuid:              conf.GetConf().Uuid,
		Tid:               tid,
		StartMeterReading: start_meter_reading,
		UserID:            userid,
		StartTime:         start_time,
		PinCode:           pin_code,
		PreCharge:         pre_charge,
		TransactionID:     transaction_id,
		Timestamp:         time_stamp,
		DBID:              id,
		StationID:         station_id,
	}

}
