package tools

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

func CreateLogFile() *os.File {
	file, err := os.OpenFile("LOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func GenerateID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func GenerateDoubleID() string {
	return strings.ReplaceAll(GenerateID()+GenerateID(), "-", "")
}

func ValidateRegistration(email, password, username *string) error {
	if strings.TrimSpace(*email) == "" {
		return errors.New("email cannot be empty")
	}
	if len(*password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if strings.TrimSpace(*username) == "" {
		return errors.New("username cannot be empty")
	}
	if !govalidator.IsEmail(*email) {
		return errors.New("invalid email format")
	}
	if !isValidUsername(*username) {
		return errors.New("username can only contain letters, numbers, and the symbols _ -")
	}
	if err := validatePasswordComplexity(*password); err != nil {
		return err
	}
	return nil
}

func isValidUsername(username string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9_\-.]{3,30}$`).MatchString(username)
}

func validatePasswordComplexity(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func ValidateEmail(email string) bool {
	return govalidator.IsEmail(email)
}
