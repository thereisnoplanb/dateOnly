package dateOnly

import (
	"errors"
	"fmt"
	"time"
)

type Date time.Time

// Returns the time date + value.
//
// # Parameters
//
//	value time.Duration
//
// Value to add to date.
//
//	result time.Time
//
// # Returns
//
//	result time.Time
//
// The time date + value.
func (date Date) Add(value time.Duration) (resut time.Time) {
	return time.Time(date).Add(value)
}

func (date Date) AddDate(years int, months int, days int) Date {
	return Date(time.Time(date).AddDate(years, months, days))
}

func (date Date) AddDays(days int) Date {
	return Date(time.Time(date).AddDate(0, 0, days))
}

func (date Date) AddMonths(months int) Date {
	return Date(time.Time(date).AddDate(0, months, 0))
}

func (date Date) AddYears(years int) Date {
	return Date(time.Time(date).AddDate(years, 0, 0))
}

// Reports whether the date date is after value.
//
// # Parameters
//
//	value Date
//
// # Returns
//
//	result bool
//
// True if date is after value, false otherwise.
func (date Date) After(value Date) (result bool) {
	return time.Time(date).After(time.Time(value))
}

// Reports whether the date date is after value.
//
// # Parameters
//
//	value time.Time
//
// # Returns
//
//	result bool
//
// True if date is after value, false otherwise.
func (date Date) AfterTime(value time.Time) (result bool) {
	return time.Time(date).After(value)
}

func (date Date) AppendBinary(bytes []byte) ([]byte, error) {
	return time.Time(date).AppendBinary(bytes)
}

// AppendFormat is like [Time.Format] but appends the textual representation to bytes and returns the extended buffer.
func (date Date) AppendFormat(bytes []byte, layout string) []byte {
	return time.Time(date).AppendFormat(bytes, layout)
}

// AppendText implements the [encoding.TextAppender] interface.
//
// # Parameters
//
//	bytes []byte
//
// Array of bytes to add YYYY-MM-DD formatted date.
//
// # Returns
//
//	result []byte
//
// Array of bytes with added YYYY-MM-DD formatted date.
//
//	err error
//
// An error if the date cannot be represented as valid YYYY-MM-DD (e.g., the year is out of range).
func (date Date) AppendText(bytes []byte) (result []byte, err error) {
	return time.Time(date).AppendText(bytes)
}

// Reports whether the date date is before value.
//
// # Parameters
//
//	value Date
//
// # Returns
//
//	result bool
//
// True if date is before value, false otherwise.
func (date Date) Before(value Date) (result bool) {
	return time.Time(date).Before(time.Time(value))
}

// Reports whether the date date is before value.
//
// # Parameters
//
//	value time.Time
//
// # Returns
//
//	result bool
//
// True if date is before value, false otherwise.
func (date Date) BeforeTime(value time.Time) bool {
	return time.Time(date).Before(value)
}

// Reports whether date and other represent the same date.
//
// # Parameters
//
//	other Date
//
// Other date to check for equality.
//
// # Returns
//
//	result bool
//
// True if date is equal to other, false otherwise.
func (date Date) Equal(other Date) bool {
	return time.Time(date).Equal(time.Time(other))
}

// func (date Date) IsDST() bool {
// 	return time.Time(date).IsDST()
// }

// Reports whether date represents the zero date, January 1, year 1.
//
// # Returns
//
// True if date is January 1, year 1; false otherwise.
func (date Date) IsZero() bool {
	return time.Time(date).IsZero()
}

// Compares the date with value.
//
// # Parameters
//
//	value Date
//
// Value to compare with date.
//
// # Returns
//
//	value int
//
// If date is before value, it returns -1; if date is after value, it returns +1; if they're the same, it returns 0.
func (date Date) Compare(value Date) int {
	return time.Time(date).Compare(time.Time(value))
}

// Returns the year, month, and day in which date occurs.
//
// # Returns
//
//	year int
//
// The year in which date occurs.
//
//	month time.Month
//
// The month of year in which date occurs.
//
//	day int
//
// The day of month in which date occurs.
func (date Date) Deconstruct() (year int, month time.Month, day int) {
	return time.Time(date).Date()
}

// Returns the day of the month specified by date.
//
// # Returns
//
//	day int
//
// The day of the month specified by date.
func (date Date) Day() int {
	return time.Time(date).Day()
}

func (date Date) Format(layout string) string {
	return time.Time(date).Format(layout)
}

// Implements [fmt.GoStringer] and formats date to be printed in Go source code.
//
// # Returns
//
//	goString string
//
// Formatted date to be printed in Go source code.
func (date Date) GoString() (goString string) {
	y, _, d := date.Deconstruct()
	m := date.Format("January")
	return fmt.Sprintf("date.New(%d, time.%s, %d)", y, m, d)
}

