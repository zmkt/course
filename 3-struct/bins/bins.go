package bins

import (
	"3-struct/bins"
	"time"
)

type Bin struct {
	ID       string
	Private  bool
	CreateAt time.Time
	Name     string
}

func NewBin(id string, private bool, createAt time.Time, name string) Bin {
	return Bin{
		ID:       id,
		Private:  private,
		CreateAt: createAt,
		Name:     name,
	}
}

type BinList struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBinList(id string, private bool, createdAt time.Time, name string) BinList {
	return BinList{
		ID:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}
