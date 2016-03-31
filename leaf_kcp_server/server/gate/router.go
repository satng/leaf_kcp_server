package gate

import "server/msg"
import "server/robot"

func init() {
	msg.JSONProcessor.SetRouter(&msg.C2S_AddUser{}, robot.ChanRPC)
}
