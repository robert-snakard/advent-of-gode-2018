package check

import (
	"fmt"
)

func Check(e error) {
	if e != nil {
		fmt.Println("Go has encountered an error:", e)
		panic(e)
	}
}
