package main

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithError(errors.New("test")).Fatal("test with error")
}
