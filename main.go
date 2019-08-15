package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type delimeter string

var (
	newline delimeter = "\n"
	space   delimeter = " "
	comma   delimeter = ","
	tab     delimeter = "\t"
)

// parse detects the format of the input string and returns a slice of values
func parse(text string) []string {
	vals := []string{}

	switch {
	case strings.Contains(text, "\n"):
		vals = strings.Split(text, "\n")
	case strings.Contains(text, ","):
		vals = strings.Split(text, ",")
	case strings.Contains(text, "\t"):
		vals = strings.Split(text, "\t")
	default:
		vals = strings.Fields(text)
	}
	out := []string{}
	for _, v := range vals {
		out = append(out, strings.TrimSpace(v))
	}
	return out
}

func stdin() string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return ""
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return ""
	}

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(buf))
}

func main() {
	n := flag.Bool("n", false, "newline")
	s := flag.Bool("s", false, "space")
	c := flag.Bool("c", false, "comma")
	t := flag.Bool("t", false, "tab")
	flag.Parse()

	list := stdin()
	if list == "" {
		fmt.Fprintln(os.Stderr, "stdin cannot be empty")
		os.Exit(1)
	}

	var outputDelim delimeter

	if *n {
		outputDelim = newline
	}
	if *s {
		outputDelim = space
	}
	if *c {
		outputDelim = comma
	}
	if *t {
		outputDelim = tab
	}

	vals := parse(list)
	fmt.Println(strings.Join(vals, string(outputDelim)))
}
