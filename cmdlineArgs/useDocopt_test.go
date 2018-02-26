package cmdlineArgs

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"testing"
)

const usage = `Naval Fate.

Usage:
	naval_fate ship new <name>...
	naval_fate ship <name> move <x> <y> [--speed=<kn>]
	naval_fate ship shoot <x> <y>
	naval_fate mine (set|remove) <x> <y> [--moored|--drifting]
	naval_fate -h | --help
	naval_fate --version

Options:
	-h --help     Show this screen.
	--version     Show version.
	--speed=<kn>  Speed in knots [default: 10].
	--moored      Moored (anchored) mine.
	--drifting    Drifting mine.`

const _usage = `Flamebird Command Line Utility

Flamebird is a collection of utilities that augment WtPhoenix.

Its core concept is based on an existing devOps tool that periodically downloads crash logs from junkbox.wetafx.co.nz, 
turning them into json documents then uploading them to elasticsearch.

A <configname> is required to determine how to generate the json documents (e.g. how to transform the information in 
the crash log to some search-friendly key-value pairs)

Flamebird also provides a client tool to search through these documents stored on elasticsearch. One needs to set the 
command to "search" to invoke this feature. It too requires <configname> as different search rules may be implemented 
per configuration.

Usage:
	flamebird <configname> download [--fromtime=<fromtimestamp>] [--totime=<totimestamp>]
	flamebird <configname> search [--fromtime=<fromtimestamp>] [--totime=<totimestamp>] [--func=<funcpattern>] 
[--relevancy=<relevancy>] [--fields=<fields>]

Options:
	-h --help     Show this screen.
	--version     Show version.
	--fromtime=<fromtimestamp> sets the search/download range
	--totime=<totimestamp> sets the search/download range
	--func=<funcpattern> a pattern to select only the records with matching function name
	--relevancy=<relevancy> a number (0-9) to select only the records with equal or higher relevancy
	--fields=<fields> to specify which field(s) to display. By default all fields are shown.
`

func prettyPrint(opts docopt.Opts) {
	for k, v := range opts {
		fmt.Println(k, v)
	}
}

func TestExpectCommand(t *testing.T) {
	arguments, err := docopt.ParseArgs(usage, []string{"ship", "new", "asd"}, "")
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, arguments["ship"].(bool))
}

func TestExpectCommandFB(t *testing.T) {
	arguments, err := docopt.ParseArgs(_usage, []string{"koru", "download"}, "")
	if err != nil {
		t.Fatal(err)
	}
	assertTrue(t, arguments["download"].(bool))
	assertStringEqual(t, "koru", arguments["<configname>"].(string))
}

func TestExpectPositionalArgValue(t *testing.T) {
	arguments, err := docopt.ParseArgs(usage, []string{"ship", "new", "asd"}, "")
	if err != nil {
		t.Fatal(err)
	}
	assertStringEqual(t, "asd", arguments["<name>"].([]string)[0])
}

func TestExpectPositionalArgValueFB(t *testing.T) {
	arguments, err := docopt.ParseArgs(_usage, []string{"koru", "download"}, "")
	if err != nil {
		t.Fatal(err)
	}
	assertStringEqual(t, "koru", arguments["<configname>"].(string))
}

func TestExpectDefaultArgValue(t *testing.T) {
	arguments, err := docopt.ParseArgs(usage, []string{"ship", "new", "asd"}, "")
	if err != nil {
		t.Fatal(err)
	}
	assertStringEqual(t, "10", arguments["--speed"].(string))
}

func TestExpectOptions(t *testing.T) {
	arguments, err := docopt.ParseArgs(_usage, []string{"koru", "download", "--fromtime=2018_01_02-00:00:00"}, "")
	if err != nil {
		t.Fatal(err)
	}
	assertStringEqual(t, "2018_01_02-00:00:00", arguments["--fromtime"].(string))
	v, _ := arguments["--totime"]
	assertEqual(t, nil, v)
}