func (date Date) GobEncode() ([]byte, error) {
	return time.Time(date).GobEncode()
}

// Returns the ISO 8601 year and week number in which date occurs.
//
// # Returns
//
//	year int
//
// The ISO 8601 year in which date occurs.
//
//	week int
//
// The ISO 8601 week number in which date occurs.
//
// # Remarks
//
// Week ranges from 1 to 53.
// Jan 01 to Jan 03 of year n might belong to week 52 or 53 of year n-1,
// and Dec 29 to Dec 31 might belong to week 1 of year n+1.
func (date Date) ISOWeek() (year int, week int) {
	return time.Time(date).ISOWeek()
}

// func (date Date) In(location *time.Location) Date {
// 	year, month, day := time.Time(date).In(location).Date()
// 	return New(year, month, day, location)
// }

// func (date Date) Local() Date {
// 	return Date(time.Time(date).Local().Truncate(24 * time.Hour))
// }

// func (date Date) Location() *time.Location {
// 	return time.Time(date).Location()
// }

// Implements the [encoding.BinaryMarshaler] interface.
//
// # Returns
//
//	data []byte
//
// The date as a slice of bytes.
//
//	err error
//
// If the date cannot be represented as a slices of bytes, then an error is reported.
func (date Date) MarshalBinary(t time.Time) (data []byte, err error) {
	return time.Time(date).MarshalBinary()
}

// Implements the [encoding/json.Marshaler] interface.
//
// # Returns
//
//	data []byte
//
// The date as a quoted string in the YYYY-MM-DD format.
//
//	err error
//
// If the date cannot be represented as YYYY-MM-DD (e.g., the year is out of range), then an error is reported.
func (date Date) MarshalJSON() (data []byte, err error) {
	return []byte(time.Time(date).Format("\"" + time.DateOnly + "\"")), nil
}

func (date Date) MarshalText() ([]byte, error) {
	return []byte(time.Time(date).Format(time.DateOnly)), nil
}

// Returns the month of the year specified by date.
//
// # Returns
//
//	month time.Month
//
// The month of the year specified by date.
func (date Date) Month() time.Month {
	return time.Time(date).Month()
}

// func (date Date) Round(value time.Duration) Date {
// 	return Date(time.Time(date).Round(value))
// }

// Returns the date formatted using the format string "2006-01-02".
// The returned string is meant for debugging; for a stable serialized representation, use date.MarshalText, date.MarshalBinary, or date.Format with an explicit format string.
func (date Date) String() string {
	return time.Time(date).Format(time.DateOnly)
}

// Returns the duration date-value.
//
// # Parameters
//
//	value time.Time
//
// # Returns
//
//	duration time.Duration
//
// The duration [time.Duration] date-value.
//
// # Remarks
//
// If the result exceeds the maximum (or minimum) value that can be stored in a [time.Duration], the maximum (or minimum) duration will be returned.
// To compute date-value for a duration duration, use date.Add(-duration).
func (date Date) Sub(value time.Time) time.Duration {
	return time.Time(date).Sub(value)
}

// func (date Date) Truncate(value time.Duration) Date {
// 	return Date(time.Time(date).Truncate(value))
// }

// func (date Date) UTC() Date {
// 	return Date(time.Time(date).UTC().Truncate(24 * time.Hour))
// }

// Returns date as a Unix time, the number of seconds elapsed since January 1, 1970 UTC.
//
// # Returns
//
//	int64
//
// The number of seconds elapsed since January 1, 1970 UTC.
//
// # Remarks
//
// Unix-like operating systems often record time as a 32-bit count of seconds, but since the method here returns a 64-bit value it is valid for billions of years into the past or future.
func (date Date) Unix() int64 {
	return time.Time(date).Unix()
}

// Returns date as a Unix time, the number of microseconds elapsed since January 1, 1970 UTC.
//
// # Returns
//
//	int64
//
// The number of microseconds elapsed since January 1, 1970 UTC.
//
// # Remarks
//
// The result is undefined if the Unix time in microseconds cannot be represented by an int64 (a date before year -290307 or after year 294246).
func (date Date) UnixMicro() int64 {
	return time.Time(date).UnixMicro()
}

// Returns date as a Unix time, the number of milliseconds elapsed since January 1, 1970 UTC.
//
// # Returns
//
//	int64
//
// The number of milliseconds elapsed since January 1, 1970 UTC.
//
// # Remarks
//
// The result is undefined if the Unix time in milliseconds cannot be represented by an int64 (a date more than 292 million years before or after 1970).
func (date Date) UnixMilli() int64 {
	return time.Time(date).UnixMilli()
}

