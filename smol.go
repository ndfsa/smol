package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	extract_mode := flag.Bool("ex", false, "set mode to encrypt")
	flag.Parse()

	args := flag.Args()

	for _, file_name := range args {
		if *extract_mode {
			log.Printf("decompressing %s", file_name)
		} else {
			compress(file_name)
		}
	}
}

func compress(file_name string) {
	abs_path, err := filepath.Abs(file_name)
	if err != nil {
		log.Fatalf("failed finding file with name: %s", file_name)
	}

	file, err := os.Open(abs_path)
	if err != nil {
		log.Fatalf("failed opening file %s", abs_path)
	}

	data, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		log.Fatalf("failed reading file %s", abs_path)
	}

	abs_path += ".gz"
	compressed_file, err := os.Create(abs_path)
	if err != nil {
		log.Fatalf("failed creating compressed file %s", abs_path)
	}

	compressed_writer := gzip.NewWriter(compressed_file)
	compressed_writer.Write(data)
	compressed_writer.Close()

}
