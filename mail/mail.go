package mail

import (
	"encoding/json"
	"errors"
	"github.com/jordan-wright/email"
	"github.com/tzzh/pod-tzzh-mail/babashka"
	"net/smtp"
	"net/textproto"
	"strconv"
)

type RawInput struct {
	Host        string
	Port        int
	Username    string
	Password    string
	ReplyTo     []string
	From        string
	To          []string
	Bcc         []string
	Cc          []string
	Subject     string
	Text        string // Plaintext message (optional)
	HTML        string // Html message (optional)
	Sender      string // override From as SMTP envelope sender (optional)
	Headers     textproto.MIMEHeader
	Attachments []string // paths to attachement files
}

func ProcessMessage(message *babashka.Message) {

	if message.Op == "describe" {
		response := &babashka.DescribeResponse{
			Format: "json",
			Namespaces: []babashka.Namespace{
				{Name: "pod.tzzh.mail",
					Vars: []babashka.Var{
						{Name: "send-mail"},
					},
				},
			},
		}
		babashka.WriteDescribeResponse(response)

	} else if message.Op == "invoke" {

		switch message.Var {
		case "pod.tzzh.mail/send-mail":
			inputList := []RawInput{}
			err := json.Unmarshal([]byte(message.Args), &inputList)
			if err != nil {
				babashka.WriteErrorResponse(message, err)
				return
			}
			if len(inputList) != 1 {
				e := errors.New("Wrong number of argument, send-mail expects 1 argument")
				babashka.WriteErrorResponse(message, e)
				return
			}
			input := &inputList[0]
			auth := smtp.PlainAuth("", input.Username, input.Password, input.Host)
			addr := input.Host + ":" + strconv.Itoa(input.Port)
			e := &email.Email{
				ReplyTo: input.ReplyTo,
				From:    input.From,
				To:      input.To,
				Bcc:     input.Bcc,
				Cc:      input.Cc,
				Subject: input.Subject,
				Text:    []byte(input.Text),
				HTML:    []byte(input.HTML),
				Sender:  input.Sender,
				Headers: input.Headers,
			}
			for _, a := range input.Attachments {
				_, err := e.AttachFile(a)
				if err != nil {
					babashka.WriteErrorResponse(message, err)
					return
				}
			}
			err = e.Send(addr, auth)
			if err != nil {
				babashka.WriteErrorResponse(message, err)
				return
			}
			babashka.WriteInvokeResponse(message, nil)
		}
	}
}
