package text

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{}<>?"
)

// GenerateSecureRandomText generates a secure random text of given length.
// By default, it includes lowercase + uppercase + digits.
// If symbol == true, symbols are included as well.
func GenerateSecureRandomText(length int, symbol bool) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	charSets := []string{lowercase, uppercase, digits}
	if symbol {
		charSets = append(charSets, symbols)
	}

	text := make([]byte, length)

	// Assign one mandatory character from each set
	for i, set := range charSets {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(set))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %w", err)
		}
		text[i] = set[num.Int64()]
	}

	// Build allChars
	all := ""
	for _, set := range charSets {
		all += set
	}

	// Fill remaining positions
	for i := len(charSets); i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(all))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %w", err)
		}
		text[i] = all[num.Int64()]
	}

	// Shuffle
	for i := len(text) - 1; i > 0; i-- {
		jNum, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", fmt.Errorf("failed to shuffle: %w", err)
		}
		j := int(jNum.Int64())
		text[i], text[j] = text[j], text[i]
	}

	return string(text), nil
}
