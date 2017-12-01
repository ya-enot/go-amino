package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/tendermint/go-wire"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("want <filepath>")
	}

	inFile := os.Args[1]
	in, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatalf("reading: %q err: %v", inFile, err)
	}
	var n int
	buf := new(bytes.Buffer)
	wire.WriteByteSlice(in, buf, &n, &err)
	if err != nil {
		log.Fatal(err)
	}

	n = 0
	// E2E
	out := wire.ReadByteSlice(buf, n, &n, &err)
	if !bytes.Equal(out, in) {
		log.Fatalf("in:  %x\nout: %x", in, out)
	}
}
