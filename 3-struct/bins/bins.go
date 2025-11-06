package bins

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	if id == "" {
		return nil, errors.New("invalid id")
	}
	if name == "" {
		return nil, errors.New("invalid name")
	}
	return &Bin{
		ID:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}, nil
}
func (b Bin) printBin() string {
	return fmt.Sprintf("Bin[ID=%s, Private=%t, CreatedAt=%s, Name=%s]",
		b.ID, b.Private, b.CreatedAt.Format(time.RFC3339), b.Name)
}

var BinList []Bin
