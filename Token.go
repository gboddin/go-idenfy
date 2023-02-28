package idenfy

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type TokenRequest struct {
	ClientId            string     `json:"clientId"`
	Firstname           string     `json:"firstName,omitempty"`
	Lastname            string     `json:"lastName,omitempty"`
	SuccessUrl          string     `json:"successUrl,omitempty"`
	ErrorUrl            string     `json:"errorUrl,omitempty"`
	UnverifiedUrl       string     `json:"unverifiedUrl,omitempty"`
	Locale              string     `json:"locale,omitempty"`
	ShowInstructions    *bool      `json:"showInstructions,omitempty"`
	ExpiryTime          *int       `json:"expiryTime,omitempty"`
	SessionLength       *int       `json:"sessionLength,omitempty"`
	Country             string     `json:"country,omitempty"`
	Documents           []Document `json:"documents,omitempty"`
	DateOfBirth         *Date      `json:"dateOfBirth,omitempty"`
	DateOfExpiry        *Date      `json:"dateOfExpiry,omitempty"`
	DateOfIssue         *Date      `json:"dateOfIssue,omitempty"`
	Nationality         string     `json:"nationality,omitempty"`
	PersonalNumber      string     `json:"personalNumber,omitempty"`
	DocumentNumber      string     `json:"documentNumber,omitempty"`
	Sex                 string     `json:"sex,omitempty"`
	GenerateDigitString *bool      `json:"generateDigitString,omitempty"`
	Address             string     `json:"address,omitempty"`
	TokenType           Type       `json:"tokenType,omitempty"`
	VideoCallQuestions  []string   `json:"videoCallQuestions,omitempty"`
	ExternalRef         string     `json:"externalRef,omitempty"`
	UtilityBill         *bool      `json:"utilityBill,omitempty"`
	CallbackUrl         string     `json:"callbackUrl,omitempty"`
	Questionnaire       string     `json:"questionnaire,omitempty"`
}

type TokenResponse struct {
	ClientId            string     `json:"clientId"`
	Firstname           string     `json:"firstName,omitempty"`
	Lastname            string     `json:"lastName,omitempty"`
	SuccessUrl          string     `json:"successUrl,omitempty"`
	ErrorUrl            string     `json:"errorUrl,omitempty"`
	UnverifiedUrl       string     `json:"unverifiedUrl,omitempty"`
	Locale              string     `json:"locale,omitempty"`
	ShowInstructions    *bool      `json:"showInstructions,omitempty"`
	ExpiryTime          *int       `json:"expiryTime,omitempty"`
	SessionLength       *int       `json:"sessionLength,omitempty"`
	Country             string     `json:"country,omitempty"`
	Documents           []Document `json:"documents,omitempty"`
	DateOfBirth         *Date      `json:"dateOfBirth,omitempty"`
	DateOfExpiry        *Date      `json:"dateOfExpiry,omitempty"`
	DateOfIssue         *Date      `json:"dateOfIssue,omitempty"`
	Nationality         string     `json:"nationality,omitempty"`
	PersonalNumber      string     `json:"personalNumber,omitempty"`
	DocumentNumber      string     `json:"documentNumber,omitempty"`
	Sex                 string     `json:"sex,omitempty"`
	GenerateDigitString *bool      `json:"generateDigitString,omitempty"`
	Address             string     `json:"address,omitempty"`
	TokenType           Type       `json:"tokenType,omitempty"`
	VideoCallQuestions  []string   `json:"videoCallQuestions,omitempty"`
	ExternalRef         string     `json:"externalRef,omitempty"`
	UtilityBill         *bool      `json:"utilityBill,omitempty"`
	CallbackUrl         string     `json:"callbackUrl,omitempty"`
	Questionnaire       string     `json:"questionnaire,omitempty"`
	Message             string     `json:"message,omitempty"`
	AuthToken           string     `json:"authToken,omitempty"`
	ScanRef             string     `json:"scanRef,omitempty"`
	DigitString         string     `json:"digitString,omitempty"`
}

// GetRedirectUrl Returns the URL to redirect the client to complete the verification
func (t TokenResponse) GetRedirectUrl() string {
	return "https://ivs.idenfy.com/api/v2/redirect?authToken=" + url.QueryEscape(t.AuthToken)
}

