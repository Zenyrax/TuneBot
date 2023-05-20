package appdata

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kirsle/configdir"
)

func Init() bool {
	first := false
	// Sets appdata path
	configPath := configdir.LocalConfig("TuneBot")

	// Ensures it exists
	err := configdir.MakePath(configPath)
	if err != nil {
		log.Println("Sorry, I couldn't create your appdata folder. Please try running me as an administrator.")
		log.Println(err)
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}

	// Set config path
	configFile := filepath.Join(configPath, "config.json")

	// Does the file not exist?
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		first = true
		// Create the new config file.
		config := Config{}
		fh, err := os.Create(configFile)
		if err != nil {
			log.Println("Sorry, I couldn't create your config file. Please try running me as an administrator.")
			log.Println(err)
			time.Sleep(10 * time.Second)
			os.Exit(1)
		}
		defer fh.Close()

		encoder := json.NewEncoder(fh)
		encoder.Encode(&config)
	} else {
		fh, err := os.Open(configFile)
		if err != nil {
			log.Println("Sorry, I couldn't read your config file. Please try running me as an administrator.")
			log.Println(err)
			time.Sleep(10 * time.Second)
			os.Exit(1)
		}
		defer fh.Close()

		decoder := json.NewDecoder(fh)
		decoder.Decode(&CurrentConfig)
	}

	// Set proxies path
	proxiesFolder := filepath.Join(configPath, "proxies")
	if _, err := os.Stat(proxiesFolder); os.IsNotExist(err) {
		// Make proxies folder
		errDir := os.MkdirAll(proxiesFolder, 0755)
		if errDir != nil {
			log.Println("Sorry, I couldn't create your proxies folder. Please try running me as an administrator.")
			log.Println(err)
			time.Sleep(10 * time.Second)
			os.Exit(1)
		}

		examplePath := filepath.Join(proxiesFolder, "proxies.txt")

		file, err := os.Create(examplePath)
		if err != nil {
			log.Println("Sorry, I couldn't create your proxies files. Please try running me as an administrator.")
			log.Println(err)
			time.Sleep(10 * time.Second)
			os.Exit(1)
		}
		file.Close()
	}

	// Set tasks path
	tasksFolder := filepath.Join(configPath, "tasks")
	if _, err := os.Stat(tasksFolder); os.IsNotExist(err) {
		// Make tasks folder
		errDir := os.MkdirAll(tasksFolder, 0755)
		if errDir != nil {
			log.Println("Sorry, I couldn't create your proxies folder. Please try running me as an administrator.")
			log.Println(err)
			time.Sleep(10 * time.Second)
			os.Exit(1)
		}

		examplePath := filepath.Join(tasksFolder, "tasks.csv")

		d1 := []byte("STORE,UPC,QUANTITY,SHIPPING OPTION,PRICE LIMIT,FIRST NAME,LAST NAME,EMAIL,PHONE NUMBER,ADDRESS LINE 1,ADDRESS LINE 2,CITY,STATE,POSTCODE / ZIP,COUNTRY,CARD NUMBER,EXPIRE MONTH,EXPIRE YEAR,CARD CVC")
		d2 := []byte("\nvintagevinyl.com,196587792114,1,standard,49.99,John,Doe,example@tunebot.io,3334443333,123 Mulberry St,Suite 2,New York City,NY,34734,US,4242424242424242,06,27,235")
		err = os.WriteFile(examplePath, append(d1, d2...), 0644)
		if err != nil {
			log.Println("Sorry, I couldn't create your example task file. Please try running me as an administrator.")
			log.Println(err)
			time.Sleep(10 * time.Second)
			os.Exit(1)
		}
	}

	return first
}
