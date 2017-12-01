package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(0)
	root := flag.String("root", ".", "the root directory in which the corpus directory will reside")
	flag.Parse()

	corpusDir := filepath.Join(*root, "corpus")
	if err := os.MkdirAll(corpusDir, 0755); err != nil {
		log.Fatal(err)
	}

	corpa := [][]byte{
		nil,
		{},
		[]byte(" "),
		[]byte("*"),
		{0x01, 0x01},
		bytes.Repeat([]byte("^ "), 100),
		bytes.Repeat([]byte("!aab^"), 100000),
		bytes.Repeat([]byte("   "), 10000),
	}

	for i, corpus := range corpa {
		outPath := filepath.Join(corpusDir, fmt.Sprintf("initial-corpus-%d", i))
		f, err := os.Create(outPath)
		if err == nil {
			f.Write(corpus)
			f.Close()
		} else {
			log.Printf("Failed to generate %q: %v", outPath, err)
		}
	}
}
