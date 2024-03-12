package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Monkey Programming Language!\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
