package internal

import (
	"fmt"
	"leaf/gate"
	"reflect"
	"server/msg"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.C2S_AddUser{}, handleAddUser)
}

func handleAddUser(args []interface{}) {
	m := args[0].(*msg.C2S_AddUser)
	a := args[1].(gate.Agent)

	a.Write("123456789")
	fmt.Println(m.UserName)
}
