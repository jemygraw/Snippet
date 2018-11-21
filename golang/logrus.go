package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal":  "walrus",
		"species": "lions",
		"age":     22,
	}).Info("A walrus appears")
}
