package menu

import (
	// "github.com/fatih/color"
	"github.com/inancgumus/screen"

	"github.com/gookit/color"

	"fmt"
	"tunebot/util/appdata"
)

func Reset() {
	// Clears the screen and moves the cursor back to the top left of the screen
	screen.Clear()
	screen.Clear()
	screen.Clear()
	screen.MoveTopLeft()

	// A bunch of stuff for making the cool title thing
	c := color.RGB(140, 103, 153)
	c.Println(" _______               ____        _   ")
	c = color.RGB(150, 103, 153)
	c.Println("|__   __|             |  _ \\      | |  ")
	c = color.RGB(160, 103, 153)
	c.Println("   | |_   _ _ __   ___| |_) | ___ | |_ ")
	c = color.RGB(170, 103, 153)
	c.Println("   | | | | | '_ \\ / _ \\  _ < / _ \\| __|")
	c = color.RGB(180, 103, 153)
	c.Println("   | | |_| | | | |  __/ |_) | (_) | |_ ")
	c = color.RGB(190, 103, 153)
	c.Print("   |_|\\__,_|_| |_|\\___|____/ \\___/ \\__|")

	if appdata.CurrentConfig.Version != "" {
		c = color.RGB(190, 103, 153)
		c.Println(fmt.Sprintf(" Version %s\n", appdata.CurrentConfig.Version))
	} else {
		fmt.Println("\n")
	}
}

// _______               ____        _
//|__   __|             |  _ \\      | |
// 	 | |_   _ _ __   ___| |_) | ___ | |_
// 	 | | | | | '_ \\ / _ \\  _ < / _ \\| __|
// 	 | | |_| | | | |  __/ |_) | (_) | |_
// 	 |_|\\__,_|_| |_|\\___|____/ \\___/ \\__|
