package bins

import (
	"encoding/json"
	"errors"
	"time"
)

type Bin struct {
	Id        int32     `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	binList []Bin
}

func NewBin(id int32, name string, private bool) (*Bin, error) {
	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}

	newBin := &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}

	return newBin, nil
}

func (bin *Bin) BinToBytes() ([]byte, error) {
	file, err := json.Marshal(bin)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewBinList() BinList {
	return BinList{
		binList: []Bin{},
	}
}

func (bl *BinList) AddBin(bin Bin) {
	bl.binList = append(bl.binList, bin)
}
