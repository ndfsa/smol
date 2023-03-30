package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
	abs_path, err := filepath.Abs(file_name)
	check(err)

	file, err := os.Open(abs_path)
	defer file.Close()
	check(err)

	data, err := ioutil.ReadAll(bufio.NewReader(file))
	check(err)

	abs_path += ".gz"
	compressed_file, err := os.Create(abs_path)
	defer compressed_file.Close()
	check(err)

	compressed_writer := gzip.NewWriter(compressed_file)
	compressed_writer.Write(data)
	compressed_writer.Close()
}

func decompress(file_name string) {
	abs_path, err := filepath.Abs(file_name)
	check(err)

	file, err := os.Open(abs_path)
	defer file.Close()
	check(err)

	reader, err := gzip.NewReader(file)
	check(err)
	data, err := ioutil.ReadAll(reader)
	check(err)

	abs_path = strings.TrimSuffix(abs_path, ".gz")
	decompressed_file, err := os.Create(abs_path)
	defer decompressed_file.Close()
	check(err)

	writer := bufio.NewWriter(decompressed_file)
	writer.Write(data)
	writer.Flush()
}
