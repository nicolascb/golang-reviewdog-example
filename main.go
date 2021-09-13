package main

import (
	"fmt"
	"os"

	"github.com/nicolascb/golang-reviewdog-example/app"
)

func main() {
	fmt.Println("vim-go")

	a := app.NewApp(os.Getenv("VERSION"))
	fmt.Println(a.Version())
}
