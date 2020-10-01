package main

import (
	"github.com/tzzh/pod-tzzh-mail/babashka"
	"github.com/tzzh/pod-tzzh-mail/mail"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	debug := os.Getenv("POD_TZZH_MAIL_DEBUG")
	if debug != "true" {
		log.SetOutput(ioutil.Discard)
	}

	for {
		message := babashka.ReadMessage()
		mail.ProcessMessage(message)
	}
}
