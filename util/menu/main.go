package menu

import (
	"os"
	"os/exec"
	"runtime"

	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/pkg/browser"

	"github.com/AlecAivazis/survey/v2"
)

func Start() {
	Reset()

	option := ""
	prompt := &survey.Select{
		Message: "What do you want to do?",
		Options: []string{"Run Tasks", "Manage Tasks", "Manage Proxies", "Manage Settings", "Quit"},
	}
	survey.AskOne(prompt, &option)

	switch option {
	case "Run Tasks":
		Tasks()
	case "Manage Tasks":
		configPath := configdir.LocalConfig("TuneBot")
		folder := filepath.Join(configPath, "tasks/")
		switch runtime.GOOS {
		case "windows":
			cmd := exec.Command(`explorer`, folder)
			cmd.Run()
		case "darwin":
			cmd := exec.Command(`open`, folder)
			cmd.Run()
		}
		Start()
	case "Manage Proxies":
		configPath := configdir.LocalConfig("TuneBot")
		folder := filepath.Join(configPath, "proxies/")
		switch runtime.GOOS {
		case "windows":
			cmd := exec.Command(`explorer`, folder)
			cmd.Run()
		case "darwin":
			cmd := exec.Command(`open`, folder)
			cmd.Run()
		}
		Start()
	case "Manage Settings":
		Settings()
	case "Get help":
		browser.OpenURL("https://google.com")
		Start()
	case "Quit":
		os.Exit(1)
	}
}
