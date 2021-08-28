package handlers

import "log"

func PanicOnError(e error) {
	if e == nil {
		return
	}
	log.Panic(e)
}