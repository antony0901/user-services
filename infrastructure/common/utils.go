package common

import "fmt"

func Check(err error) {
	if err != nil {
		fmt.Println(err)
		// log.Printf("%s", err)
		panic(err)
	}
}
