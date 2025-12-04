package randomnumber

type FakeRandomNumberService struct {
}

func (r *FakeRandomNumberService) NewRandomNumber(maxNumber int) int {
	return 2
}
