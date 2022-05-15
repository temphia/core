package lib

import (
	"crypto/rand"
	"math/big"
	"os"

	"github.com/k0kubun/pp"
)

func Die(args ...interface{}) {

	pp.Println(args...)
	os.Exit(1)

}

func ReadFile(file string, paths ...string) ([]byte, error) {
	return nil, nil
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
