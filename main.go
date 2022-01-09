package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {

	newFileName := "/heat1.html"

	filePath := flag.String("file", "", "path to the html file with the iracing results")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("please pass required -file command line option")
		os.Exit(1)
	}

	fmt.Printf("Processing %s...\n", *filePath)

	inputFile, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Println("Encountered error reading file", *filePath)
		os.Exit(1)
	}

	inputFileString := string(inputFile)

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Encountered error getting current directory")
		os.Exit(1)
	}

	outputFile, err := os.Create(workingDir + newFileName)
	if err != nil {
		fmt.Println("Encountered error creating file")
		os.Exit(1)
	}

	defer outputFile.Close()

	outputString := ""

	resultOBJ := regexp.MustCompile(`var resultOBJ ={(\n.+)+FEATURE",(\n.+)+};`)

	outputString = resultOBJ.ReplaceAllString(inputFileString, "")

	heat1 := regexp.MustCompile(`HEAT 1`)

	outputString = heat1.ReplaceAllString(outputString, "FEATURE")

	simSessNum := regexp.MustCompile(`simSessNum:-2`)

	outputString = simSessNum.ReplaceAllString(outputString, "simSessNum:0")

	_, err = outputFile.WriteString(outputString)
	if err != nil {
		fmt.Println("Encountered error writing to file")
		os.Exit(1)
	}
	fmt.Println("Created " + newFileName)
}
