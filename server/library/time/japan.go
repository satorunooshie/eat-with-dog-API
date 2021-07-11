package time

import (
	"fmt"
	"strings"
	"time"
)

const (
	JPDateLayout     = "1月2日"
	JPYearDateLayout = "2006年1月2日"
)

type Japan struct {
	weeks    map[time.Weekday]string
	location *time.Location
}

func NewJapan() *Japan {
	location, _ := time.LoadLocation("Asia/Tokyo")
	return &Japan{
		weeks: map[time.Weekday]string{
			time.Sunday:    "日",
			time.Monday:    "月",
			time.Tuesday:   "火",
			time.Wednesday: "水",
			time.Thursday:  "木",
			time.Friday:    "金",
			time.Saturday:  "土",
		},
		location: location,
	}
}

func (j *Japan) Location() *time.Location {
	return j.location
}

func (j *Japan) DateFormat(t time.Time) string {
	t = t.In(j.location)
	return t.Format(JPDateLayout)
}

func (j *Japan) YearDateFormat(t time.Time) string {
	t = t.In(j.location)
	return t.Format(JPYearDateLayout)
}

func (j *Japan) YearDateFormatFromDateString(s string) string {
	if _, err := ParseDateFormat(s); err != nil {
		return ""
	}
	d := make([]string, 0, 3)
	d = strings.Split(s, "-")
	if len(d) != 3 {
		d = strings.Split(s, "/")
		if len(d) != 3 {
			return ""
		}
	}
	return fmt.Sprintf("%s年%s月%s日", d[0], j.trimLeftZero(d[1]), j.trimLeftZero(d[2]))
}

func (j *Japan) DateFormatWithWeekday(t time.Time) string {
	t = t.In(j.location)
	return fmt.Sprintf("%s（%s）", t.Format(JPDateLayout), j.weeks[t.Weekday()])
}

func (j *Japan) YearDateFormatWithWeekday(t time.Time) string {
	t = t.In(j.location)
	return fmt.Sprintf("%s（%s）", t.Format(JPYearDateLayout), j.weeks[t.Weekday()])
}

func (j *Japan) DateTimeFormatWithWeekday(t time.Time) string {
	t = t.In(j.location)
	return fmt.Sprintf("%s（%s）%s", t.Format(JPDateLayout), j.weeks[t.Weekday()], HourMinuteFormat(t))
}

func (j *Japan) YearDateTimeFormatWithWeekday(t time.Time) string {
	t = t.In(j.location)
	return fmt.Sprintf("%s（%s）%s", t.Format(JPYearDateLayout), j.weeks[t.Weekday()], HourMinuteFormat(t))
}

func (j *Japan) trimLeftZero(s string) string {
	return strings.TrimLeft(s, "0")
}
