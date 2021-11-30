package fields

import (
	"bytes"
	"database/sql/driver"
	"strconv"
	"strings"
)

type Uint64s []uint64

func (u *Uint64s) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	v, ok := src.([]byte)
	if !ok {
		return ErrInvalidDBValueForUint64ToString
	}

	strS := strings.Split(string(v), ",")

	*u = make([]uint64, 0, len(strS))
	for _, v := range strS {
		num, _ := strconv.ParseUint(v, 10, 64)
		*u = append(*u, num)
	}
	return nil
}

func (u Uint64s) Value() (driver.Value, error) {
	if len(u) == 0 {
		return "", nil
	}

	var buffer bytes.Buffer
	for _, v := range u {
		buffer.WriteString(strconv.FormatUint(v, 10))
		buffer.WriteString(",")
	}
	return strings.Trim(buffer.String(), ","), nil
}
