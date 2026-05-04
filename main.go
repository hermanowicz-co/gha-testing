package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world from blacksmith, go version. current time is:", time.Now().Format(time.RFC3339))
	fmt.Println("Current time in unix format is:", time.Now().Unix())
	fmt.Println("Running init test of review agent on aws.")
}
