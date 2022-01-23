package main

import (
	"fmt"
	"funscript-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Funscript programming language!\n", user.Username)
	fmt.Printf("Enjoy the ride!\n")
	repl.Start(os.Stdin, os.Stdout)
}
