package updates

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/inconshreveable/go-update"
)

// Keeping this flexible so I can maybe make a standalone package for it
func Check(owner, repo, version string) {
	res, err := http.DefaultClient.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo))
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		var releases githubResponse
		json.Unmarshal(body, &releases)

		if len(releases) == 0 {
			log.Println("No releases found")
			return
		}

		if releases[0].Name != version {
			log.Printf("A new version of %s is out (%s -> %s)\n", repo, version, releases[0].Name)
			if len(releases[0].Assets) == 0 {
				log.Println("No assets found")
				return
			}

			for i := 0; i < len(releases[0].Assets); i++ {
				if runtime.GOOS == "windows" && releases[0].Assets[i].ContentType == "application/x-msdownload" {
					log.Println("Installing...")
					updateExe(releases[0].Assets[i].BrowserDownloadURL)
					// I used to have code that automatically restarted after updating, but it doesn't seem to work anymore
					// I'll continue to look into this
					log.Printf("Please relaunch %s, this window will close in 10 seconds\n", repo)
					<-time.After(10 * time.Second)
					os.Exit(1)
				}
			}
		}
	}
}

func updateExe(link string) {
	res, err := http.DefaultClient.Get(link)
	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		err := update.Apply(res.Body, update.Options{})
		if err == nil {
			log.Println("Successfully updated executable")
		}
	}
}

// Code below isn't used anymore, but I still think it's neat
func SetVersion() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}

	hash, _ := md5sum(ex)
	if len(hash) > 10 {
		return hash[:6]
	}
	return hash
}

func md5sum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
