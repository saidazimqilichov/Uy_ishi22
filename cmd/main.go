package main

import (
	"UY_ISHI22/git"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("storich/username.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	name, err := git.GetUserName()
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(name)

	file.Close()

}
