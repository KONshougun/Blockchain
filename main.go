package main

import (
	"fmt"

	"github.com/KONshougun/Blockchain/storage"
)

func main() {
	if err := storage.VerifyStorage(); err != nil{
		fmt.Printf("err: %v\n", err)
	}else{
		fmt.Println("Blocco genesis aggiunto con successo")
	}
}
