package common

import "log"

func PrintErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
