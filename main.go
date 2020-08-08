package main

import (
	_ "net/http/pprof"

	"github.com/claytonblythe/fast_cli/fast_cli"
)

func main() {
	fast_cli.Get_urls()
}
