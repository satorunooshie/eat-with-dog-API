package time

import (
	"time"
)

type Time time.Time

func New(t time.Time) Time {
	return Time(t)
}

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t *Time) IsZero() bool {
	return t.Time().IsZero()
}

func (t *Time) DateFormat() string {
	return DateFormat(t.Time())
}

func (t *Time) DateTimeFormat() string {
	return DateTimeFormat(t.Time())
}

func (t *Time) ISO8601Format() string {
	return ISO8601Format(t.Time())
}

func (t *Time) ISO8601UTCFormat() string {
	return ISO8601UTCFormat(t.Time())
}

func (t *Time) ISO8601ExtendedFormat() string {
	return ISO8601ExtendedFormat(t.Time())
}

var fixedTime Time

func Now() time.Time {
	if !fixedTime.IsZero() {
		return fixedTime.Time()
	}
	return time.Now()
}

func NowDateTime() string {
	return DateTimeFormat(Now())
}

func DeltaMinutes(n time.Duration) time.Time {
	return Now().Add(n * time.Minute)
}

func RecordNow() int64 {
	return Now().UnixNano() / int64(1000)
}

func CreateTimeFromRecordTime(rt int64) time.Time {
	return time.Unix(rt/int64(1000000), 0)
}

func CreateRecordTimeFromTime(t time.Time) int64 {
	return t.UnixNano() / int64(1000)
}

func CreateSecondTimeFromTime(t time.Time) int64 {
	return t.UnixNano() / int64(1000000000) * 1000000
}

func CreateTimeFromRecordTimeContainNanoSecond(rt int64) time.Time {
	return time.Unix(rt/int64(100000), rt%int64(1000000)*1000)
}

func ConvertTimeToDateString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return DateFormat(t.In(AppLocation()))
}

func ConvertTimeToTimezoneString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return DateTimeFormat(t.In(AppLocation()))
}

func ConvertTimeToDateTimeMinuteFormatString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return DateTimeMinuteFormat(t.In(AppLocation()))
}

func EqualDate(t1, t2 time.Time) bool {
	return TimeDayStart(t1).Equal(TimeDayStart(t2))
}

func EqualWeekday(t1, t2 time.Time) bool {
	return t1.Weekday() == t2.Weekday()
}

func IsToday(rt int64) bool {
	today := TimeDayStart(Now())
	day := TimeDayStart(CreateTimeFromRecordTime(rt))
	return today.Equal(day)
}

func IsTodayFromAppLocation(t time.Time) bool {
	today := TimeDayStart(Now().In(AppLocation()))
	day := TimeDayStart(t.In(AppLocation()))
	return today.Equal(day)
}

func IsTodayFromUTC(t time.Time) bool {
	today := TimeDayStart(Now().In(AppLocation()))
	day := TimeDayStart(t.In(AppLocation()))
	return today.Equal(day)
}

func TimeLastMonthStart(utcTime time.Time) time.Time {
	t := utcTime.In(AppLocation())
	t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())

	return t.AddDate(0, -1, 0).UTC()
}

func TimeLastMonthEnd(t time.Time) time.Time {
	t = time.Date(t.Year(), t.Month(), 1, 23, 59, 59, 999999999, t.Location())
	return t.AddDate(0, 0, -1)
}

func TimeHourStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

func TimeHourEnd(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 59, 59, 999999999, t.Location())
}

func TimeMinStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

func TimeMinEnd(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 59, 999999999, t.Location())
}

func Time10MinStart(t time.Time) time.Time {
	min := (t.Minute() / 10) * 10 // 1の位を0にする
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), min, 0, 0, t.Location())
}

func Time10MinEnd(t time.Time) time.Time {
	min := (t.Minute() / 10) * 10 // 1の位を0にする
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), min+9, 59, 999999999, t.Location())
}

func FutureSecond(dt time.Time, i int) time.Time {
	dt = dt.In(AppLocation())
	return dt.Add(time.Duration(i) * time.Second)
}

func FutureSecondOneSec(dt time.Time) time.Time {
	return FutureSecond(dt, 1)
}

func FutureMinute(dt time.Time, i int) time.Time {
	dt = dt.In(AppLocation())
	return dt.Add(time.Duration(i) * time.Minute)
}

func FutureHour(dt time.Time, i int) time.Time {
	dt = dt.In(AppLocation())
	return dt.Add(time.Duration(i) * time.Hour)
}

func FutureDay(dt time.Time, i int) time.Time {
	dt = dt.In(AppLocation())
	return dt.AddDate(0, 0, i)
}

func FutureMonth(dt time.Time, i int) time.Time {
	dt = dt.In(AppLocation())
	return dt.AddDate(0, i, 0)
}

func FutureYear(dt time.Time, i int) time.Time {
	dt = dt.In(AppLocation())
	return dt.AddDate(i, 0, 0)
}

func TimeFutureMinute(i int) time.Time {
	dt := Now().In(AppLocation())
	return dt.Add(time.Duration(i) * time.Minute)
}

