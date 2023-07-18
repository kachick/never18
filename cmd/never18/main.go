package main

import (
	"flag"
	"fmt"
	"log"
	"never18"
	"os"
	"time"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	birthDateFlag := flag.String("birth", "", "your birthday - first time mewling and puking in this world")
	momentDateFlag := flag.String("moment", "", "time flies like an arrow")
	dateFormatFlag := flag.String("dateFormat", time.DateOnly, "Go date format for parsing")
	limitFlag := flag.Int("limit", 17, "you believe that the number is the truth or your own soul")
	versionFlag := flag.Bool("version", false, "print the version of this program")
	doctorFlag := flag.Bool("doctor", false, "print information to stderr for bug reporting")

	const usage = `Usage: never18 --birth=[YOUR_BIRTH_DAY] <flags>

	$ never18 --birth 1962-08-07
	$ never18 --birth 1962-08-07 --limit 12
	$ never18 --birth 1962-08-07 --moment 2112-09-03
	$ never18 --birth 1962-08-07 --doctor
	$ never18 --version`

	flag.Usage = func() {
		// https://github.com/golang/go/issues/57059#issuecomment-1336036866
		fmt.Printf("%s", usage+"\n\n")
		fmt.Println("Usage of command:")
		flag.PrintDefaults()
	}

	revision := commit[:7]
	version := fmt.Sprintf("%s\n", "never18"+" "+version+" "+"("+revision+")")

	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		return
	}

	if *birthDateFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	birth, err := time.ParseInLocation(*dateFormatFlag, *birthDateFlag, time.Local)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var moment time.Time
	if *momentDateFlag == "" {
		moment = time.Now()
	} else {
		moment, err = time.Parse(*dateFormatFlag, *momentDateFlag)
		if err != nil {
			log.Fatalf("%v", err)
			flag.Usage()
			os.Exit(1)
		}
	}

	age := never18.Age{
		Birth: birth,
	}

	truth, err := age.Truth(moment, *limitFlag)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if *doctorFlag {
		nominally, err := age.Nominally(moment)
		if err != nil {
			log.Fatalf("%v", err)
		}

		fmt.Fprintln(os.Stderr, version)
		fmt.Fprintf(os.Stderr, "birth: %v, moment: %v, limit: %v\n", birth, moment, *limitFlag)
		fmt.Fprintf(os.Stderr, "TruthAge: %v\n", truth)
		fmt.Fprintf(os.Stderr, "NominallyAge: %v\n", nominally)
	}

	fmt.Println(truth)
}
