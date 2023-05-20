package main

import (
	"tunebot/util/appdata"
	"tunebot/util/menu"
	"tunebot/util/updates"
)

func main() {
	appdata.Init()
	appdata.CurrentConfig.Version = updates.SetVersion()
	menu.Start()
}
