package factory

import (
	"crypto/rand"
	"fmt"
	"os"
)

func NewKey() (key []byte) {
	key, err := readKey()
	if err != nil {
		key = make([]byte, 32)
		if _, err = rand.Read(key); err != nil {
			panic(fmt.Sprintf("ошибка генерации ключа: %s", err.Error()))
		}

		if err = writeKey(key); err != nil {
			panic(fmt.Sprintf("ошибка записи ключа: %s", err.Error()))
		}
	}

	return key
}

// readKey attemps to read key from secret.key.
// If len(secret.key) > 32, readKey truncates it to 32 bytes.
// If len(secret.key) < 32, readKey returns error
func readKey() (key []byte, err error) {
	file, err := os.Open("./secret.key")
	if err != nil {
		return
	}

	key = make([]byte, 32)
	_, err = file.Read(key)

	return
}

// writeKey writes provided key to secret.key
func writeKey(key []byte) (err error) {
	file, err := os.Create("./secret.key")
	if err != nil {
		return
	}

	_, err = file.Write(key)

	return
}
