package main

import (
	"compress/gzip"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	extractFlag := flag.Bool("x", false, "extract archive")
	compressFlag := flag.Bool("c", false, "compress file")
	flag.Parse()

	args := flag.Args()

	for _, file_name := range args {
		if *extractFlag && *compressFlag {
			log.Fatal("could not parse flags")
		} else if *extractFlag {
			extract(file_name)
		} else if *compressFlag {
			compress(file_name)
		} else {
			flag.PrintDefaults()
		}
	}
}

func compress(file_name string) {
	// get file absolute path
	absPath, err := filepath.Abs(file_name)
	if err != nil {
		log.Fatal(err)
	}

	// open file
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create compressed file
	absPath += ".gz"
	compressedFile, err := os.Create(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer compressedFile.Close()

	// use gzip algorithm
	compressedWriter := gzip.NewWriter(compressedFile)
	defer compressedWriter.Close()

	if _, err := io.Copy(compressedWriter, file); err != nil {
		log.Fatal(err)
	}
}

func extract(file_name string) {
	// get file absolute path
	absPath, err := filepath.Abs(file_name)
	if err != nil {
		log.Fatal(err)
	}

	// open file
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// use gzip algorithm
	extractReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	absPath = strings.TrimSuffix(absPath, ".gz")
	extractedFile, err := os.Create(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer extractedFile.Close()

	if _, err := io.Copy(extractedFile, extractReader); err != nil {
		log.Fatal(err)
	}
}
