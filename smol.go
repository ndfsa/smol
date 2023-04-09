package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	extract_mode := flag.Bool("ex", false, "set mode to encrypt")
	flag.Parse()

	args := flag.Args()

	for _, file_name := range args {
		if *extract_mode {
			decompress(file_name)
		} else {
			compress(file_name)
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func compress(file_name string) {
	file, err := os.Open(file_name)
	check(err)
	defer file.Close()

	compressed_file, err := os.Create(file_name + ".gz")
	check(err)
	defer compressed_file.Close()

	compressed_writer := gzip.NewWriter(compressed_file)
	defer compressed_writer.Close()
	_, err = io.Copy(compressed_writer, file)
	check(err)

	compressed_writer.Flush()
}

func decompress(file_name string) {
	file, err := os.Open(file_name)
	check(err)
	defer file.Close()

	decompressed_file, err := os.Create(strings.TrimSuffix(file_name, ".gz"))
	check(err)
	defer decompressed_file.Close()

	decompressed_reader, err := gzip.NewReader(file)
	check(err)
	defer decompressed_reader.Close()

	writer := bufio.NewWriter(decompressed_file)
	_, err = io.Copy(decompressed_file, decompressed_reader)
	check(err)

	writer.Flush()
}
