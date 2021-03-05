package cmdtab

import (
	"fmt"
	"os/exec"
	"runtime"
	"github.com/rwxrob/cmdtab"
)

var (
	err error
)

func init() {
	x := cmdtab.New("open")
	x.Version = `v1.0.0`
	x.Usage = `[<resource>]`
	x.Summary = `opens <resource> with the default browser`
	x.Description = `The *open* subcommand opens args[0] using the detected
	platforms default browser. In this context a resource is any valid URL
	or file including but not limited to: .md, .html, .txt, .mp4, .mp3, ...`
	x.Method = func(args []string) error {
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", args[0]).Run()
		case "windows":
			err = exec.Command("open", args[0]).Run()
		default:
			// I dont support apple(darwin)
			err = fmt.Errorf("platform not supported ...\n")
		}
		return err
	}
}

