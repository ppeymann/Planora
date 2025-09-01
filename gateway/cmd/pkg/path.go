package pkg

import (
	"fmt"
	"runtime"
)

func getSchemaPath(mod string) string {
	if runtime.GOOS == "windows" {
		return fmt.Sprintf(".\\gateway\\schemas\\%s", mod)
	}

	return fmt.Sprintf("./gateway/schemas/%s", mod)
}
