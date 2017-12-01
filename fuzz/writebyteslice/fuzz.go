package writebyteslice

import (
	"bytes"
	"fmt"

	"github.com/tendermint/go-wire"
)

func Fuzz(in []byte) int {
	var n int
	var err error
	buf := new(bytes.Buffer)
	wire.WriteByteSlice(in, buf, &n, &err)

	n = 0
	// E2E
	out := wire.ReadByteSlice(buf, n, &n, &err)
	if !bytes.Equal(out, in) {
		panic(fmt.Errorf("in:  %x\nout: %x", in, out))
	}

	return 1
}
