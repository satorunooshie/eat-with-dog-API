package time

import (
	"time"
)

const (
	HourMinuteLayout           = "15:04"
	TimeLayout                 = "15:04:05"
	MonthDayLayout             = "1/2"
	DateLayout                 = "2006-01-02"
	DateSlashLayout            = "2006/01/02"
	DateNoSeparatorsLayout     = "20060102"
	DateTimeMinuteLayout       = "2006-01-02 15:04"
	DateTimeLayout             = "2006-01-02 15:04:05"
	DateTimeMinuteSlashLayout  = "2006/01/02 15:04"
	DateTimeSlashLayout        = "2006/01/02 15:04:05"
	ISO8601Layout              = "2006-01-02T15:04:05"
	ISO8601ExtendedLayout      = "2006-01-02T15:04:05-07:00"
	ISO8601ExtendedSlashLayout = "2006/01/02T15:04:05-07:00"
	ISO8601UTCLayout           = "2006-01-02T15:04:05Z"
	ISO8601NoTZLayout          = "2006-01-02T15:04:05"
)

func HourMinuteFormat(t time.Time) string {
	return t.Format(HourMinuteLayout)
}

func TimeFormat(t time.Time) string {
	return t.Format(TimeLayout)
}

func TimeFormatInApplication(t time.Time) string {
	return t.In(AppLocation()).Format(TimeLayout)
}

func MonthDayFormat(t time.Time) string {
	return t.Format(MonthDayLayout)
}

func DateFormat(t time.Time) string {
	return t.Format(DateLayout)
}

func NoSeparatorsFormat(t time.Time) string {
	return t.Format(DateNoSeparatorsLayout)
}

func ISO8601(t time.Time) string {
	return t.Format(ISO8601Layout)
}

func ParseDateFormat(s string) (time.Time, error) {
	createdAt, err := time.Parse(DateLayout, s)
	if err != nil {
		return time.Parse(DateSlashLayout, s)
	}
	return createdAt, nil
}

func ParseDateFormatInAppLocation(s string) (time.Time, error) {
	createdAt, err := time.ParseInLocation(DateLayout, s, AppLocation())
	if err != nil {
		return time.ParseInLocation(DateSlashLayout, s, AppLocation())
	}
	return createdAt, nil
}

func MustParseDateFormatInAppLocation(s string) time.Time {
	t, _ := ParseDateFormatInAppLocation(s)
	return t
}

func ParseDateTimeFormatInAppLocation(s string) (time.Time, error) {
	createdAt, err := time.ParseInLocation(DateTimeLayout, s, AppLocation())
	if err != nil {
		return time.ParseInLocation(DateTimeSlashLayout, s, AppLocation())
	}
	return createdAt, nil
}

func MustParseDateTimeFormatInAppLocation(s string) time.Time {
	t, _ := ParseDateTimeFormatInAppLocation(s)
	return t
}

func MustParseDateFormat(s string) time.Time {
	t, _ := ParseDateFormat(s)
	return t
}

func DateFormatAppLocation(t time.Time) string {
	return DateFormat(t.In(AppLocation()))
}

// 現地時間換算で朝5:00から翌日4:59までを当日とした日付文字列を返却する
func DateFormatForHuman(t time.Time) string {
	const dayStart = 5
	t = t.In(AppLocation())
	return FutureHour(t, -dayStart).Format(DateLayout)
}

func DateTimeFormat(t time.Time) string {
	return t.Format(DateTimeLayout)
}

func DateTiemFormatAppLocation(t time.Time) string {
	return DateTimeFormat(t.In(AppLocation()))
}

func ParseDateTimeFormat(s string) (time.Time, error) {
	createdAt, err := time.Parse(DateTimeLayout, s)
	if err != nil {
		return time.Parse(DateTimeSlashLayout, s)
	}
	return createdAt, nil
}

func MustParseDateTimeFormat(s string) time.Time {
	t, _ := ParseDateTimeFormat(s)
	return t
}

func DateTimeMinuteFormat(t time.Time) string {
	return t.Format(DateTimeMinuteLayout)
}

func DateTimeMinuteFomartAppLocation(t time.Time) string {
	return DateTimeMinuteFormat(t.In(AppLocation()))
}

func ParseDateTimeMinuteFormatAppLocation(s string) (time.Time, error) {
	createdAt, err := time.ParseInLocation(DateTimeMinuteLayout, s, AppLocation())
	if err != nil {
		return time.ParseInLocation(DateTimeMinuteLayout, s, AppLocation())
	}
	return createdAt, nil
}

func MustParseDateTimeMinuteFormatAppLocation(s string) time.Time {
	t, _ := ParseDateTimeMinuteFormatAppLocation(s)
	return t
}

func ISO8601ExtendedFormat(t time.Time) string {
	return t.Format(ISO8601ExtendedLayout)
}

func ISO8601Format(t time.Time) string {
	return t.Format(ISO8601Layout)
}

func ParseISO8601Format(s string) (time.Time, error) {
	createdAt, err := time.Parse(ISO8601Layout, s)
	if err != nil {
		return createdAt, err
	}
	return createdAt, nil
}

func ISO8601UTCFormat(t time.Time) string {
	return t.Format(ISO8601UTCLayout)
}

func ParseISO8601UTCFormat(s string) (time.Time, error) {
	createdAt, err := time.Parse(ISO8601UTCLayout, s)
	if err != nil {
		return createdAt, err
	}
	return createdAt, nil
}

func ParseISO8601ExtendedFormat(s string) (time.Time, error) {
	createdAt, err := time.Parse(ISO8601ExtendedLayout, s)
	if err != nil {
		return time.Parse(ISO8601ExtendedSlashLayout, s)
	}
	return createdAt, nil
}

func MustParseISO8601ExtendedFormat(s string) time.Time {
	t, _ := ParseISO8601ExtendedFormat(s)
	return t
}

func ParseISO8601FormatUTC(s string) (time.Time, error) {
	return time.Parse(ISO8601UTCLayout, s)
}
