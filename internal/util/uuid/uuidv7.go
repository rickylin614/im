package uuid

import (
	"sync"

	"github.com/gofrs/uuid"
)

var mu = &sync.Mutex{}

func New() string {
	mu.Lock()
	defer mu.Unlock()
	id, _ := uuid.NewV7()
	return id.String()
}
