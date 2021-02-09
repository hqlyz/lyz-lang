package main

import (
	"fmt"
	"log"
	"lyz-lang/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Hello %s! This is the LYZ programming language!\n", u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
