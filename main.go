package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/claytonblythe/fast_cli/fast_cli"
)

func main() {
	fmt.Println("Success!")
	fast_cli.Get_urls()
}
