package sense

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := int64(t.Unix()) * 1000
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}

func (t Timestamp) EncodeValues(key string, v *url.Values) error {
	ts := t.Unix() * 1000
	v.Set(key, fmt.Sprint(ts))
	return nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	sec := ts / 1000
	nsec := (ts % 1000) * 1000
	*t = Timestamp{time.Unix(int64(sec), int64(nsec))}
	return nil
}
