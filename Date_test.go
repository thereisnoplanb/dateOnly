package date

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDate_Add(t *testing.T) {
	type args struct {
		value time.Duration
	}
	tests := []struct {
		name string
		date Date
		args args
		want time.Time
	}{
		{
			name: "Short duration",
			date: New(2000, time.January, 1),
			args: args{
				value: 8*time.Hour + 7*time.Minute + 6*time.Second,
			},
			want: time.Date(2000, time.January, 1, 8, 7, 6, 0, time.UTC),
		},
		{
			name: "Long duration",
			date: New(2000, time.January, 1),
			args: args{
				value: 24*time.Hour + 8*time.Hour + 7*time.Minute + 6*time.Second,
			},
			want: time.Date(2000, time.January, 2, 8, 7, 6, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.Add(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Date.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getPtr[T any](value T) *T {
	return &value
}

func TestDate_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		date     *Date
		args     args
		wantDate *Date
		wantErr  bool
	}{
		{
			name: "Unmarshalable",
			date: new(Date),
			args: args{
				data: []byte(`"2000-01-01"`),
			},
			wantDate: getPtr(New(2000, time.January, 1)),
			wantErr:  false,
		},
		{
			name: "Unmarshalable - null",
			date: new(Date),
			args: args{
				data: []byte(`null`),
			},
			wantDate: &Date{},
			wantErr:  false,
		},
		{
			name: "Not unmarshalable",
			date: new(Date),
			args: args{
				data: []byte(`"ala ma kota"`),
			},
			wantDate: &Date{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.date.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Date.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.date, tt.wantDate) {
				t.Errorf("Date.UnmarshalJSON() = %v, want %v", *tt.date, *tt.wantDate)
			}
		})
	}
}

func TestDate_AddDate(t *testing.T) {
	type args struct {
		years  int
		months int
		days   int
	}
	tests := []struct {
		name string
		date Date
		args args
		want Date
	}{
		{
			name: "Zero",
			date: New(2000, time.January, 1),
			args: args{
				years:  0,
				months: 0,
				days:   0,
			},
			want: New(2000, time.January, 1),
		},
		{
			name: "1 day",
			date: New(2000, time.January, 1),
			args: args{
				years:  0,
				months: 0,
				days:   1,
			},
			want: New(2000, time.January, 2),
		},
		{
			name: "1 month",
			date: New(2000, time.January, 1),
			args: args{
				years:  0,
				months: 1,
				days:   0,
			},
			want: New(2000, time.February, 1),
		},
		{
			name: "1 year",
			date: New(2000, time.January, 1),
			args: args{
				years:  1,
				months: 0,
				days:   0,
			},
			want: New(2001, time.January, 1),
		},
		{
			name: "1 year, 1 month, 1 day",
			date: New(2000, time.January, 1),
			args: args{
				years:  1,
				months: 1,
				days:   1,
			},
			want: New(2001, time.February, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.AddDate(tt.args.years, tt.args.months, tt.args.days); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Date.AddDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_After(t *testing.T) {
	type args struct {
		value Date
	}
	tests := []struct {
		name string
		date Date
		args args
		want bool
	}{
		{
			name: "Date is after value",
			date: New(2000, time.January, 1),
			args: args{
				value: New(1999, time.December, 31),
			},
			want: true,
		},
		{
			name: "Date is equal value - is not after",
			date: New(2000, time.January, 1),
			args: args{
				value: New(2000, time.January, 1),
			},
			want: false,
		},
		{
			name: "Date is before value - is not after",
			date: New(2000, time.January, 1),
			args: args{
				value: New(2000, time.January, 1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.After(tt.args.value); got != tt.want {
				t.Errorf("Date.After() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_AppendBinary(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		date    Date
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test",
			date: New(2000, time.January, 1),
			args: args{
				bytes: []byte("X"),
			},
			want:    []byte{88, 1, 0, 0, 0, 14, 175, 255, 58, 128, 0, 0, 0, 0, 255, 255},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.date.AppendBinary(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Date.AppendBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Date.AppendBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Before(t *testing.T) {
	type args struct {
		value Date
	}
	tests := []struct {
		name string
		date Date
		args args
		want bool
	}{
		{
			name: "Date is before value",
			date: New(1999, time.December, 31),
			args: args{
				value: New(2000, time.January, 1),
			},
			want: true,
		},
		{
			name: "Date is equal value - is not before",
			date: New(2000, time.January, 1),
			args: args{
				value: New(2000, time.January, 1),
			},
			want: false,
		},
		{
			name: "Date is after value - is not before",
			date: New(2000, time.January, 1),
			args: args{
				value: New(1999, time.December, 31),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.Before(tt.args.value); got != tt.want {
				t.Errorf("Date.Before() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Equal(t *testing.T) {
	type args struct {
		other Date
	}
	tests := []struct {
		name string
		date Date
		args args
		want bool
	}{
		{
			name: "Equal",
			date: New(2000, 1, 1),
			args: args{
				other: New(2000, 1, 1),
			},
			want: true,
		},
		{
			name: "Before",
			date: New(2000, 1, 1),
			args: args{
				other: New(1999, 12, 31),
			},
			want: false,
		},
		{
			name: "After",
			date: New(2000, 1, 1),
			args: args{
				other: New(2000, 1, 2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.Equal(tt.args.other); got != tt.want {
				t.Errorf("Date.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_IsZero(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want bool
	}{
		{
			name: "Is zero",
			date: Date{},
			want: true,
		},
		{
			name: "Is non-zero",
			date: New(2000, time.January, 1),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.IsZero(); got != tt.want {
				t.Errorf("Date.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Compare(t *testing.T) {
	type args struct {
		value Date
	}
	tests := []struct {
		name string
		date Date
		args args
		want int
	}{
		{
			name: "Before",
			date: New(2000, 1, 1),
			args: args{
				value: New(1999, 12, 31),
			},
			want: 1,
		},
		{
			name: "Equal",
			date: New(2000, 1, 1),
			args: args{
				value: New(2000, 1, 1),
			},
			want: 0,
		},
		{
			name: "After",
			date: New(2000, 1, 1),
			args: args{
				value: New(2000, 1, 2),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.Compare(tt.args.value); got != tt.want {
				t.Errorf("Date.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Date(t *testing.T) {
	tests := []struct {
		name      string
		date      Date
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{
			name:      "Composite",
			date:      New(2000, time.January, 2),
			wantYear:  2000,
			wantMonth: time.January,
			wantDay:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotYear, gotMonth, gotDay := tt.date.Deconstruct()
			if gotYear != tt.wantYear {
				t.Errorf("Date.Date() gotYear = %v, want %v", gotYear, tt.wantYear)
			}
			if !reflect.DeepEqual(gotMonth, tt.wantMonth) {
				t.Errorf("Date.Date() gotMonth = %v, want %v", gotMonth, tt.wantMonth)
			}
			if gotDay != tt.wantDay {
				t.Errorf("Date.Date() gotDay = %v, want %v", gotDay, tt.wantDay)
			}
		})
	}
}

func TestDate_Day(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want int
	}{
		{
			name: "Day1",
			date: New(2000, time.January, 1),
			want: 1,
		},
		{
			name: "Day29",
			date: New(2000, time.February, 29),
			want: 29,
		},
		{
			name: "Day30",
			date: New(2000, time.April, 30),
			want: 30,
		},
		{
			name: "Day31",
			date: New(2000, time.January, 31),
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.Day(); got != tt.want {
				t.Errorf("Date.Day() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Format(t *testing.T) {
	type args struct {
		layout string
	}
	tests := []struct {
		name string
		date Date
		args args
		want string
	}{
		{
			name: "DateOnly",
			date: New(2000, 1, 2),
			args: args{
				layout: "2006-01-02",
			},
			want: "2000-01-02",
		},
		{
			name: "US",
			date: New(2000, 1, 2),
			args: args{
				layout: "2006-02-01",
			},
			want: "2000-02-01",
		},
		{
			name: "YearOnly",
			date: New(2000, 1, 1),
			args: args{
				layout: "2006",
			},
			want: "2000",
		},
		{
			name: "YearMonthOnly",
			date: New(2000, 1, 2),
			args: args{
				layout: "2006-01",
			},
			want: "2000-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.Format(tt.args.layout); got != tt.want {
				t.Errorf("Date.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_GoString(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want string
	}{
		{
			name: "GoString - long year",
			date: New(2000, time.March, 5),
			want: "date.New(2000, time.March, 5)",
		},
		{
			name: "GoString - short year + time.UTC",
			date: New(1, time.December, 31),
			want: "date.New(1, time.December, 31)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.GoString(); got != tt.want {
				t.Errorf("Date.GoString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_ISOWeek(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		wantYear int
		wantWeek int
	}{
		{
			name:     "Last week previous year",
			date:     New(2000, time.January, 1),
			wantYear: 1999,
			wantWeek: 52,
		},
		{
			name:     "First week current year",
			date:     New(2000, time.January, 3),
			wantYear: 2000,
			wantWeek: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotYear, gotWeek := tt.date.ISOWeek()
			if gotYear != tt.wantYear {
				t.Errorf("Date.ISOWeek() gotYear = %v, want %v", gotYear, tt.wantYear)
			}
			if gotWeek != tt.wantWeek {
				t.Errorf("Date.ISOWeek() gotWeek = %v, want %v", gotWeek, tt.wantWeek)
			}
		})
	}
}

// func TestDate_In(t *testing.T) {
// 	type args struct {
// 		location *time.Location
// 	}
// 	NewYorkLocation, _ := time.LoadLocation("America/New_York")
// 	tests := []struct {
// 		name string
// 		date Date
// 		args args
// 		want Date
// 	}{
// 		{
// 			name: "Date in New York",
// 			date: New(2000, time.January, 1, time.UTC),
// 			args: args{
// 				location: NewYorkLocation,
// 			},
// 			want: New(1999, time.December, 31, NewYorkLocation),
// 		},
// 		{
// 			name: "Date in UTC",
// 			date: New(2000, time.January, 1, NewYorkLocation),
// 			args: args{
// 				location: time.UTC,
// 			},
// 			want: New(2000, time.January, 1, time.UTC),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := tt.date.In(tt.args.location)
// 			// if !got.Equal(tt.want) {
// 			// 	t.Errorf("Date.In() = %v, want %v", got, tt.want)
// 			// }
// 			fmt.Println(time.Time(tt.date))
// 			fmt.Println(time.Time(got))
// 			fmt.Println(time.Time(tt.want))
// 			fmt.Println(tt.date)
// 			fmt.Println(got)
// 			fmt.Println(tt.want)
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Date.In() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestDate_UTC(t *testing.T) {
// 	NewYorkLocation, _ := time.LoadLocation("America/New_York")
// 	tests := []struct {
// 		name string
// 		date Date
// 		want Date
// 	}{
// 		{
// 			name: "Test",
// 			date: New(2000, 1, 1, NewYorkLocation),
// 			want: New(2000, 1, 1, time.UTC),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := tt.date.UTC()
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Date.UTC() = %v, want %v", got, tt.want)
// 			}
// 			fmt.Println(time.Time(tt.date))
// 			fmt.Println(time.Time(got))
// 			fmt.Println(time.Time(tt.want))
// 		})
// 	}
// }

func TestToday(t *testing.T) {
	tests := []struct {
		name string
		want Date
	}{
		{
			name: "Test",
			want: Date(time.Now().UTC().Truncate(24 * time.Hour)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Today()
			fmt.Println(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Today() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Sub(t *testing.T) {
	type args struct {
		value time.Time
	}
	tests := []struct {
		name string
		date Date
		args args
		want time.Duration
	}{
		{
			name: "UTC",
			date: New(2000, time.January, 2),
			args: args{
				value: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			want: 24 * time.Hour,
		},
		{
			name: "Not UTC",
			date: New(2000, time.January, 2),
			args: args{
				value: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.FixedZone("My Zone", 3600)),
			},
			want: 24 * time.Hour,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.Sub(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Date.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
