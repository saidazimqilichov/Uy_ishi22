package main

import (
	"UY_ISHI22/git"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("/home/saidazim/NT_Homeworks/Uy_ishi22/storage/username.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	name, err := git.GetUserName()
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(name)

	email, err := git.GetUserEmail()
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(email)

	file.Close()

}
