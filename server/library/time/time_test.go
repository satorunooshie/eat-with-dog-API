package time

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEqualWeekday(t *testing.T) {
	t1, _ := ParseDateTimeFormat("2016-04-15 01:02:03")
	t21, _ := ParseDateTimeFormat("2016-04-01 02:03:04")
	t22, _ := ParseDateTimeFormat("2016-04-02 03:04:05")
	tests := []struct {
		t1   time.Time
		t2   time.Time
		want bool
	}{
		{
			t1:   t1,
			t2:   t21,
			want: true,
		},
		{
			t1:   t1,
			t2:   t22,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := EqualWeekday(tt.t1, tt.t2); got != tt.want {
				t.Errorf("EqualWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetElapsedDay(t *testing.T) {
	_ = Init("jp")
	tests := []struct {
		dt     time.Time
		want   int
		format string
	}{
		{
			dt:     TimePastDay(30),
			want:   30,
			format: "2006-01-02 15:04:05",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GetElapsedDay(tt.dt); got != tt.want {
				t.Errorf("GetElapsedDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetElapsedHour(t *testing.T) {
	_ = Init("jp")
	tests := []struct {
		dt     time.Time
		want   int
		format string
	}{
		{
			dt:     TimePastHour(3),
			want:   3,
			format: "2006-01-02 15:04:05",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GetElapsedHour(tt.dt); got != tt.want {
				t.Errorf("GetElapsedHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetElapsedMinute(t *testing.T) {
	_ = Init("jp")
	tests := []struct {
		dt     time.Time
		want   int64
		format string
	}{
		{
			dt:     TimePastMinute(30),
			want:   int64(30),
			format: "2006-01-02 15:04:05",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GetElapsedMinute(tt.dt); got != tt.want {
				t.Errorf("GetElapsedMinute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTodayFromAppLocation(t *testing.T) {
	_ = Init("jp")
	todayStart := TimeDayStart(Now().In(AppLocation()))
	todayEnd := TimeDayEnd(Now().In(AppLocation()))
	tests := []struct {
		t    time.Time
		want bool
	}{
		{
			t:    todayStart.Add(1 * time.Second),
			want: true,
		},
		{
			t:    todayStart.Add(-1 * time.Second),
			want: false,
		},
		{
			t:    todayEnd.Add(-1 * time.Second),
			want: true,
		},
		{
			t:    todayEnd.Add(1 * time.Second),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsTodayFromAppLocation(tt.t); got != tt.want {
				t.Errorf("IsTodayFromAppLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNow(t *testing.T) {
	now1 := Now()
	assert.False(t, now1.IsZero())
	time.Sleep(10 * time.Millisecond)
	now2 := Now()
	nano1, nano2 := now1.UnixNano(), now2.UnixNano()
	assert.NotEqual(t, now1, now2)
	assert.NotEqual(t, nano1, nano2)

	const (
		layout  = "2006-01-02 15:04:05 -07:00"
		timeStr = "2006-01-02 23:12:34 +00:00"
	)

	tm, _ := time.Parse(layout, timeStr)
	setFixedTime(tm)

	assert.Equal(t, tm, Now())
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, tm, Now())
}

func TestPastDay(t *testing.T) {
	_ = Init("jp")
	type args struct {
		dt time.Time
		i  int
	}
	dt := Now().In(AppLocation())
	tests := []struct {
		args args
		want time.Time
	}{
		{
			args: args{
				dt: dt,
				i:  3,
			},
			want: dt.AddDate(0, 0, -3),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := PastDay(tt.args.dt, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PastDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwitchingTime_ReviseToYesterdayIfNeed(t *testing.T) {
	_ = Init("jp")
	tests := []struct {
		now           string
		switchingTime int
		want          string
	}{
		{
			now:           "2016-07-10 01:00:00",
			switchingTime: 4,
			want:          "2016-07-09",
		},
		{
			now:           "2016-07-10 05:00:00",
			switchingTime: 4,
			want:          "2016-07-10",
		},
		{
			now:           "2016-07-10 05:00:00",
			switchingTime: 11,
			want:          "2016-07-09",
		},
		{
			now:           "2016-07-10 11:00:00",
			switchingTime: 11,
			want:          "2016-07-10",
		},
		{
			now:           "2016-04-01 10:00:00",
			switchingTime: 11,
			want:          "2016-03-31",
		},
	}
	for _, tt := range tests {
		now, _ := time.ParseInLocation(DateTimeLayout, tt.now, AppLocation())
		TheWorld(func() {
			s := &SwitchingTime{
				Time: GetAppNow(),
			}
			st := s.ReviseToYesterdayIfNeed(tt.switchingTime)
			assert.Equal(t, DateFormat(st), tt.want)
		}, DateTimeFormat(now.UTC()))
	}
}

func TestTimeLastMonthEnd(t *testing.T) {
	_ = Init("jp")
	tests := []struct {
		utcTime time.Time
		want    time.Time
	}{
		{
			utcTime: time.Date(
				2015, 11, 30, 15, 0, 0, 0, time.UTC,
			),
			want: time.Date(
				2015, 10, 31, 23, 59, 59, 999999999, time.UTC,
			),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := TimeLastMonthEnd(tt.utcTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeLastMonthEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeLastMonthStart(t *testing.T) {
	_ = Init("jp")
	tests := []struct {
		utcTime time.Time
		want    time.Time
	}{
		{
			utcTime: time.Date(
				2016, 1, 21, 22, 33, 44, 1, time.UTC,
			),
			want: time.Date(
				2015, 11, 30, 15, 0, 0, 0, time.UTC,
			),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := TimeLastMonthStart(tt.utcTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeLastMonthStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime_BetweenHour(t *testing.T) {
	_ = Init("jp")
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		t    Time
		args args
		want bool
	}{
		{
			t: New(
				time.Date(2017, 11, 28, 0, 0, 0, 0, AppLocation()),
			),
			args: args{
				start: 0,
				end:   1,
			},
			want: true,
		},
		{
			t: New(
				time.Date(2017, 11, 28, 1, 0, 0, 0, AppLocation()),
			),
			args: args{
				start: 0,
				end:   1,
			},
			want: true,
		},
		{
			t: New(
				time.Date(2017, 11, 28, 2, 0, 0, 0, AppLocation()),
			),
			args: args{
				start: 0,
				end:   1,
			},
			want: false,
		},
		{
			t: New(
				time.Date(2017, 11, 28, 22, 0, 0, 0, AppLocation()),
			),
			args: args{
				start: 22,
				end:   23,
			},
			want: true,
		},
		{
			t: New(
				time.Date(2017, 11, 28, 23, 0, 0, 0, AppLocation()),
			),
			args: args{
				start: 22,
				end:   23,
			},
			want: true,
		},
		{
			t: New(
				time.Date(2017, 11, 28, 0, 0, 0, 0, AppLocation()),
			),
			args: args{
				start: 22,
				end:   23,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.t.BetweenHour(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("BetweenHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime(t *testing.T) {
	tm := Time{}
	assert.True(t, tm.IsZero())

	fixed := FixedTime(nil)
	assert.Equal(t, "2006-01-02 15:04:05", DateTimeFormat(fixed))
	assert.NotEqual(t, "2006-01-02 15:04:05", DateTimeFormat(Now()))

	TheWorld(func() {
		now := Time(Now())
		assert.False(t, now.IsZero())
		assert.Equal(t, "2006-01-02", now.DateFormat())
		assert.Equal(t, "2006-01-02 15:04:05", now.DateTimeFormat())
		assert.Equal(t, "2006-01-02T15:04:05+00:00", now.ISO8601ExtendedFormat())
	})
}
