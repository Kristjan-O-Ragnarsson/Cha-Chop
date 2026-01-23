package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	numberOfRuns := flag.Int("r", 1000, "Number of runs of the binary excluding warm-up")
	binaryName := flag.String("binary", "ICIS_Encryption", "Name of the binary to run")
	warmUp := flag.Bool("no-warm-up", false, "Whether to warm-up the binary")

	path := flag.String("path", "/tmp", "Path to the binary")
	flag.Parse()
	dryRuns := 0
	if !*warmUp {
		dryRuns = *numberOfRuns / -20
	}

	binaryPath := fmt.Sprintf("%s/%s", *path, *binaryName)
	var inputs = [4]string{ // TODO some modular/configurable way to do input
		fmt.Sprintf("%s/text/28.txt", *path),
		fmt.Sprintf("%s/text/56.txt", *path),
		fmt.Sprintf("%s/text/112.txt", *path),
		fmt.Sprintf("%s/text/224.txt", *path),
	}

	allOutputs := make(map[string]map[int][]string)
	for _, input := range inputs {
		allOutputs[input] = make(map[int][]string)
		for i := dryRuns; i < *numberOfRuns; i++ {
			cmd := exec.Command(binaryPath)
			cmd.Dir = *path

			inputFile, err := os.Open(input)
			if err != nil {
				log.Fatalf("Run %d: Failed to open input: %v", i, err)
			}

			cmd.Stdin = inputFile

			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("ERORR at Run %d: %v\n %s", i, err, string(out))

			} else {
				//println(string(out))
				if i >= 0 {
					allOutputs[input][i] = strings.Split(strings.TrimRight(string(out), "\r\n"), ",")
				}
			}

			err = inputFile.Close()
			if err != nil {
				log.Fatalf("Run %d: Failed to close input: %v", i, err)
			}
		}
	}

	for i, input := range inputs {
		file, err := os.Create(fmt.Sprintf("result_%d.csv", i))
		if err != nil {
			log.Fatal(err)
		}

		writer := csv.NewWriter(file)
		for _, output := range allOutputs[input] {
			extraData := output

			err := writer.Write(extraData)
			if err != nil {
				log.Printf("Failed to write iteration %d to output file: %v", i, err)
			}

		}
		writer.Flush()
		err = file.Close()
		if err != nil {
			log.Printf("Writer %s: Failed to close output: %v", input, err)
		}
	}
}
