package appdata

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"tunebot/modules/tuneportals"
	"tunebot/util/plog"
	"tunebot/util/proxies"

	"github.com/kirsle/configdir"
)

func ListTasks() []string {
	tasks := make([]string, 0)

	configPath := configdir.LocalConfig("TuneBot")
	tasksPath := filepath.Join(configPath, "tasks")

	files, err := ioutil.ReadDir(tasksPath)
	if err != nil {
		log.Println("Sorry, I couldn't read your tasks folder. Please try running me as an administrator.")
		log.Println(err)
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".csv") {
			path := filepath.Join(tasksPath, f.Name())

			dat, err := os.ReadFile(path)
			if err != nil {
				log.Println("Sorry, I couldn't read your tasks folder. Please try running me as an administrator.")
				log.Println(err)
				time.Sleep(10 * time.Second)
				os.Exit(1)
			}
			lines := strings.Split(string(dat), "\n")
			count := len(lines)

			count--

			s := "s"
			if count == 1 {
				s = ""
			}

			tasks = append(tasks, fmt.Sprintf("%s - %d task%s", f.Name(), count, s))
		}
	}

	tasks = append(tasks, "Return")

	return tasks
}

func StartTasks(tasks, proxyList string, restart int, currentRestart *int) {
	configPath := configdir.LocalConfig("TuneBot")
	tasksPath := filepath.Join(configPath, "tasks")
	path := filepath.Join(tasksPath, tasks)

	dat, err := os.ReadFile(path)
	if err != nil {
		log.Println("Sorry, I couldn't read your tasks file. Please try running me as an administrator.")
		log.Println(err)
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}

	lines := strings.Split(string(dat), "\n")

	proxiesPath := filepath.Join(configPath, "proxies")
	path = filepath.Join(proxiesPath, proxyList)

	dat, err = os.ReadFile(path)
	if err != nil {
		log.Println("Sorry, I couldn't read your proxies file. Please try running me as an administrator.")
		log.Println(err)
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}

	readyList := proxies.NewList(string(dat))

	useProxy := true

	if len(readyList.Proxies) == 0 {
		useProxy = false
	}

	var wg sync.WaitGroup

	wg.Add(1)

	for i := 1; i < len(lines); i++ {
		fields := strings.Split(lines[i], ",")

		price, err := strconv.ParseFloat(fields[4], 64)
		if err != nil {
			plog.TaskStatus(i, "red", "Fatal error", "Price must be a number")
			continue
		}

		task := &tuneportals.Task{
			Site:           strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fields[0], "/", ""), "https:", ""), "http:", ""), // We do a little nesting
			UPC:            fields[1],
			Quantity:       fields[2],
			ShippingOption: fields[3],
			PriceLimit:     price,

			FirstName:    fields[5],
			LastName:     fields[6],
			Email:        fields[7],
			Phone:        fields[8],
			AddressLine1: fields[9],
			AddressLine2: fields[10],
			City:         fields[11],
			State:        fields[12],
			Zip:          fields[13],
			Country:      fields[14],
			CardNumber:   fields[15],
			CardMonth:    fields[16],
			CardYear:     fields[17],
			CVC:          fields[18],

			Count:    i,
			UseProxy: useProxy,
			Stage:    tuneportals.INIT,
			Proxies:  readyList,
			Webhook:  CurrentConfig.Webhook,
		}

		go tuneportals.Start(task)
	}

	wg.Wait()
}
