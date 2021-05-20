package log

import (
	"fmt"
	"os"
)

func Fatal(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(-1)
}
