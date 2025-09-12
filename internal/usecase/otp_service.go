package usecase

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"
)

type OtpStore interface {
	Create(userID, otp string, ttl time.Duration) error
	Get(userID string) (string, error)
	Delete(userID string) error
}

type OtpSender interface {
	Send(message string, receptors ...string) error
}
type OtpService struct {
	len     int
	ttl     time.Duration
	storage OtpStore
	sender  OtpSender
}

func NewOtpService(len int, ttl time.Duration, storage OtpStore, sender  OtpSender) *OtpService {

	return &OtpService{
		len: len,
		ttl: ttl,
		storage: storage,
		sender: sender,
	}
}

func (s *OtpService) RequestOTP(phone string) error {
	log.Println("sending for ", phone)
	code, err := generateCode(s.len)
	if err != nil {
		return err
	}
	log.Println("code is ", code)
	if err := s.storage.Create(phone, code, s.ttl); err != nil {
		return err
	}

	log.Println("lest's go to send otp")

	if err := s.sender.Send(code, phone); err != nil {
		return err
	}

	log.Println("otp", code)

	return nil
}

func (s *OtpService) VerifyOTP(userID, otp string) error {
	storedOTP, err := s.storage.Get(userID)
	if err != nil {
		return err
	}

	if storedOTP != otp {
		return fmt.Errorf("invalid otp")
	}

	if err := s.storage.Delete(userID); err != nil {
		return errors.New("failed to delete OTP: " + err.Error())
	}
	return nil
}

// --------------- helper ------------------------
func generateCode(length int) (string, error) {
	const digits = "0123456789"
	otp := make([]byte, length)

	for i := range length {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[n.Int64()]
	}

	return string(otp), nil
}
