package pkg

import (
	"os"
	"strconv"
)

func GetEnvInt(txt string, def string) int {
	if len(txt) == 0 {
		res, _ := strconv.Atoi(def)
		return res
	}
	res, _ := strconv.Atoi(os.Getenv(txt))
	return res
}
