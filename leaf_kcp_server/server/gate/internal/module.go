package internal

import (
	"server/conf"
	"server/msg"
	"server/robot"

	"leaf/gate"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,

		LenMsgLen:    conf.LenMsgLen,
		LittleEndian: conf.LittleEndian,
		Processor:    msg.JSONProcessor,
		AgentChanRPC: robot.ChanRPC,

		UDPAddr: conf.Server.UDPAddr,
	}
}
