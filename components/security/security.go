/*
 * File: security.go
 * File Created: Thursday, 22nd June 2023 3:45:02 pm
 * Last Modified: Friday, 23rd June 2023 1:01:40 am
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package security

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"golang.org/x/crypto/sha3"
)

// Decrypt decrypts the ciphertext using AES decryption with the provided key.
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Separate the nonce and ciphertext
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the ciphertext
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// DeriveKey derives a valid AES encryption key from the hardware ID using SHA-256.
func DeriveKey(hardwareID string) []byte {
	hasher := sha3.New256()
	hasher.Write([]byte(hardwareID))
	return hasher.Sum(nil)
}
