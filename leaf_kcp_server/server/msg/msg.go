package msg

import (
	"leaf/network/json"
)

var JSONProcessor = json.NewProcessor()

func init() {
	JSONProcessor.Register(&C2S_AddUser{})
}

type C2S_AddUser struct {
	UserName string
}
