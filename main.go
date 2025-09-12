package main

import (
	"fmt"

	"github.com/DKeshavarz/sinar/internal/interface/server"
)

func main(){
	server := server.New()

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}