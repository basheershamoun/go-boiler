package main

import (
	"main/app"
	"main/cli"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "web")
	}
	mode := os.Args[1]
	switch mode {
	case "cli":
		cli.Main()
	case "web":
		app.Main()
	default:
		app.Main()

	}
}
