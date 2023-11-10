package main

import (
	"os"
	"time"

	"goWithTests/maths/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}