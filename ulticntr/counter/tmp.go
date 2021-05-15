package counter

import (
	"encoding/binary"
	"io"
	"os"
)

type tmp struct{}

func newTmp() (*tmp, error) {
	return &tmp{}, nil
}

func (t *tmp) Count() (c uint64, err error) {
	f, err := os.OpenFile(".tmpcntr", os.O_RDWR|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		return c, err
	}
	err = binary.Read(f, binary.LittleEndian, &c)
	if err != nil && err != io.EOF {
		return c, err
	}
	c++
	f.Truncate(0)
	f.Seek(0, 0)
	err = binary.Write(f, binary.LittleEndian, c)
	return c, err
}
