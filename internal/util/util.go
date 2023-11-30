package util

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"strings"
	"time"
)

// CalculateAge calculate age from birthdate
func CalculateAge(birthDate time.Time) int {
	currentTime := time.Now()
	years := currentTime.Year() - birthDate.Year()

	if currentTime.Month() < birthDate.Month() ||
		(currentTime.Month() == birthDate.Month() && currentTime.Day() < birthDate.Day()) {
		years--
	}

	return years
}

// GenerateHashPassword generate bcrypt hash password
func GenerateHashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
}

// StringIsEmpty checks if a string is empty ” or if it only has empty spaces
func StringIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsEmailValid checks if email is valid
func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// ParseDate parse string date with format 2006-01-02 to time.Time
func ParseDate(dateString string) (time.Time, error) {
	return time.Parse("2006-01-02", dateString)
}

// IsUUID checks if uuid is valid
func IsUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
