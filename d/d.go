package d

import (
	"errors"
)

//DerivedApplicationSecretKey
type D interface {
	// Compute computes derived application secret key(dask)
	Compute(R2 []byte, ask []byte) ([]byte, error)
}

var (
	_ D = AppleD{}
)

type AppleD struct {
}

func (d AppleD) Compute(R2 []byte, ask []byte) ([]byte, error) {
	if len(R2) == 0 {
		return nil, errors.New("R2 block doesn't exist.")
	}

	// b, err := hex.DecodeString("d87ce7a26081de2e8eb8acef3a6dc179") //Apple provided

	if len(ask) != 16 {
		return nil, errors.New("ask key length doesn't equal 16")
	}

	return ask, nil
}
