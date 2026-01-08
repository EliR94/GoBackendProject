package uuid

type FakeUUIDService struct {
	fakeUUID string
}

func (r *FakeUUIDService) NewUUID() string {
	return r.fakeUUID
}

func (r *FakeUUIDService) StoreFakeUUID(newFakeUUID string) {
	r.fakeUUID = newFakeUUID
}
