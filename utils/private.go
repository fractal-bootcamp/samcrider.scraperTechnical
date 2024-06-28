package utils

import (
	"fmt"
	"os"
)

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
