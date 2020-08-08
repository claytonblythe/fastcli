package main

import (
	_ "net/http/pprof"

	fastcli "github.com/claytonblythe/fastcli/fastcli"
)

func main() {
	fastcli.Test_Speed()
}
