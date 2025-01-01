package time

import (
	"fmt"
	"strings"
	"time"
)

var TimeFormat = map[string]string{
	"yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
	"yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
	"yyyy-mm-dd hh":       "2006-01-02 15",
	"yyyy-mm-dd":          "2006-01-02",
	"yyyy-mm":             "2006-01",
	"mm-dd":               "01-02",
	"dd-mm-yy hh:mm:ss":   "02-01-06 15:04:05",

	"yyyy/mm/dd hh:mm:ss": "2006/01/02 15:04:05",
	"yyyy/mm/dd hh:mm":    "2006/01/02 15:04",
	"yyyy/mm/dd hh":       "2006/01/02 15",
	"yyyy/mm/dd":          "2006/01/02",
	"yyyy/mm":             "2006/01",
	"mm/dd":               "01/02",
	"dd/mm/yy hh:mm:ss":   "02/01/06 15:04:05",

	"yyyymmdd": "20060102",
	"mmddyy":   "010206",
	"yyyy":     "2006",
	"yy":       "06",
	"mm":       "01",
	"hh:mm:ss": "15:04:05",
	"hh:mm":    "15:04",
	"mm:ss":    "04:05",
}

// Time holds time.Time and provides additional methods
type Time struct{ time.Time }

// NewTime returns a new Time with current time
func NewTime() *Time                   { return &Time{time.Now()} }
func NewTimeFromTm(tm time.Time) *Time { return &Time{tm} }
func NewTimeFromUnix(ts int64) *Time   { return &Time{time.Unix(ts, 0)} }

// NewTimeFormStr converts string to DateTime
//
//	Usage:
//		NewTimeFormStr("2024-09-10 23:24:25", "yyyy-mm-dd hh:mm:ss")
//		NewTimeFormStr("2024-09-10 23:24:25", "yyyy-mm-dd hh:mm:ss", time.Local.String())
//		NewTimeFormStr("2024-09-10", "yyyy-mm-dd")
//		NewTimeFormStr("10-09-24 23:24:25", "dd-mm-yy hh:mm:ss")
func NewTimeFormStr(str, format string, timezone ...string) (*Time, error) {
	tf, ok := TimeFormat[strings.ToLower(format)]
	if !ok {
		return nil, fmt.Errorf("invalid format: %s", format)
	}

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return nil, err
		}

		dt, err := time.ParseInLocation(tf, str, loc)
		if err != nil {
			return nil, err
		}

		return NewTimeFromTm(dt), nil
	}

	dt, err := time.Parse(tf, str)
	if err != nil {
		return nil, err
	}

	return NewTimeFromTm(dt), nil
}

// FormatTimeToStr converts time to string
func (t *Time) FormatTimeToStr(format string, timezone ...string) string {
	tf, ok := TimeFormat[strings.ToLower(format)]
	if !ok {
		return ""
	}

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return ""
		}

		return t.In(loc).Format(tf)
	}

	return t.Format(tf)
}

// AddMinute adds or subs minute
func (t *Time) AddMinute(minute int64) { t.Time = t.Add(time.Minute * time.Duration(minute)) }
func (t *Time) AddHour(hour int64)     { t.Time = t.Add(time.Hour * time.Duration(hour)) }
func (t *Time) AddDay(day int64)       { t.Time = t.AddDate(0, 0, int(day)) }
func (t *Time) AddMonth(month int64)   { t.Time = t.AddDate(0, int(month), 0) }
func (t *Time) AddYear(year int64)     { t.Time = t.AddDate(int(year), 0, 0) }

// DateStr returns format "2006-01-02" of current time
func (t *Time) DateStr() string { return t.Format(time.DateOnly) }

// TimeStr returns format "15:04:05" of current time
func (t *Time) TimeStr() string { return t.Format(time.TimeOnly) }

// DateTimeStr returns format "2006-01-02 15:04:05" of current datetime
func (t *Time) DateTimeStr() string { return t.Format(time.DateTime) }

// TodayStartTimeStr returns the start time of today, format: yyyy-mm-dd 00:00:00.
func (t *Time) TodayStartTimeStr() string { return t.DateStr() + " 00:00:00" }

// TodayEndTimeStr returns the end time of today, format: yyyy-mm-dd 23:59:59.
func (t *Time) TodayEndTimeStr() string { return t.DateStr() + " 23:59:59" }

// Timestamp returns current second timestamp.
func (t *Time) Timestamp(timezone ...string) int64 {
	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}

		return t.In(loc).Unix()
	}

	return t.Unix()
}

