package menu

import (
	"strings"
	"tunebot/util/appdata"

	"github.com/AlecAivazis/survey/v2"
)

func Tasks() {
	tasks := appdata.ListTasks()

	taskOption := ""
	prompt := &survey.Select{
		Message: "What CSV do you want to run?",
		Options: tasks,
	}
	survey.AskOne(prompt, &taskOption)

	if taskOption == "Return" {
		Start()
		return
	}

	proxies := appdata.ListProxies()

	proxyOption := ""
	prompt = &survey.Select{
		Message: "What proxy list do you want to use?",
		Options: proxies,
	}
	survey.AskOne(prompt, &proxyOption)

	if proxyOption == "Return" {
		Start()
		return
	}

	currentRestart := 0
	// go appdata.Restart(strings.Split(taskOption, " - ")[0], strings.Split(proxyOption, " - ")[0], &currentRestart)
	appdata.StartTasks(strings.Split(taskOption, " - ")[0], strings.Split(proxyOption, " - ")[0], 0, &currentRestart)
}
