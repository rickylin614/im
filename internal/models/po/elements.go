package po

import (
	"strconv"
	"time"

	"github.com/goccy/go-json"
)

type DateTime time.Time

func NewDateTime(t time.Time) *DateTime {
	dt := DateTime(t)
	return &dt
}

// MarshalJSON 實現 json.Marshaler 接口
func (mt DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(mt)
	formattedTime := int(t.Unix())
	return []byte(strconv.Itoa(formattedTime)), nil
}

// UnmarshalJSON 實現 json.Unmarshaler 接口
func (mt *DateTime) UnmarshalJSON(data []byte) error {
	var unixSeconds int64
	if err := json.Unmarshal(data, &unixSeconds); err != nil {
		return err
	}

	t := time.Unix(unixSeconds, 0)
	*mt = DateTime(t)
	return nil
}
