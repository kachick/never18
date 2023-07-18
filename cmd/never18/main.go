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
	versionFlag := flag.Bool("version", false, "print the version of this program")
	birthDateFlag := flag.String("birth", "", "your birthday - first time mewling and puking in this world")
	momentDateFlag := flag.String("moment", "", "Time flies like an arrow")
	dateFormatFlag := flag.String("dateFormat", time.DateOnly, "Go date format for parsing")
	limitFlag := flag.Int("limit", 17, "the limit of truth age")
	doctorFlag := flag.Bool("doctor", false, "print the truth and falsehood age for debugging")

	const usage = `Usage: never18 --birth=[YOUR_BIRTH_DAY] <flags>

	$ never18 --birth 1962-08-07
	$ never18 --birth 1962-08-07 --limit 13
	$ never18 --birth 1962-08-07 --moment 2112-09-03
	$ never18 --birth 1962-08-07 --doctor
	$ never18 --version`

	flag.Usage = func() {
		// https://github.com/golang/go/issues/57059#issuecomment-1336036866
		fmt.Printf("%s", usage+"\n\n")
		fmt.Println("Usage of command:")
		flag.PrintDefaults()
	}

	flag.Parse()
	if *versionFlag {
		revision := commit[:7]
		fmt.Printf("%s\n", "never18"+" "+version+" "+"("+revision+")")
		return
	}

	if *birthDateFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	birth, err := time.Parse(*dateFormatFlag, *birthDateFlag)
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

	if *doctorFlag {
		fmt.Printf("TruthAge: %v\n", age.Truth(moment, *limitFlag))
		fmt.Printf("FalsehoodAge: %v\n", age.Falsehood(moment))
	}

	fmt.Println(age.Truth(moment, *limitFlag))
}
