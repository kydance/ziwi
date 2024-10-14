package datetime

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
	"yyyymmdd":            "20060102",
	"mmddyy":              "010206",
	"yyyy":                "2006",
	"yy":                  "06",
	"mm":                  "01",
	"hh:mm:ss":            "15:04:05",
	"hh:mm":               "15:04",
	"mm:ss":               "04:05",
}

// DateTime holds time.Time
type DateTime struct {
	tm time.Time
}

// NewDateTime
func NewDateTime() *DateTime {
	return &DateTime{tm: time.Now()}
}

// NewDateTimeFromTime
func NewDateTimeFromTime(t time.Time) *DateTime {
	return &DateTime{tm: t}
}

// NewDateTimeFormStr converts string to DateTime
//
//	Usage:
//		NewDateTimeFormStr("2024-09-10 23:24:25", "yyyy-mm-dd hh:mm:ss")
//		NewDateTimeFormStr("2024-09-10 23:24:25", "yyyy-mm-dd hh:mm:ss", time.Local.String())
//		NewDateTimeFormStr("2024-09-10", "yyyy-mm-dd")
//		NewDateTimeFormStr("10-09-24 23:24:25", "dd-mm-yy hh:mm:ss")
func NewDateTimeFormStr(str, format string, timezone ...string) (*DateTime, error) {
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

		return NewDateTimeFromTime(dt), nil
	}

	dt, err := time.Parse(tf, str)
	if err != nil {
		return nil, err
	}

	return NewDateTimeFromTime(dt), nil
}

// FormatTimeToStr converts time to string
func (d *DateTime) FormatTimeToStr(format string, timezone ...string) string {
	tf, ok := TimeFormat[strings.ToLower(format)]
	if !ok {
		return ""
	}

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return ""
		}

		return d.tm.In(loc).Format(tf)
	}

	return d.tm.Format(tf)
}

// AddMinute adds or subs minute
//
//	 Usage:
//			dt.AddMinute(1)
func (d *DateTime) AddMinute(minute int64) {
	d.tm = d.tm.Add(time.Minute * time.Duration(minute))
}

// AddHour adds or subs hour
func (d *DateTime) AddHour(hour int64) {
	d.tm = d.tm.Add(time.Hour * time.Duration(hour))
}

// AddDay adds or subs day
func (d *DateTime) AddDay(day int64) {
	d.tm = d.tm.Add(24 * time.Hour * time.Duration(day))
}

// AddYear adds or subs year(365 days)
func (d *DateTime) AddYear(year int64) {
	d.tm = d.tm.Add(365 * 24 * time.Hour * time.Duration(year))
}

// Date returns the year, month, and day.
func (d *DateTime) Date() (int, time.Month, int) {
	return d.tm.Date()
}

// DateStr returns format "2006-01-02" of current date
func (d *DateTime) DateStr() string {
	return d.tm.Format("2006-01-02")
}

// TimeStr returns format "15:04:05" of current time
func (d *DateTime) TimeStr() string {
	return d.tm.Format("15:04:05")
}

// DateTimeStr returns format "2006-01-02 15:04:05" of current datetime
func (d *DateTime) DateTimeStr() string {
	return d.tm.Format("2006-01-02 15:04:05")
}

// TodayStartTimeStr returns the start time of today, format: yyyy-mm-dd 00:00:00.
func (d *DateTime) TodayStartTimeStr() string {
	return d.DateStr() + " 00:00:00"
}

// TodayEndTimeStr returns the end time of today, format: yyyy-mm-dd 23:59:59.
func (d *DateTime) TodayEndTimeStr() string {
	return d.DateStr() + " 23:59:59"
}

// Timestamp returns current second timestamp.
func (d *DateTime) Timestamp(timezone ...string) int64 {
	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}

		return d.tm.In(loc).Unix()
	}

	return d.tm.Unix()
}

// TimestampMill return current millisecond timestamp
func (d *DateTime) TimestampMilli(timezone ...string) int64 {
	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}

		// return d.dt.In(loc).UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)
		return d.tm.In(loc).UnixMilli()
	}

	// return d.dt.UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)
	return d.tm.UnixMilli()
}

// TimestampMill return current microsecond timestamp
func (d *DateTime) TimestampMicro(timezone ...string) int64 {
	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}

		return d.tm.In(loc).UnixNano() * int64(time.Nanosecond) / int64(time.Microsecond)
	}

	return d.tm.UnixNano() * int64(time.Nanosecond) / int64(time.Microsecond)
}

// TimestampMill return current nanosecond timestamp
func (d *DateTime) TimestampNano(timezone ...string) int64 {
	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}

		return d.tm.In(loc).UnixNano()
	}

	return d.tm.UnixNano()
}

// ZeroHourTimestamp return timestamp of zero hour (timestamp of 00:00).
func (d *DateTime) ZeroHourTimestamp() int64 {
	dt := d.DateStr()
	t, _ := time.Parse("2006-01-02", dt)
	return t.UTC().Unix() - 8*3600 // XXX 8*3600
}

