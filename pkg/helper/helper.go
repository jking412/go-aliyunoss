package helper

import (
	"fmt"
	"os"
	"time"
)

func FileIsExist(filePath string) error {
	if _, err := os.Stat(filePath); err != nil {
		return err
	}
	return nil
}

func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}
