package helper

import "os"

func FileIsExist(filePath string) error {
	if _, err := os.Stat(filePath); err != nil {
		return err
	}
	return nil
}