// Validate Validates the token request against Idenfy specifications
func (t TokenRequest) Validate() error {
	if err := t.validateLen("clientId", t.ClientId, 1, 100); err != nil {
		return err
	}
	if len(t.ClientId) < 1 {
		return errors.New("clientId is empty")
	}
	if len(t.ClientId) > 100 {
		return errors.New("clientId must be max 100 chars")
	}
	if err := t.validateName("firstName", t.Firstname); err != nil {
		return err
	}
	if err := t.validateName("lastName", t.Lastname); err != nil {
		return err
	}
	if err := t.validateUrl("successUrl", t.SuccessUrl); err != nil {
		return err
	}
	if err := t.validateUrl("errorUrl", t.ErrorUrl); err != nil {
		return err
	}
	if err := t.validateUrl("unverifiedUrl", t.UnverifiedUrl); err != nil {
		return err
	}
	if len(t.Locale) != 0 && len(t.Locale) != 2 {
		return errors.New("locale must be 2 char")
	}
	if t.ExpiryTime != nil && (*t.ExpiryTime < 60 || *t.ExpiryTime > 2592000) {
		return errors.New("expiryTime must be between 60 and 2592000")
	}
	if t.SessionLength != nil && (*t.SessionLength < 60 || *t.SessionLength > 3600) {
		return errors.New("expiryTime must be between 60 and 3600")
	}
	if len(t.Country) != 0 && len(t.Country) != 2 {
		return errors.New("country must be 2 char")
	}
	for _, document := range t.Documents {
		if !t.isValidDocument(document) {
			return errors.New("invalid document type")
		}
	}
	// Date is handled by MarshalJSON
	if len(t.Nationality) != 0 && len(t.Nationality) != 2 {
		return errors.New("nationality must be 2 char")
	}
	if len(t.Sex) > 0 {
		if len(t.Sex) != 1 {
			return errors.New("sex must be 1 char")
		}
		if t.Sex != "M" && t.Sex != "F" {
			return errors.New("sex must be M or F")
		}
	}
	if err := t.validateLen("address", t.Address, 0, 255); err != nil {
		return err
	}
	if len(t.TokenType) > 1 && !t.isValidSessionType(t.TokenType) {
		return errors.New("not a valid session type : " + string(t.TokenType))
	}
	if err := t.validateLen("externalRef", t.ExternalRef, 0, 40); err != nil {
		return err
	}
	if err := t.validateUrl("callbackUrl", t.CallbackUrl); err != nil {
		return err
	}
	return nil
}

func (t TokenRequest) validateUrl(fieldName, value string) error {
	if len(value) == 0 {
		return nil
	}
	parsedUrl, err := url.Parse(value)
	if err != nil {
		return errors.New(fmt.Sprintf("%s : %s", fieldName, err))
	}
	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return errors.New(fmt.Sprintf("%s : %s %s", fieldName, "unsupported scheme", parsedUrl.Scheme))
	}
	return nil
}

func (t TokenRequest) validateLen(fieldName, value string, min, max int) error {
	if len(value) < min {
		return errors.New(fmt.Sprintf("%s : must be at least %d char", fieldName, min))
	}
	if max > 0 && len(value) > max {
		return errors.New(fmt.Sprintf("%s : must be max %d char", fieldName, max))
	}
	return nil
}

func (t TokenRequest) validateName(fieldName, value string) error {
	if len(value) == 0 {
		return nil
	}
	if err := t.validateLen(fieldName, value, 1, 100); err != nil {
		return err
	}
	if strings.ContainsAny(value, "01234567890~!@#$%^*()_+={}[]\\|:;\",<>/?") {
		return errors.New(fieldName + " contains illegal chars")
	}
	return nil
}

func (t TokenRequest) isValidDocument(doc Document) bool {
	for _, docTest := range validDocuments {
		if docTest == doc {
			return true
		}
	}
	return false
}

func (t TokenRequest) isValidSessionType(sessType Type) bool {
	for _, sessTypeTest := range validSessionTypes {
		if sessTypeTest == sessType {
			return true
		}
	}
	return false
}

type Type string

const (
	DOCUMENT                  Type = "DOCUMENT"
	IDENTIFICATION            Type = "IDENTIFICATION"
	VIDEO_CALL                Type = "VIDEO_CALL"
	VIDEO_CALL_PHOTOS         Type = "VIDEO_CALL_PHOTOS"
	VIDEO_CALL_IDENTIFICATION Type = "VIDEO_CALL_IDENTIFICATION"
)

var validSessionTypes = []Type{DOCUMENT, IDENTIFICATION, VIDEO_CALL, VIDEO_CALL_PHOTOS, VIDEO_CALL_IDENTIFICATION}
