package validator

import (
	"errors"
	"regexp"
	"strings"
)

var (
	rxPhone        = regexp.MustCompile(`^(13|14|15|16|17|18|19)\d{9}$`)
	ErrPhoneFormat = errors.New("phone format error")
)

func ValidateRxPhone(phone string) error {
	phone = strings.TrimSpace(phone)
	if !rxPhone.MatchString(phone) {
		return ErrPhoneFormat
	}

	return nil
}
