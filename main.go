package main

import (
	"tunebot/util/appdata"
	"tunebot/util/menu"
	"tunebot/util/updates"
)

func main() {
	version := "v1.0.3"

	updates.Check("zenyrax", "TuneBot", version, false)
	appdata.Init(version)
	menu.Start()
}
