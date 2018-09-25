package lrmp

import (
	"fmt"
	"log"
)

func logDebug(args ...interface{}) {
	if isDebug() {
		fmt.Println(args)
	}
}
func isDebug() bool { return true }

func logError(args ...interface{}) {
	log.Println(args)
}
