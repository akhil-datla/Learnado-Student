/*
 * File: courses.go
 * File Created: Thursday, 24th November 2022 9:04:52 pm
 * Last Modified: Friday, 23rd June 2023 1:00:50 am
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package courses

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"main/components/security"
	"net/http"
	"os"

	"github.com/denisbrodbeck/machineid"
	"github.com/spf13/afero"
)

var AppFs = afero.NewMemMapFs()

// RegisterLicense sends a license key and hardware ID to the server for registration.
func RegisterLicense(licenseKey string) (string, error) {
	hardwareID := getHardwareID()

	// Prepare the request body
	postBody, _ := json.Marshal(map[string]string{
		"licenseKey": licenseKey,
		"hardwareID": hardwareID,
	})
	responseBody := bytes.NewBuffer(postBody)

	// Send the POST request to register the license
	resp, err := http.Post(fmt.Sprintf("%s/licenses/register", os.Getenv("URL")), "application/json", responseBody)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	sb := string(body)
	return sb, nil
}

// DownloadCourses retrieves course data from the server and saves it to a file.
func DownloadCourses() error {
	hardwareID := getHardwareID()

	// Prepare the request body
	postBody, _ := json.Marshal(map[string]string{
		"hardwareID": hardwareID,
	})
	responseBody := bytes.NewBuffer(postBody)

	// Send the POST request to download courses
	resp, err := http.Post(fmt.Sprintf("%s/download", os.Getenv("URL")), "application/json", responseBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for an error response
	if string(body) == "Error downloading courses" {
		return fmt.Errorf("error downloading courses")
	}

	// Save the downloaded courses to a file
	err = ioutil.WriteFile("fs.gob", body, 0644)
	if err != nil {
		return err
	}

	return nil
}

// LoadCourses loads the previously downloaded courses from the file system.
func LoadCourses() error {
	encryptedData, err := os.ReadFile("fs.gob")
	if err != nil {
		return err
	}

	// Decrypt and decompress the map
	decryptedDecompressedMap, err := DecryptAndDecompressMap(encryptedData, getHardwareID())
	if err != nil {
		return err
	}

	// Convert the map to Afero file system
	MapToAferoFS(decryptedDecompressedMap)

	return nil
}

// GobDecodeMapFromBytes decodes a gob-encoded byte slice into a map[string][]byte.
func GobDecodeMapFromBytes(data []byte) (map[string][]byte, error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	var m map[string][]byte
	err := dec.Decode(&m)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}

	return m, nil
}

// MapToAferoFS writes the content of a map to the Afero file system.
func MapToAferoFS(m map[string][]byte) {
	for path, content := range m {
		fullPath := "/" + path // add the leading slash

		// Write the file or create the directory
		if len(content) > 0 {
			afero.WriteFile(AppFs, fullPath, content, 0644)
		} else {
			AppFs.Mkdir(fullPath, 0755)
		}
	}
}

// DecryptAndDecompressMap decrypts and decompresses a map using AES decryption with a given key.
func DecryptAndDecompressMap(data []byte, key string) (map[string][]byte, error) {
	// Decrypt the encrypted map
	decryptedMap, err := security.Decrypt(data, security.DeriveKey(key))
	if err != nil {
		return nil, err
	}

	// Decompress the decrypted map
	decompressedMap, err := decompress(decryptedMap)
	if err != nil {
		return nil, err
	}

	// Decode the decompressed map
	decodedMap, err := GobDecodeMapFromBytes(decompressedMap)
	if err != nil {
		return nil, err
	}

	return decodedMap, nil
}

// getHardwareID retrieves the unique hardware ID of the machine.
func getHardwareID() string {
	hardwareID, err := machineid.ID()
	if err != nil {
		panic(err)
	}
	return hardwareID
}

// decompress performs gzip decompression on the input data.
func decompress(data []byte) ([]byte, error) {
	buf := bytes.NewReader(data)
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	var result bytes.Buffer
	_, err = io.Copy(&result, r)
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}
