package main

import (
	"fmt"
	"os"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/dependencyinjection"
	"github.com/joshuaswickirl/learn-go-with-tests/internal/hello"
)

func main() {
	fmt.Println(hello.Hello("Josh", ""))

	dependencyinjection.Greet(os.Stdout, "Eoldie")
}
