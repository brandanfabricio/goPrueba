package main

import (
	"flag"
	"log"
	"strconv"
	"task/commands"
)

func main() {

	var expens []float32
	var export string
	flag.StringVar(&export, "export", "", "exportar detalle en .txt")
	flag.Parse()

	for {

		input, err := commands.GetInput()

		if err != nil {
			log.Fatal(err)
		}

		if input == "cls" {
			break
		}
		expen, err := strconv.ParseFloat(input, 32)
		if err != nil {
			// log.Fatal(err)
			continue
		}
		expens = append(expens, float32(expen))
	}

	if export == "" {

		commands.ShowConsole(expens)

	} else {
		commands.Export(export, expens)
	}

	log.Println("closing...")
}
