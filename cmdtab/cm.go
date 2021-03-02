package cmdtab

import (
	"html/template"
	"os"

	"github.com/rwxrob/cmdtab"
)

type KN struct {
	Imports  []string
	Commands []string
	Version  string
	Git      string
}

var (
	tpl  *template.Template
	cmds = []string{
		"kn",
		"day",
		"cd",
		"add",
		"tstamp",
		"html",
		"epoch",
		"pomo",
	}

	imports = []string{
		"github.com/rwxrob/cmdtab-pomo",
	}

	// TODO: Determine if there is a way to allow `` instead of ""
	knFile = `package cmdtab
import (
	"github.com/rwxrob/cmdtab"
{{ range .Imports }}
	_ "{{ . }}"
{{ end }}
)

func Execute(s string) {
	//      cmdtab.Dump(cmdtab.Index)
	cmdtab.Execute(s)
}

func init() {
	x := cmdtab.New(
{{ range .Commands }}
	"{{ . }}",
{{ end }}
	)
	x.Summary = ""
	x.Version = "{{ .Version }}"
	x.Author = "Rob Muhlestein <rob@rwx.gg> (rob.rwx.gg)"
	x.Git = "{{ .Git }}"
	x.Copyright = "(c) Rob Muhlestein"
	x.License = "MPL-v2.0"
	x.Description = ""
}`
)

func init() {
	x := cmdtab.New("cm")
	x.Version = `v1.0.0`
	x.Usage = `([<subcommand name> <subcommand link>])`
	x.Summary = `install and manage cmdtab-based subcommands`
	x.Description = ``
	x.Examples = `
		cm links github.com/oglinuk/cmdtab-links
	`
	x.Method = func(args []string) error {
		// We expect <subcommand name> and <subcommand link>
		if len(args) < 1 {
			return x.UsageError()
		}

		// TODO: Parse template rather than just append, otherwise
		// it will only ever save one command.
		cmds = append(cmds, args[0])
		imports = append(imports, args[1])

		// TODO: This will error if not in the kn directory, todo better ...
		f, err := os.Create("cmdtab/kn.go")
		if err != nil {
			return err
		}
		defer f.Close()

		tpl, err := template.New("kn.go").Parse(knFile)
		if err != nil {
			return err
		}

		// TODO: Generate kn.go using KN configuration
		tpl.Execute(f, &KN{
			Imports:  imports,
			Commands: cmds,
			Version:  "v1.0.0",
			Git:      "github.com/rwxrob/kn",
		})

		// go install
		// complete -C kn kn
		return nil
	}
}
