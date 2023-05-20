package appdata

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
)

var CurrentConfig Config

func WriteConfig() {
	configPath := configdir.LocalConfig("TuneBot")
	configFile := filepath.Join(configPath, "config.json")
	configString, _ := json.Marshal(CurrentConfig)
	os.WriteFile(configFile, configString, 0644)
}
