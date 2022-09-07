package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func getDiff(previousFileLines []string, NewFileLines []string) []string {
	var diff []string

	for _, newLine := range NewFileLines {
		existsInOldFile := false

		for _, oldLine := range previousFileLines {
			if oldLine == newLine {
				existsInOldFile = true
			}
		}

		if !existsInOldFile {
			diff = append(diff, newLine)
		}
	}

	return diff
}

func retrieveFileData(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		os.Exit(-1)
	}

	splittedData := strings.Split(string(data), "\n")

	return splittedData
}

func main() {
	var previousFileFlag = flag.String("p", "", "Previous file")
	var newFileFlag = flag.String("n", "", "Next file")
	var outputFlag = flag.String("o", "diff", "Diff file")
	flag.Parse()
	previousFileData := retrieveFileData(*previousFileFlag)
	newFileData := retrieveFileData(*newFileFlag)

	diff := getDiff(previousFileData, newFileData)
	diffString := strings.Join(diff, "\n")

	fmt.Println(diffString)
	os.WriteFile(*outputFlag, []byte(diffString), 0655)
}
