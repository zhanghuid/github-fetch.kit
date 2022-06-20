package util

import (
	"fmt"
	"os"

	"github.com/golang-module/carbon"
	"github.com/gookit/color"
)

const (
	LOG_TYPE_DANGER = "danger"
	LOG_TYPE_INFO   = "info"
)

func Log(level string, message string, isExit bool) {
	str := fmt.Sprintf("[%s] %s", carbon.Now().ToDateTimeString(), message)
	color.GetTheme(level).Println(str)

	if isExit {
		os.Exit(9)
	}
}
