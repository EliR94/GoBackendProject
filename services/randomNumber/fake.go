package randomnumber

type FakeRandomNumberService struct {
	FakeRandomNumber int
}

func (r *FakeRandomNumberService) NewRandomNumber(maxNumber int) int {
	return r.FakeRandomNumber
}

func (r *FakeRandomNumberService) StoreFakeRandomNumber(newFakeRandomNumber int) {
	r.FakeRandomNumber = newFakeRandomNumber
}
