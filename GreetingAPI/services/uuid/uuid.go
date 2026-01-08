package uuid

import (
	"github.com/google/uuid"
)

// the interface which oulines anything that has a NewUUID func is a UUIDService
type UUIDService interface {
	NewUUID() string
}

// this struct respresents the real uuid service
type RealUUIDService struct {
}

// this function provides the uuid service with a NewUUID method and enables the stucture on line 13 to become a UUIDService
func (r *RealUUIDService) NewUUID() string {
	return uuid.NewString()
}
