package gate

type Agent interface {
	Write(msg interface{})
	Close()
	UserData() interface{}
	SetUserData(data interface{})
}
