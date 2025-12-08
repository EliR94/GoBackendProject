package randomnumber

type FakeRandomNumberService struct {
	fakeRandomNumber int
}

func (r *FakeRandomNumberService) NewRandomNumber(maxNumber int) int {
	return r.fakeRandomNumber
}

func (r *FakeRandomNumberService) StoreFakeRandomNumber(newFakeRandomNumber int) {
	r.fakeRandomNumber = newFakeRandomNumber
}
