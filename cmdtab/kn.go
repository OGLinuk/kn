package cmdtab
import (
	"github.com/rwxrob/cmdtab"

	_ "github.com/rwxrob/cmdtab-pomo"

	_ "github.com/oglinuk/cmdtab-links"

)

func Execute(s string) {
	//      cmdtab.Dump(cmdtab.Index)
	cmdtab.Execute(s)
}

func init() {
	x := cmdtab.New(

	"kn",

	"day",

	"cd",

	"add",

	"tstamp",

	"html",

	"epoch",

	"pomo",

	"links",

	)
	x.Summary = ""
	x.Version = "v1.0.0"
	x.Author = "Rob Muhlestein <rob@rwx.gg> (rob.rwx.gg)"
	x.Git = "github.com/rwxrob/kn"
	x.Copyright = "(c) Rob Muhlestein"
	x.License = "MPL-v2.0"
	x.Description = ""
}