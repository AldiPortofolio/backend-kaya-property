package main

import (
	"fmt"

	logger "kaya-backend/library/logger/v2"
)

func main() {
	kayalog := logger.InitLog()

	for i := 0; i < 15; i++ {
		kayalog.Info(fmt.Sprintf("Mesage Haloe: %d", 1))
		kayalog.Info(fmt.Sprintf("Mesage Hi: %d", 2))
		kayalog.Info(fmt.Sprintf("Mesage Request: %d", 3))
	}
	fmt.Println("test")

	fmt.Scanln()
}
