package main

import (
	"fmt"
	"go-compiler/repl"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s. Feel free to type in commands\n", currentUser.Username)
	repl.Start(os.Stdin, os.Stdout)
}
