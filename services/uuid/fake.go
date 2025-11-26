package uuid

type FakeUUIDService struct {
	FakeUUID string
}

func (r *FakeUUIDService) NewUUID() string {
	return r.FakeUUID
}

func (r *FakeUUIDService) StoreFakeUUID(newFakeUUID string) {
	r.FakeUUID = newFakeUUID
}
