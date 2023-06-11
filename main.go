package main

import (
	"tunebot/util/appdata"
	"tunebot/util/menu"
	"tunebot/util/updates"
)

func main() {
	version := "v1.0.2"

	updates.Check("zenyrax", "TuneBot", version)
	appdata.Init(version)
	menu.Start()
}
