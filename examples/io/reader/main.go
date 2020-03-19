package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 14)

	fmt.Println("------ LIMIT READER ------")
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	// ------------------------------------------------

	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	multi := io.MultiReader(r1, r2, r3)

	fmt.Println("------- MULTI READER ------")
	if _, err := io.Copy(os.Stdout, multi); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
