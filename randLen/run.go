package randLen

import (
	"fmt"
	"os"
	"strconv"
)

func Run() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(RandLenStringGenerator(8))
	} else if n, err := strconv.Atoi(args[0]); err == nil {
		fmt.Println(RandLenStringGenerator(n))
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
