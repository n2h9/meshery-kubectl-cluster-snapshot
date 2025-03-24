package main

import (
	"errors"
	"fmt"
	"os"
)

func fileExists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	if os.IsNotExist(err) {
		return errors.New(
			fmt.Sprintf("file not exists %s", path),
		)
	}
	return nil
}
