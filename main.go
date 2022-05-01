package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	err := aplicacao.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
