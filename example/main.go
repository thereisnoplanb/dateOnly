package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/thereisnoplanb/dateOnly"
)

func main() {
	a := dateOnly.Date{}
	_ = a

	b := time.Time{}
	_ = b

	c := dateOnly.Date{}
	_ = c

	jsond := `{
		"date1" : null,
		"date2" : "2000-01-01",
		"time1" : null,
		"time2" : "2000-01-01T15:00:00Z"
	}`

	type testowe struct {
		Date1 *dateOnly.Date `json:"date1"`
		Date2 dateOnly.Date  `json:"date2"`
		Time1 *time.Time     `json:"time1"`
		Time2 time.Time      `json:"time2"`
	}

	t := new(testowe)

	err := json.Unmarshal([]byte(jsond), t)
	_ = err

	bytes, _ := json.Marshal(t)

	aa := string(bytes)
	_ = aa

	fmt.Printf("%s", a)
	fmt.Printf("%#v", a)
	fmt.Printf("%#v", b)

	a.After(c)
}
