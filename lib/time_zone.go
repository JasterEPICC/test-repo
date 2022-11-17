package lib

import (
	"fmt"
	"time"
)

func InitTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		LogError(fmt.Sprintf("Cannot start server: %v", err), true)
	}

	time.Local = ict
}