// TimestampMill return current millisecond timestamp
func (t *Time) TimestampMilli(timezone ...string) int64 {
	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}

		return t.In(loc).UnixMilli()
	}

	return t.UnixMilli()
}

// TimestampMill return current microsecond timestamp
func (t *Time) TimestampMicro(timezone ...string) int64 {
	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}

		return t.In(loc).UnixNano() * int64(time.Nanosecond) / int64(time.Microsecond)
	}

	return t.UnixNano() * int64(time.Nanosecond) / int64(time.Microsecond)
}

// TimestampMill return current nanosecond timestamp
func (t *Time) TimestampNano(timezone ...string) int64 {
	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}

		return t.In(loc).UnixNano()
	}

	return t.UnixNano()
}

// ZeroHourTimestamp return timestamp of zero hour (timestamp of 00:00).
func (t *Time) ZeroHourTimestamp() int64 {
	tm, err := time.Parse(time.DateOnly, t.DateStr())
	if err != nil {
		return 0
	}

	return tm.UTC().Unix() - 8*3600 // XXX 8*3600
}

// NightTimestamp returns timestamp of zero hour (timestamp of 23:59)
func (t *Time) NightTimestamp() int64 {
	return t.ZeroHourTimestamp() + 86400 - 1 // XXX 86400-1
}

// BeginOfMinute returns beginning minute time of day
func (t *Time) BeginOfMinute() time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute returns end minute time of day
func (t *Time) EndOfMinute() time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfHour returns beginning hour time of day
func (t *Time) BeginOfHour() time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour returns end hour time of day
func (t *Time) EndOfHour() time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfDay returns beginning day time of day
func (t *Time) BeginOfDay() time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns end day time of day
func (t *Time) EndOfDay() time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfWeek returns beginning week, default week begin from Sunday.
func (t *Time) BeginOfWeek(begFrom ...time.Weekday) time.Time {
	begFromWeek := time.Sunday
	if len(begFrom) > 0 {
		begFromWeek = begFrom[0]
	}

	year, month, day := t.AddDate(0, 0, int(begFromWeek-t.Weekday())).Date()
	begOfWeek := time.Date(year, month, day, 0, 0, 0, 0, t.Location())

	if begOfWeek.After(t.Time) {
		begOfWeek = begOfWeek.AddDate(0, 0, -7)
	}

	return begOfWeek
}

// EndOfWeek returns ending week, default week end to Saturday.
func (t *Time) EndOfWeek(endWith ...time.Weekday) time.Time {
	endWithWeek := time.Saturday
	if len(endWith) > 0 {
		endWithWeek = endWith[0]
	}

	year, month, day := t.AddDate(0, 0, int(endWithWeek-t.Weekday())).Date()
	endOfWeek := time.Date(year, month, day, 23, 59, 59,
		int(time.Second-time.Nanosecond), t.Location())

	if endOfWeek.Before(t.Time) {
		endOfWeek = endOfWeek.AddDate(0, 0, 7)
	}

	return endOfWeek
}

// BeginOfMonth returns current date of month begin time.
func (t *Time) BeginOfMonth() time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns current date of month end time.
func (t *Time) EndOfMonth() time.Time {
	return t.BeginOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfYear returns current date of year begin time.
func (t *Time) BeginOfYear() time.Time {
	// year, _, _ := t.Date()
	return time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear returns current date of year end time.
func (t *Time) EndOfYear() time.Time {
	return t.BeginOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// DayOfYears returns which day of the year. firstday: 0
func (t *Time) DayOfYear() int {
	y_, m_, d_ := t.Date()
	firstDay := time.Date(y_, 1, 1, 0, 0, 0, 0, t.Location())
	currDay := time.Date(y_, m_, d_, 0, 0, 0, 0, t.Location())

	return int(currDay.Sub(firstDay).Hours() / 24)
}

// Weekend judge the day is weekend or not.
func (t *Time) Weekend() bool {
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

// IsLeapYear check if param year is leap year or not.
func (t *Time) IsLeapYear() bool {
	year := t.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// BetweenSeconds returns the number of seconds between two times.
func (t *Time) BetweenSeconds(other time.Time) int64 { return other.Unix() - t.Unix() }

// TraceFuncCost return the func costed time(milliseconds).
//
//	usage: `defer TraceFuncCost()`
func TraceFuncCost() func() {
	t1 := time.Now()

	return func() {
		cost := time.Since(t1)
		// XXX Print <--> log
		fmt.Println("Cost time: \t", cost)
	}
}
