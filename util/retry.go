package util

import (
	"log"
	"time"
)

func RetryUntilOk(fn func() error, errBaseMsg string) {

	var err error
	for {
		err = fn()
		if err == nil {
			return
		}
		log.Println(errBaseMsg)
		log.Println(err.Error())
		time.Sleep(10 * time.Second)
	}
}
