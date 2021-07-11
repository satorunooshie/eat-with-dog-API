package time

import (
	"fmt"
	"time"
)

var week = map[string]time.Weekday{
	time.Sunday.String():    time.Sunday,
	time.Monday.String():    time.Monday,
	time.Tuesday.String():   time.Tuesday,
	time.Wednesday.String(): time.Wednesday,
	time.Thursday.String():  time.Thursday,
	time.Friday.String():    time.Friday,
	time.Saturday.String():  time.Saturday,
}

func StringToWeekday(s string) (time.Weekday, error) {
	if w, ok := week[s]; ok {
		return w, nil
	}
	return time.Sunday, fmt.Errorf("invalid days of the week: %s", s)
}
