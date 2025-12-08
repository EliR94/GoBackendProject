package randomnumber

import (
	"crypto/rand"
	"math/big"
)

// the interface which oulines anything that has a NewRandomNumber func is a RandomNumberService
type RandomNumberService interface {
	NewRandomNumber(maxNumber int) int
}

// this struct respresents the real randomNumber service
type RealRandomNumberService struct {
}

// this function provides the randomNumber service with a NewRandomNumber method and enables the stucture on line ? to become a RandomNumberService
func (r *RealRandomNumberService) NewRandomNumber(maxNumber int) int {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(int64(maxNumber)))
	return int(randomNumber.Int64())
}
