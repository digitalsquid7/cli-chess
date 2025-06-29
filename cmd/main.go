package main

import (
	"fmt"
	"github.com/digitalsquid7/cli-chess/internal/app"
)

func main() {
	application := app.New()
	if err := application.Run(); err != nil {
		fmt.Print(err.Error())
	}
}
