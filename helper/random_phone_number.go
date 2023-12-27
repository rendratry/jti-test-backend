package helper

import (
	"crypto/rand"
	"io"
	"math/big"
)

func RandomPhoneNumber() string {
	maxNumber := 9
	b := make([]byte, maxNumber)
	n, err := io.ReadAtLeast(rand.Reader, b, maxNumber)
	if n != maxNumber {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	resultNumber := "08" + string(b)
	return resultNumber
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func RandomProvider() string {
	stringProvider := []string{"TELKOMSEL", "INDOSAT", "XL", "By.U", "SMARTFREN", "MENTARI"}
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(stringProvider))))
	if err != nil {
		panic(err)
	}
	index := randomIndex.Int64()

	return stringProvider[index]
}
