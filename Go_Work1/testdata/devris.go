package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func main78() {
	// Generate a random device ID
	deviceID, err := generateDeviceID()
	if err != nil {
		log.Fatal("Error generating device ID:", err)
	}

	fmt.Println("Device ID:", deviceID)
}

func generateDeviceID() (string, error) {
	// Generate 16 bytes of random data
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Convert the random bytes to a hexadecimal string
	deviceID := hex.EncodeToString(randomBytes)
	fmt.Println(randomBytes)

	return deviceID, nil
}
