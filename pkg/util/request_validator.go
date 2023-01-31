package util

import (
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
	"regexp"
)

const (
	ethAddressRegexString = `^0x[0-9a-fA-F]{40}$`
)

var (
	ethAddressRegex = regexp.MustCompile(ethAddressRegexString)
)

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidation(val *validator.Validate) *RequestValidator {
	isEthAddress := func(fl validator.FieldLevel) bool {
		return ethAddressRegex.MatchString(fl.Field().String())
	}
	err := val.RegisterValidation("isEthAddress", isEthAddress)
	if err != nil {
		log.Fatalf("Error registring custom isEthAddress validator: %s", err)
	}
	return &RequestValidator{
		validator: val,
	}
}

func (r *RequestValidator) Validate(i interface{}) error {
	if err := r.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
