package main

import (
	start "go-reloaded/pkg/start"
)

func main() {
	start.Start("./sample.txt", "result.txt")
	// start.Test()
}
