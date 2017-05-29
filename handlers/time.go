package handlers

import (
	"fmt"
	"time"
)

func getTime(args []string) string {
	return fmt.Sprintf("unix timestamp: %d", int32(time.Now().Unix()))
}
