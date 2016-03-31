package conf

import (
	"encoding/json"
	"io/ioutil"
	"leaf/log"
)

var Server struct {
	LogLevel   string
	LogPath    string
	MaxConnNum int

	UDPAddr string
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
