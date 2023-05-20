package updates

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

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
