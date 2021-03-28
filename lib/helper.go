package lib

import "log"

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
