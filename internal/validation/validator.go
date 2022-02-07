// Package validation contain CustomValidator for echo server
package validation

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	"net/http"
)

// CustomValidator instance is used as echo server validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate checks structure's fields for correspondence to its tags
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// NewValidator returns new instance of CustomValidator
func NewValidator() (*CustomValidator, error) {
	translator := en.New()
	uni := ut.New(translator, translator)

	// this is usually known or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, found := uni.GetTranslator("en")
	if !found {
		return nil, fmt.Errorf("translator not found")
	}
	v := validator.New()
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		return nil, fmt.Errorf("en_translation error %w", err)
	}
	err := v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	if err != nil {
		return nil, fmt.Errorf("register translation error %w", err)
	}
	return &CustomValidator{validator: v}, nil
}
