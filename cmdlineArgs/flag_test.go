package cmdlineArgs

import (
	"testing"
	"flag"
)

func TestExpectDefaultFlagValue(t *testing.T) {
	cmdline := flag.NewFlagSet("", flag.ExitOnError)
	aInt := cmdline.Int("a", 42, "an int")
	cmdline.Parse([]string{""})
	assertEqual(t, 42, *aInt)
}

func TestExpectParsedFlagValue(t *testing.T) {
	cmdline := flag.NewFlagSet("", flag.ExitOnError)
	aInt := cmdline.Int("a", 42, "an int")
	cmdline.Parse([]string{"-a=1337"})
	assertEqual(t, 1337, *aInt)
}

func testExpectInvalidFlagValueIgnored(t *testing.T) {
	cmdline := flag.NewFlagSet("", flag.ExitOnError)
	aInt := cmdline.Int("a", 42, "an int")

	// throws an error: invalid value "13aa37" for flag -a
	cmdline.Parse([]string{"-a=13aa37"})
	assertEqual(t, 42, *aInt)
}