// Returns date as a Unix time, the number of nanoseconds elapsed since January 1, 1970 UTC.
//
// # Returns
//
//	int64
//
// The number of nanoseconds elapsed since January 1, 1970 UTC.
//
// # Remarks
//
// The result is undefined if the Unix time in nanoseconds cannot be represented by an int64 (a date before the year 1678 or after 2262).
func (date Date) UnixNano() int64 {
	return time.Time(date).UnixNano()
}

// Implements the [encoding.BinaryUnmarshaler] interface.
//
// # Parameters
//
//	data []byte
//
// Binary data.
//
// # Returns
//
//	err error
//
// Error when unmarshal problems, nil otherwise.
func (date *Date) UnmarshalBinary(data []byte) error {
	time := &time.Time{}
	err := time.UnmarshalBinary(data)
	if err != nil {
		return err
	}
	*date = Date(*time)
	return nil
}

// Implements the [encoding/json.Unmarshaler] interface.
//
// # Parameters
//
//	data []byte
//
// JSON data.
//
// # Returns
//
//	err error
//
// Error when unmarshal problems, nil otherwise.
//
// # Remarks
//
// The date must be a quoted string in the [time.DateOnly] format.
func (date *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// TODO(https://go.dev/issue/47353): Properly unescape a JSON string.
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("date.UnmarshalJSON: input is not a JSON string")
	}
	data = data[1 : len(data)-1]
	time, err := time.Parse(time.DateOnly, string(data))
	if err != nil {
		return err
	}
	*date = Date(time)
	return nil
}

// Implements the [encoding.TextUnmarshaler] interface.
//
// # Parameters
//
//	data []byte
//
// Text data.
//
// # Returns
//
//	err error
//
// Error when unmarshal problems, nil otherwise.
//
// # Remarks
//
// The date must be a string in the [time.DateOnly] format.
func (date *Date) UnmarshalText(data []byte) error {
	time, err := time.Parse(time.DateOnly, string(data))
	if err != nil {
		return err
	}
	*date = Date(time)
	return nil
}

// Returns the day of the week specified by date.
//
// # Returns
//
//	weekday time.Weekday
//
// The day of the week specified by date.
func (date Date) Weekday() time.Weekday {
	return time.Time(date).Weekday()
}

// Returns the year in which date occurs.
//
// # Returns
//
//	year int
//
// The year in which date occurs.
func (date Date) Year() int {
	return time.Time(date).Year()
}

// Returns the day of the year specified by date.
//
// # Returns
//
//	yearDay int
//
// The day of the year specified by date.
//
// # Remarks
//
// Returned value is in the range [1,365] for non-leap years, and [1,366] in leap years.
func (date Date) YearDay() int {
	return time.Time(date).YearDay()
}

// func (date Date) Zone() (name string, offset int) {
// 	return time.Time(date).Zone()
// }

// func (date Date) ZoneBounds(t time.Time) (start time.Time, end time.Time) {
// 	return time.Time(date).ZoneBounds()
// }

// Creates a new instance of the Date structure to the specified year, month, and day.
//
// # Parameters
//
//	year int
//
// The year (1 through 9999).
//
//	month time.Month
//
// The month (1 through 12).
//
//	day int
//
// The day (1 througthe number of days in month).
//
// # Returns
//
//	date Date
//
// A new instance of the Date structure to the specified year, month, and day.
func New(year int, month time.Month, day int) Date {
	return Date(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

// func FromTime(time time.Time) Date {
// 	year, month, day := time.Date()
// 	return New(year, month, day)
// }

// func (date Date) ToTime() time.Time {
// 	return time.Time(date)
// }

func Parse(layout string, value string) (Date, error) {
	t, err := time.Parse(layout, value)
	return Date(t.In(time.UTC).Truncate(24 * time.Hour)), err
}

// Returns the current local date.
//
// # Returns
//
//	date Date
//
// Current local date.
func Today() Date {
	year, month, day := time.Now().Date()
	return New(year, month, day)
}

// func ParseInLocation(layout string, value string, location *time.Location) (Date, error) {
// 	t, err := time.ParseInLocation(layout, value, location)
// 	return Date(t.Truncate(24 * time.Hour)), err
// }

// Returns the time elapsed since value.
//
// # Parameters
//
//	value time.Time
//
// The value from which the elapsed time is returned.
//
// # Returns
//
//	duration time.Duration
//
// The time elapsed since value.
//
// # Remarks
//
// It is shorthand for date.Today().Sub(value).
func Since(value time.Time) time.Duration {
	return Today().Sub(value)
}

// Returns the duration until value.
//
// # Parameters
//
//	value time.Time
//
// The value to which the duration is returned.
//
//	duration time.Duration
//
// The duration until value.
//
// # Remarks
//
// It is shorthand for value.Sub(date.Today()).
func Until(value time.Time) time.Duration {
	return value.Sub(time.Time(Today()))
}
