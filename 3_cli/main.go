package main

import (
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	binList []Bin
}

func newBin(id, name string, private bool) Bin {
	return Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func newBinList() BinList {
	return BinList{
		binList: []Bin{},
	}
}

func (bl *BinList) addBin(bin Bin) {
	bl.binList = append(bl.binList, bin)
}

func main() {

}
