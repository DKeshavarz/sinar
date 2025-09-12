package sms

import (
	"fmt"

	"github.com/DKeshavarz/sinar/internal/usecase"
	"github.com/kavenegar/kavenegar-go"
)

type smsSender struct {
	api    *kavenegar.Kavenegar
	sender string
}

func New(apiKey, sender string) usecase.OtpSender {
	return &smsSender{
		api: kavenegar.New(apiKey),
		sender: sender,
	}
}
func (s *smsSender) Send(message string, receptor... string) error {
	if res, err := s.api.Message.Send(s.sender, receptor, message, nil); err != nil {
		// switch err := err.(type) {
		// case *kavenegar.APIError:
		//   fmt.Println(err.Error())
		// case *kavenegar.HTTPError:
		//   fmt.Println(err.Error())
		// default:
		//   fmt.Println(err.Error())
		// }
		return err
	}else {
		for _, r := range res {
		  fmt.Println("MessageID   = ", r.MessageID)
		  fmt.Println("Status      = ", r.Status)
		}
	}
	
	return nil
}
