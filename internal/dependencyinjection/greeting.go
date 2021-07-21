package dependencyinjection

import (
	"fmt"
	"io"
)

func Greet(b io.Writer, name string) {
	// b.Write([]byte("Hello, " + name))
	fmt.Fprintf(b, "Hello, %s.\n", name)
}
