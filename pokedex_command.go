package main

import "fmt"

func commandPokedex(c config, args []string) error {
	fmt.Println(c.pokeDex.List())
	return nil
}
