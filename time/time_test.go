package time

import (
	"reflect"
	"testing"
	"time"
)

func TestNewDateTime(t *testing.T) {
	tm := NewTime()
	now := time.Now()

	if tm.Year() != now.Year() || tm.Month() != now.Month() || tm.Day() != now.Day() {
		t.Errorf("NewDateTime() failed, expected date to match current date")
	}
}

func TestSetTime(t *testing.T) {
	testTime := time.Date(2024, 11, 20, 12, 34, 56, 0, time.UTC)
	tm := NewTimeFromTm(testTime)

	if tm.Year() != testTime.Year() || tm.Month() != testTime.Month() || tm.Day() != testTime.Day() {
		t.Errorf("SetTime() failed, expected date to match provided date")
	}

	if tm.Hour() != testTime.Hour() || tm.Minute() != testTime.Minute() || tm.Second() != testTime.Second() {
		t.Errorf("SetTime() failed, expected time to match provided time")
	}
}

func TestNewDateTimeFormStr(t *testing.T) {
	type args struct {
		str      string
		format   string
		timezone []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Time
		wantErr bool
	}{
		{
			name: "yyyy-mm-dd hh:mm:ss",
			args: args{
				str:    "2024-09-10 23:24:25",
				format: "yyyy-mm-dd hh:mm:ss",
			},
			want: &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)},
		},
		{
			name: "yyyy-mm-dd hh:mm:ss with timezone",
			args: args{
				str:      "2024-09-10 23:24:25",
				format:   "yyyy-mm-dd hh:mm:ss",
				timezone: []string{"UTC"},
			},
			want: &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)},
		},
		{
			name: "yyyy-mm-dd hh:mm",
			args: args{
				str:    "2024-09-10 23:24",
				format: "yyyy-mm-dd hh:mm",
			},
			want: &Time{time.Date(2024, 9, 10, 23, 24, 0, 0, time.UTC)},
		},
		{
			name: "yyyy-mm-dd",
			args: args{
				str:    "2024-09-10",
				format: "yyyy-mm-dd",
			},
			want: &Time{time.Date(2024, 9, 10, 0, 0, 0, 0, time.UTC)},
		},
		{
			name: "dd-mm-yy hh:mm:ss",
			args: args{
				str:    "10-09-24 23:24:25",
				format: "dd-mm-yy hh:mm:ss",
			},
			want: &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)},
		},
		{
			name: "invalid format",
			args: args{
				str:    "2024-09-10 23:24:25",
				format: "invalid format",
			},
			wantErr: true,
		},
		{
			name: "invalid timezone",
			args: args{
				str:      "2024-09-10 23:24:25",
				format:   "yyyy-mm-dd hh:mm:ss",
				timezone: []string{"invalid timezone"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTimeFormStr(tt.args.str, tt.args.format, tt.args.timezone...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDateTimeFormStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && got != nil {
				if !got.Equal(tt.want.Time) {
					t.Errorf("NewDateTimeFormStr() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestDateTime_FormatTimeToStr(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}

	type args struct {
		format   string
		timezone []string
	}
	tests := []struct {
		name string
		dt   *Time
		args args
		want string
	}{
		{
			name: "yyyy-mm-dd hh:mm:ss",
			dt:   dt,
			args: args{
				format: "yyyy-mm-dd hh:mm:ss",
			},
			want: "2024-09-10 23:24:25",
		},
		{
			name: "yyyy-mm-dd hh:mm:ss with timezone",
			dt:   dt,
			args: args{
				format:   "yyyy-mm-dd hh:mm:ss",
				timezone: []string{"Asia/Shanghai"},
			},
			want: "2024-09-11 07:24:25",
		},
		{
			name: "yyyy-mm-dd hh:mm",
			dt:   dt,
			args: args{
				format: "yyyy-mm-dd hh:mm",
			},
			want: "2024-09-10 23:24",
		},
		{
			name: "yyyy-mm-dd",
			dt:   dt,
			args: args{
				format: "yyyy-mm-dd",
			},
			want: "2024-09-10",
		},
		{
			name: "dd-mm-yy hh:mm:ss",
			dt:   dt,
			args: args{
				format: "dd-mm-yy hh:mm:ss",
			},
			want: "10-09-24 23:24:25",
		},
		{
			name: "invalid format",
			dt:   dt,
			args: args{
				format: "invalid format",
			},
			want: "",
		},
		{
			name: "invalid timezone",
			dt:   dt,
			args: args{
				format:   "yyyy-mm-dd hh:mm:ss",
				timezone: []string{"invalid timezone"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.FormatTimeToStr(tt.args.format, tt.args.timezone...); got != tt.want {
				t.Errorf("DateTime.FormatTimeToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_AddMinute(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}
	dt.AddMinute(1)

	year, month, day := dt.Date()
	if year != 2024 || month != 9 || day != 10 {
		t.Errorf("DateTime.AddMinute() date error, got = %d-%d-%d, want = 2024-09-10", year, month, day)
	}

	hour, min, sec := dt.Hour(), dt.Minute(), dt.Second()
	if hour != 23 || min != 25 || sec != 25 {
		t.Errorf("DateTime.AddMinute() time error, got = %d:%d:%d, want = 23:25:25", hour, min, sec)
	}
}

func TestDateTime_AddHour(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}
	dt.AddHour(1)

	year, month, day := dt.Date()
	if year != 2024 || month != 9 || day != 11 {
		t.Errorf(
			"DateTime.AddHour() date error, got = %d-%d-%d, want = 2024-09-11",
			year,
			month,
			day,
		)
	}

	hour, min, sec := dt.Hour(), dt.Minute(), dt.Second()
	if hour != 0 || min != 24 || sec != 25 {
		t.Errorf("DateTime.AddHour() time error, got = %d:%d:%d, want = 00:24:25", hour, min, sec)
	}
}

func TestDateTime_AddDay(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}
	dt.AddDay(1)

	year, month, day := dt.Date()
	if year != 2024 || month != 9 || day != 11 {
		t.Errorf(
			"DateTime.AddDay() date error, got = %d-%d-%d, want = 2024-09-11",
			year,
			month,
			day,
		)
	}

	hour, min, sec := dt.Hour(), dt.Minute(), dt.Second()
	if hour != 23 || min != 24 || sec != 25 {
		t.Errorf("DateTime.AddDay() time error, got = %d:%d:%d, want = 23:24:25", hour, min, sec)
	}
}

func TestDateTime_AddYear(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}
	dt.AddYear(1)

	year, month, day := dt.Date()
	if year != 2025 || month != 9 || day != 10 {
		t.Errorf(
			"DateTime.AddYear() date error, got = %d-%d-%d, want = 2025-09-10",
			year,
			month,
			day,
		)
	}

	hour, min, sec := dt.Hour(), dt.Minute(), dt.Second()
	if hour != 23 || min != 24 || sec != 25 {
		t.Errorf("DateTime.AddYear() time error, got = %d:%d:%d, want = 23:24:25", hour, min, sec)
	}
}

func TestDateTime_Date(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}
	year, month, day := dt.Date()
	if year != 2024 || month != 9 || day != 10 {
		t.Errorf("DateTime.Date() error, got = %d-%d-%d, want = 2024-09-10", year, month, day)
	}
}

func TestDateTime_DateStr(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}
	want := "2024-09-10"
	if got := dt.DateStr(); got != want {
		t.Errorf("DateTime.DateStr() = %v, want %v", got, want)
	}
}

func TestDateTime_TimeStr(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}
	want := "23:24:25"
	if got := dt.TimeStr(); got != want {
		t.Errorf("DateTime.TimeStr() = %v, want %v", got, want)
	}
}

func TestDateTime_DateTimeStr(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}
	want := "2024-09-10 23:24:25"
	if got := dt.DateTimeStr(); got != want {
		t.Errorf("DateTime.DateTimeStr() = %v, want %v", got, want)
	}
}

func TestDateTime_TodayStartTimeStr(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}
	want := "2024-09-10 00:00:00"
	if got := dt.TodayStartTimeStr(); got != want {
		t.Errorf("DateTime.TodayStartTimeStr() = %v, want %v", got, want)
	}
}

func TestDateTime_TodayEndTimeStr(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}
	want := "2024-09-10 23:59:59"
	if got := dt.TodayEndTimeStr(); got != want {
		t.Errorf("DateTime.TodayEndTimeStr() = %v, want %v", got, want)
	}
}

func TestDateTime_Timestamp(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 0, time.UTC)}

	type args struct {
		timezone []string
	}
	tests := []struct {
		name string
		dt   *Time
		args args
		want int64
	}{
		{
			name: "Timestamp without timezone",
			dt:   dt,
			args: args{
				timezone: nil,
			},
			want: 1726010665,
		},
		{
			name: "Timestamp with timezone",
			dt:   dt,
			args: args{
				timezone: []string{"Asia/Shanghai"},
			},
			want: 1726010665,
		},
		{
			name: "Timestamp with timezone",
			dt:   dt,
			args: args{
				timezone: []string{time.UTC.String()},
			},
			want: 1726010665,
		},
		{
			name: "Timestamp with invalid timezone",
			dt:   dt,
			args: args{
				timezone: []string{"Invalid/Timezone"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.Timestamp(tt.args.timezone...); got != tt.want {
				t.Errorf("DateTime.Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_TimestampMilli(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 123456789, time.UTC)}

	type args struct {
		timezone []string
	}
	tests := []struct {
		name string
		dt   *Time
		args args
		want int64
	}{
		{
			name: "TimestampMilli without timezone",
			dt:   dt,
			args: args{
				timezone: nil,
			},
			want: 1726010665123,
		},
		{
			name: "TimestampMilli with timezone",
			dt:   dt,
			args: args{
				timezone: []string{"Asia/Shanghai"},
			},
			want: 1726010665123,
		},
		{
			name: "TimestampMilli with invalid timezone",
			dt:   dt,
			args: args{
				timezone: []string{"Invalid/Timezone"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.TimestampMilli(tt.args.timezone...); got != tt.want {
				t.Errorf("DateTime.TimestampMilli() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_TimestampMicro(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 123456789, time.UTC)}

	type args struct {
		timezone []string
	}
	tests := []struct {
		name string
		dt   *Time
		args args
		want int64
	}{
		{
			name: "TimestampMicro without timezone",
			dt:   dt,
			args: args{
				timezone: nil,
			},
			want: 1726010665123456,
		},
		{
			name: "TimestampMicro with timezone",
			dt:   dt,
			args: args{
				timezone: []string{"Asia/Shanghai"},
			},
			want: 1726010665123456,
		},
		{
			name: "TimestampMicro with invalid timezone",
			dt:   dt,
			args: args{
				timezone: []string{"Invalid/Timezone"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.TimestampMicro(tt.args.timezone...); got != tt.want {
				t.Errorf("DateTime.TimestampMicro() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_TimestampNano(t *testing.T) {
	dt := &Time{time.Date(2024, 9, 10, 23, 24, 25, 123456789, time.UTC)}

	type args struct {
		timezone []string
	}
	tests := []struct {
		name string
		dt   *Time
		args args
		want int64
	}{
		{
			name: "TimestampNano without timezone",
			dt:   dt,
			args: args{
				timezone: nil,
			},
			want: 1726010665123456789,
		},
		{
			name: "TimestampNano with timezone",
			dt:   dt,
			args: args{
				timezone: []string{"Asia/Shanghai"},
			},
			want: 1726010665123456789,
		},
		{
			name: "TimestampNano with invalid timezone",
			dt:   dt,
			args: args{
				timezone: []string{"Invalid/Timezone"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.TimestampNano(tt.args.timezone...); got != tt.want {
				t.Errorf("DateTime.TimestampNano() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_ZeroHourTimestamp(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want int64
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 0, time.FixedZone("CST", 8*3600))},
			want: 1725897600,
		},
		{
			name: "Test case 2",
			dt:   &Time{time.Date(2024, 12, 31, 23, 59, 59, 0, time.FixedZone("CST", 8*3600))},
			want: 1735574400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.ZeroHourTimestamp(); got != tt.want {
				t.Errorf("DateTime.ZeroHourTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_NightTimestamp(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want int64
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 0, time.FixedZone("CST", 8*3600))},
			want: 1725983999,
		},
		{
			name: "Test case 2",
			dt:   &Time{time.Date(2024, 12, 31, 23, 59, 59, 0, time.FixedZone("CST", 8*3600))},
			want: 1735660799,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.NightTimestamp(); got != tt.want {
				t.Errorf("DateTime.NightTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_BeginOfMinute(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want time.Time
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 9, 10, 10, 24, 0, 0, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.BeginOfMinute(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.BeginOfMinute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_EndOfMinute(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want time.Time
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 9, 10, 10, 24, 59, 999999999, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.EndOfMinute(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.EndOfMinute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_BeginOfHour(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want time.Time
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 9, 10, 10, 0, 0, 0, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.BeginOfHour(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.BeginOfHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_EndOfHour(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want time.Time
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 9, 10, 10, 59, 59, 999999999, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.EndOfHour(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.EndOfHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_BeginOfDay(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want time.Time
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 9, 10, 0, 0, 0, 0, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.BeginOfDay(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.BeginOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_EndOfDay(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want time.Time
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 9, 10, 23, 59, 59, 999999999, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.EndOfDay(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.EndOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_BeginOfWeek(t *testing.T) {
	tests := []struct {
		name     string
		dt       *Time
		begFrom  []time.Weekday
		wantTime time.Time
	}{
		{
			name:     "Test case 1: Default Sunday",
			dt:       &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			begFrom:  []time.Weekday{time.Sunday},
			wantTime: time.Date(2024, 9, 8, 0, 0, 0, 0, time.FixedZone("CST", 8*3600)),
		},
		{
			name:     "Test case 2: Monday",
			dt:       &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			begFrom:  []time.Weekday{time.Monday},
			wantTime: time.Date(2024, 9, 9, 0, 0, 0, 0, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTime := tt.dt.BeginOfWeek(tt.begFrom...); !reflect.DeepEqual(
				gotTime,
				tt.wantTime,
			) {
				t.Errorf("DateTime.BeginOfWeek() = %v, want %v", gotTime, tt.wantTime)
			}
		})
	}
}

func TestDateTime_EndOfWeek(t *testing.T) {
	tests := []struct {
		name     string
		dt       *Time
		endWith  []time.Weekday
		wantTime time.Time
	}{
		{
			name:     "Test case 1: Default Saturday",
			dt:       &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			endWith:  []time.Weekday{time.Saturday},
			wantTime: time.Date(2024, 9, 14, 23, 59, 59, 999999999, time.FixedZone("CST", 8*3600)),
		},
		{
			name:     "Test case 2: Friday",
			dt:       &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			endWith:  []time.Weekday{time.Friday},
			wantTime: time.Date(2024, 9, 13, 23, 59, 59, 999999999, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTime := tt.dt.EndOfWeek(tt.endWith...); !reflect.DeepEqual(gotTime, tt.wantTime) {
				t.Errorf("DateTime.EndOfWeek() = %v, want %v", gotTime, tt.wantTime)
			}
		})
	}
}

func TestDateTime_BeginOfMonth(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want time.Time
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 9, 1, 0, 0, 0, 0, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.BeginOfMonth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.BeginOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_EndOfMonth(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want time.Time
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 9, 30, 23, 59, 59, 999999999, time.FixedZone("CST", 8*3600)),
		},
		{
			name: "Test case 2: Leap year",
			dt:   &Time{time.Date(2024, 2, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 2, 29, 23, 59, 59, 999999999, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.EndOfMonth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.EndOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_BeginOfYear(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want time.Time
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.BeginOfYear(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.BeginOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_EndOfYear(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want time.Time
	}{
		{
			name: "Test case 1",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.FixedZone("CST", 8*3600)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.EndOfYear(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTime.EndOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_DayOfYear(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want int
	}{
		{
			name: "Test case 1: January 1st",
			dt:   &Time{time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
			want: 0,
		},
		{
			name: "Test case 2: February 29th (Leap Year)",
			dt:   &Time{time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)},
			want: 59,
		},
		{
			name: "Test case 3: December 31st",
			dt:   &Time{time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
			want: 365,
		},
		{
			name: "Test case 4: September 10th",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: 253,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.DayOfYear(); got != tt.want {
				t.Errorf("DateTime.DayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_Weekend(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want bool
	}{
		{
			name: "Test case 1: Saturday",
			dt:   &Time{time.Date(2024, 9, 14, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: true,
		},
		{
			name: "Test case 2: Sunday",
			dt:   &Time{time.Date(2024, 9, 15, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: true,
		},
		{
			name: "Test case 3: Monday",
			dt:   &Time{time.Date(2024, 9, 16, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.Weekend(); got != tt.want {
				t.Errorf("DateTime.Weekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime_IsLeapYear(t *testing.T) {
	tests := []struct {
		name string
		dt   *Time
		want bool
	}{
		{
			name: "Test case 1: Leap year",
			dt:   &Time{time.Date(2024, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: true,
		},
		{
			name: "Test case 2: Not leap year",
			dt:   &Time{time.Date(2023, 9, 10, 10, 24, 25, 123456789, time.FixedZone("CST", 8*3600))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.IsLeapYear(); got != tt.want {
				t.Errorf("DateTime.IsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetweenSeconds(t *testing.T) {
	t1, _ := NewTimeFormStr("2024-09-10 11:24:25", "yyyy-mm-dd hh:mm:ss")
	t2, _ := NewTimeFormStr("2024-09-10 12:24:25", "yyyy-mm-dd hh:mm:ss")

	got := t1.BetweenSeconds(t2.Time)
	want := int64(3600)

	if got != want {
		t.Errorf("BetweenSeconds() = %v, want %v", got, want)
	}
}

func TestNewTimeFromUnix(t *testing.T) {
	type args struct {
		ts int64
	}
	tests := []struct {
		name string
		args args
		want *Time
	}{
		{
			name: "Test case 1: Valid timestamp",
			args: args{ts: 1234567890},
			want: &Time{time.Unix(1234567890, 0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTimeFromUnix(tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTimeFromUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestA(t *testing.T) {
	t.Parallel()

	tma := NewTimeFromUnix(1732982400)
	t.Log(tma.Format(time.DateTime))
	tmb := NewTimeFromUnix(1735635074)
	t.Log(tmb.Format(time.DateTime))

	//
	tmc := NewTimeFromUnix(1733043074)
	t.Log(tmc.Format(time.DateTime))

	tmd := NewTimeFromUnix(1735635074)
	t.Log(tmd.Format(time.DateTime))

	// 1732982400,1735647420
	tme := NewTimeFromUnix(1732982400)
	t.Log(tme.Format(time.DateTime))

	tmf := NewTimeFromUnix(1735647420)
	t.Log(tmf.Format(time.DateTime))
}
