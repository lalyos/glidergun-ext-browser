package browser

import (
	"os"

	"github.com/gliderlabs/glidergun/extpoints"
	"github.com/skratchdot/open-golang/open"
)

func init() {
	extpoints.Register(new(BrowserCommands), "browser")
}

type BrowserCommands struct {
}

func (b *BrowserCommands) Commands() map[string]extpoints.CommandFunc {
	return map[string]extpoints.CommandFunc{
		"openbrowser": OpenBrowser,
	}
}

func fatal(msg string) {
	println("!!", msg)
	os.Exit(2)
}

func OpenBrowser(args []string) {
	err := open.Start(args[0])
	if err != nil {
		fatal("Can't open browser: '" + err.Error())
	}
}
