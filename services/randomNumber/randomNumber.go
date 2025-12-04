package randomnumber

import (
	"math/rand/v2"
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
	return rand.IntN(maxNumber)
}
