package main

import (
	"server/conf"
	"server/gate"
	"server/login"
	"server/robot"

	"leaf"
	lconf "leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath

	leaf.Run(
		robot.Module,
		gate.Module,
		login.Module,
	)
}