func TimeFutureHour(i int) time.Time {
	dt := Now().In(AppLocation())
	return dt.Add(time.Duration(i) * time.Hour)
}

func TimeFutureDay(i int) time.Time {
	dt := Now().In(AppLocation())
	return dt.AddDate(0, 0, i)
}

func TimeFutureMonth(i int) time.Time {
	dt := Now().In(AppLocation())
	return dt.AddDate(0, i, 0)
}

func TimeFutureYear(i int) time.Time {
	dt := Now().In(AppLocation())
	return dt.AddDate(i, 0, 0)
}

func PastSecond(dt time.Time, i int) time.Time {
	return FutureSecond(dt, i*-1)
}

func PastMinute(dt time.Time, i int) time.Time {
	return FutureMinute(dt, i*-1)
}

func PastHour(dt time.Time, i int) time.Time {
	return FutureHour(dt, i*-1)
}

func PastDay(dt time.Time, i int) time.Time {
	return FutureDay(dt, i*-1)
}

func PastMonth(dt time.Time, i int) time.Time {
	return FutureMonth(dt, i*-1)
}

func PastYear(dt time.Time, i int) time.Time {
	return FutureYear(dt, i*-1)
}

func TimePastMinute(i int) time.Time {
	return TimeFutureMinute(i * -1)
}

func TimePastHour(i int) time.Time {
	return TimeFutureHour(i * -1)
}

func TimePastDay(i int) time.Time {
	return TimeFutureDay(i * -1)
}

func TimePastMonth(i int) time.Time {
	return TimeFutureMonth(i * -1)
}

func TimePastYear(i int) time.Time {
	return TimeFutureYear(i * -1)
}

func DurationDays(dt time.Time) int {
	dt = TimeDayStart(dt)
	now := TimeDayStart(Now())
	delta := dt.Sub(now)
	days := int(delta / time.Hour / 24)
	if days < 0 {
		days = 0
	}
	return days
}

func GetElapsedSecond(dt time.Time) int64 {
	return int64(since(dt) / time.Second)
}

func GetElapsedMinute(dt time.Time) int64 {
	return int64(since(dt) / time.Minute)
}

func GetElapsedHour(dt time.Time) int {
	return int(since(dt) / time.Hour)
}

func GetElapsedDay(dt time.Time) int {
	return int(Now().Sub(dt)/time.Hour) / 24
}

func GetAppNow() time.Time {
	return Now().In(AppLocation())
}

type SwitchingTime struct{ time.Time }

func (st *SwitchingTime) ReviseToYesterdayIfNeed(hour int) time.Time {
	cp := SwitchingTime{Time: st.In(AppLocation())}
	if st.Hour() < hour {
		return time.Date(
			st.Year(),
			st.Month(),
			st.Day()-1, // Yesterday
			st.Hour(),
			st.Minute(),
			st.Second(),
			st.Nanosecond(),
			time.Local,
		)
	}
	return cp.Time
}

func RemoveNano(dt time.Time) time.Time {
	return time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), dt.Second(), 0, dt.Location())
}

func GetFirstDayOfNextMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month()+1, 1, 12, 0, 0, 0, AppLocation())
}

func DateRangeByAge(age int) (from, to time.Time) {
	dt := Now()
	return dt.AddDate(-(age + 1), 0, 0), dt.AddDate(-age, 0, 1)
}

func (t Time) BetweenHour(start, end int) bool {
	hour := t.Time().Hour()
	return start <= hour && hour <= end
}

func since(t time.Time) time.Duration {
	return Now().Sub(t)
}

func TimeDayStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func TimeDayStartToday() time.Time {
	return TimeDayStart(Now())
}

func TimeDayEnd(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

func TimeDayEndToday() time.Time {
	return TimeDayEnd(Now())
}

////////////
// Debug //
//////////

// alias
var FixedTime = TheWorld

func TheWorld(dio func(), args ...interface{}) time.Time {
	const definedTime = "2006-01-02 15:04:05"
	if len(args) == 0 {
		return TheWorld(dio, definedTime, true)
	}
	switch t := args[0].(type) {
	case bool:
		return TheWorld(dio, definedTime, t)
	case string, time.Time:
		if len(args) == 1 {
			return TheWorld(dio, t, true)
		}
		if isTheWorld, _ := args[1].(bool); isTheWorld {
			setFixedTimeParseValue(t)
		}
		if dio != nil {
			dio()
		}
		tm := fixedTime.Time()
		resetFixedTime()
		return tm
	}
	return fixedTime.Time()
}

func setFixedTime(t time.Time) {
	fixedTime = New(t)
}

func resetFixedTime() {
	fixedTime = Time{}
}

func setFixedTimeParseValue(value interface{}) {
	switch value := value.(type) {
	case string:
		if t, err := ParseDateTimeFormat(value); err == nil {
			fixedTime = New(t)
		}
	case time.Time:
		fixedTime = New(value)
	}
}
