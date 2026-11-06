package main

import (
	"errors"
	"fmt"
)

// Sentinel Error
var err1 = errors.New("player not found")
var err2 = fmt.Errorf("player %d not found", 123)

func main() {
	println(err1)
	println(err2)

	// error比較
	if errors.Is(err1, err1) {
		fmt.Println("error!! err1")
	}
}
