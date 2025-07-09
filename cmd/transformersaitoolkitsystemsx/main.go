// cmd/transformersaitoolkitsystemsx/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"transformersaitoolkitsystemsx/internal/transformersaitoolkitsystemsx"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := transformersaitoolkitsystemsx.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
