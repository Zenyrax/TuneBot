package plog

import (
	"fmt"
	s "strings"
	"time"

	"github.com/gookit/color"
)

func TaskStatus(count int, c, status, message string) {
	t := time.Now()
	switch c {
	case "red":
		color.Red.Printf("[%s][Task %d][%s] %s\n", t.Format("2006-01-02 15:04:05"), count, s.ToUpper(status), message)
	case "green":
		color.Green.Printf("[%s][Task %d][%s] %s\n", t.Format("2006-01-02 15:04:05"), count, s.ToUpper(status), message)
	case "yellow":
		color.Yellow.Printf("[%s][Task %d][%s] %s\n", t.Format("2006-01-02 15:04:05"), count, s.ToUpper(status), message)
	case "blue":
		color.Cyan.Printf("[%s][Task %d][%s] %s\n", t.Format("2006-01-02 15:04:05"), count, s.ToUpper(status), message)
	case "purple":
		clr := color.RGB(150, 103, 153)
		clr.Printf("[%s][Task %d][%s] %s\n", t.Format("2006-01-02 15:04:05"), count, s.ToUpper(status), message)
	default:
		fmt.Printf("[%s][Task %d][%s] %s\n", t.Format("2006-01-02 15:04:05"), count, s.ToUpper(status), message)
	}
}
