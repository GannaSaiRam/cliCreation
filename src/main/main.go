package main

import (
	"fmt"
	"flag"
	"errors"

	// Internal Packages
	"header"
	log "logger"
)

func showHelp() {
	fmt.Print(`
Usage: CLI Template [OPTIONS]

Options:
	-s, --string	Prints string input.	(default='some string').
	-i, --int		Prints int.				(default=0).
	-e, --error		Prints a custom made error.
	-w, --warning	Prints the warning you entered.
	-h, --help		Prints this help info.
`)
}

func setFlag(flag *flag.FlagSet) {
	flag.Usage = func() {
		showHelp()
	}
}

func main() {
	var (
		sStr string
		sHelp bool
		sInt int
		sErr bool
	)
	flag.StringVar(&sStr, "s", "", "")
	flag.StringVar(&sStr, "string", "", "")
	flag.BoolVar(&sHelp, "h", false, "")
	flag.BoolVar(&sHelp, "help", false, "")
	flag.IntVar(&sInt, "i", 0, "")
	flag.IntVar(&sInt, "int", 0, "")
	flag.BoolVar(&sErr, "e", false, "")
	flag.BoolVar(&sErr, "error", false, "")

	setFlag(flag.CommandLine)
	flag.Parse()
	if sHelp {
		showHelp()
		return
	}
	header.GetHeader()
	log.InitLogger()

	if sStr != "" {
		log.Printf("This is user string input: %s\n", sStr)
	}

	if sInt != 0 {
		log.Printf("This is user integer input: %d\n", sInt)
	}

	if sErr {
		log.Errorf("Oh! Got some error %s", errors.New("Something broke"))
	}
}
