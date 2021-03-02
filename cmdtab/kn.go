package cmdtab

import (
	"github.com/rwxrob/cmdtab"
	_ "github.com/rwxrob/cmdtab-pomo"
)

func Execute(s string) {
	//      cmdtab.Dump(cmdtab.Index)
	cmdtab.Execute(s)
}

func init() {
	x := cmdtab.New("kn", "day", "cd", "cm", "add", "tstamp", "html", "epoch", "pomo")
	x.Summary = ""
	x.Version = "v1.0.0"
	x.Author = "Rob Muhlestein <rob@rwx.gg> (rob.rwx.gg)"
	x.Git = "https://github.com/rwxrob/kn"
	x.Copyright = "(c) Rob Muhlestein"
	x.License = "MPL-v2.0"
	x.Description = ""
}
