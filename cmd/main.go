package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/dependencyinjection"
	"github.com/joshuaswickirl/learn-go-with-tests/internal/hello"
	"github.com/joshuaswickirl/learn-go-with-tests/internal/mocking"
)

func main() {
	fmt.Println(hello.Hello("Josh", ""))

	dependencyinjection.Greet(os.Stdout, "Eoldie")

	mocking.Countdown(os.Stdout, &mocking.ConfigurableSleeper{
		Duration: 1 * time.Second, Sleeper: time.Sleep,
	})
}
