package appdata

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kirsle/configdir"
)

func ListProxies() []string {
	proxies := make([]string, 0)

	configPath := configdir.LocalConfig("TuneBot")
	tasksPath := filepath.Join(configPath, "proxies")

	files, err := ioutil.ReadDir(tasksPath)
	if err != nil {
		log.Println("Sorry, I couldn't read your proxies folder. Please try running me as an administrator.")
		log.Println(err)
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			path := filepath.Join(tasksPath, f.Name())

			dat, err := os.ReadFile(path)
			if err != nil {
				log.Println("Sorry, I couldn't read your proxies folder. Please try running me as an administrator.")
				log.Println(err)
				time.Sleep(10 * time.Second)
				os.Exit(1)
			}
			lines := strings.Split(string(dat), "\n")
			count := len(lines)

			// fmt.Println(lines)

			for i := 0; i < len(lines); i++ {
				if len(lines[i]) == 0 {
					count--
				}
			}

			s := "proxies"
			if count == 1 {
				s = "proxy"
			}

			proxies = append(proxies, fmt.Sprintf("%s - %d %s", f.Name(), count, s))
		}
	}

	proxies = append(proxies, "Return")

	return proxies
}