// NightTimestamp returns timestamp of zero hour (timestamp of 23:59)
func (d *DateTime) NightTimestamp() int64 {
	return d.ZeroHourTimestamp() + 86400 - 1 // XXX 86400-1
}

// BeginOfMinute returns beginning minute time of day
func (d *DateTime) BeginOfMinute() time.Time {
	y_, m_, d_ := d.tm.Date()
	return time.Date(y_, m_, d_, d.tm.Hour(), d.tm.Minute(), 0, 0, d.tm.Location())
}

// EndOfMinute returns end minute time of day
func (d *DateTime) EndOfMinute() time.Time {
	y_, m_, d_ := d.tm.Date()
	return time.Date(y_, m_, d_,
		d.tm.Hour(), d.tm.Minute(), 59,
		int(time.Second-time.Nanosecond), d.tm.Location())
}

// BeginOfHour returns beginning hour time of day
func (d *DateTime) BeginOfHour() time.Time {
	y_, m_, d_ := d.tm.Date()
	return time.Date(y_, m_, d_, d.tm.Hour(), 0, 0, 0, d.tm.Location())
}

// EndOfHour returns end hour time of day
func (d *DateTime) EndOfHour() time.Time {
	y_, m_, d_ := d.tm.Date()
	return time.Date(y_, m_, d_, d.tm.Hour(), 59, 59,
		int(time.Second-time.Nanosecond), d.tm.Location())
}

// BeginOfDay returns beginning day time of day
func (d *DateTime) BeginOfDay() time.Time {
	y_, m_, d_ := d.tm.Date()
	return time.Date(y_, m_, d_, 0, 0, 0, 0, d.tm.Location())
}

// EndOfDay returns end day time of day
func (d *DateTime) EndOfDay() time.Time {
	y_, m_, d_ := d.tm.Date()
	return time.Date(y_, m_, d_, 23, 59, 59,
		int(time.Second-time.Nanosecond), d.tm.Location())
}

// BeginOfWeek returns beginning week, default week begin from Sunday.
func (d *DateTime) BeginOfWeek(begFrom ...time.Weekday) time.Time {
	begFromWeek := time.Sunday
	if len(begFrom) > 0 {
		begFromWeek = begFrom[0]
	}

	y_, m_, d_ := d.tm.AddDate(0, 0, int(begFromWeek-d.tm.Weekday())).Date()
	begOfWeek := time.Date(y_, m_, d_, 0, 0, 0, 0, d.tm.Location())

	if begOfWeek.After(d.tm) {
		begOfWeek = begOfWeek.AddDate(0, 0, -7)
	}

	return begOfWeek
}

// EndOfWeek returns ending week, default week end to Saturday.
func (d *DateTime) EndOfWeek(endWith ...time.Weekday) time.Time {
	endWithWeek := time.Saturday
	if len(endWith) > 0 {
		endWithWeek = endWith[0]
	}

	y_, m_, d_ := d.tm.AddDate(0, 0, int(endWithWeek-d.tm.Weekday())).Date()
	endOfWeek := time.Date(y_, m_, d_, 23, 59, 59,
		int(time.Second-time.Nanosecond), d.tm.Location())

	if endOfWeek.Before(d.tm) {
		endOfWeek = endOfWeek.AddDate(0, 0, 7)
	}

	return endOfWeek
}

// BeginOfMonth returns current date of month begin time.
func (d *DateTime) BeginOfMonth() time.Time {
	y_, m_, _ := d.tm.Date()
	return time.Date(y_, m_, 1, 0, 0, 0, 0, d.tm.Location())
}

// EndOfMonth returns current date of month end time.
func (d *DateTime) EndOfMonth() time.Time {
	return d.BeginOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfYear returns current date of year begin time.
func (d *DateTime) BeginOfYear() time.Time {
	y_, _, _ := d.tm.Date()
	return time.Date(y_, time.January, 1, 0, 0, 0, 0, d.tm.Location())
}

// EndOfYear returns current date of year end time.
func (d *DateTime) EndOfYear() time.Time {
	return d.BeginOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// DayOfYears returns which day of the year. firstday: 0
func (d *DateTime) DayOfYear() int {
	y_, m_, d_ := d.tm.Date()
	firstDay := time.Date(y_, 1, 1, 0, 0, 0, 0, d.tm.Location())
	currDay := time.Date(y_, m_, d_, 0, 0, 0, 0, d.tm.Location())

	return int(currDay.Sub(firstDay).Hours() / 24)
}

// Weekend judge the day is weekend or not.
func (d *DateTime) Weekend() bool {
	return d.tm.Weekday() == time.Saturday || d.tm.Weekday() == time.Sunday
}

// IsLeapYear check if param year is leap year or not.
func (d *DateTime) IsLeapYear() bool {
	year := d.tm.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// BetweenSeconds returns the number of seconds between two times.
func BetweenSeconds(t1 time.Time, t2 time.Time) int64 { return t2.Unix() - t1.Unix() }

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
