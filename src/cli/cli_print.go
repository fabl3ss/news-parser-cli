package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func PrintTitle() {
	if header, err := os.Open("cli/header.txt"); err != nil {
		fmt.Println("Error to load header! ", err)
		log.Fatal(err)
	} else {
		text, _ := ioutil.ReadAll(header)
		fmt.Print(string(text))
	}
}
