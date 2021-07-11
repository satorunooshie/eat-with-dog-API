package time

import (
	"testing"
	"time"
)

func TestYearDateFormatFromDateString(t *testing.T) {
	tests := []struct {
		name  string
		japan *Japan
		s     string
		want  string
	}{
		{
			name:  "success",
			japan: NewJapan(),
			s:     "2020-01-01",
			want:  "2020年1月1日",
		},
		{
			name:  "success slash",
			japan: NewJapan(),
			s:     "2020/01/01",
			want:  "2020年1月1日",
		},
		{
			name:  "failure",
			japan: NewJapan(),
			s:     "2020-01",
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.japan.YearDateFormatFromDateString(tt.s); got != tt.want {
				t.Errorf("YearDateFormatFromDateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateFormatWithWeekday(t *testing.T) {
	tests := []struct {
		name  string
		japan *Japan
		time  time.Time
		want  string
	}{
		{
			name:  "success",
			japan: NewJapan(),
			time: func() time.Time {
				t, _ := time.Parse(DateLayout, "2020-01-01")
				return t
			}(),
			want: "1月1日（水）",
		},
		{
			name:  "success slash",
			japan: NewJapan(),
			time: func() time.Time {
				t, _ := time.Parse(DateSlashLayout, "2020/01/01")
				return t
			}(),
			want: "1月1日（水）",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.japan.DateFormatWithWeekday(tt.time); got != tt.want {
				t.Errorf("DateFormatWithWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearDateFormatWithWeekday(t *testing.T) {
	tests := []struct {
		name  string
		japan *Japan
		time  time.Time
		want  string
	}{
		{
			name:  "success",
			japan: NewJapan(),
			time: func() time.Time {
				t, _ := time.Parse(DateLayout, "2020-01-01")
				return t
			}(),
			want: "2020年1月1日（水）",
		},
		{
			name:  "success slash",
			japan: NewJapan(),
			time: func() time.Time {
				t, _ := time.Parse(DateSlashLayout, "2020/01/01")
				return t
			}(),
			want: "2020年1月1日（水）",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.japan.YearDateFormatWithWeekday(tt.time); got != tt.want {
				t.Errorf("YearDateFormatWithWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTimeFormatWithWeekday(t *testing.T) {
	tests := []struct {
		name  string
		japan *Japan
		time  time.Time
		want  string
	}{
		{
			name:  "success",
			japan: NewJapan(),
			time: func() time.Time {
				t, _ := time.Parse(DateTimeLayout, "2020-01-01 00:00:00")
				return t
			}(),
			want: "1月1日（水）09:00",
		},
		{
			name:  "success slash",
			japan: NewJapan(),
			time: func() time.Time {
				t, _ := time.Parse(DateTimeSlashLayout, "2020/01/01 00:00:00")
				return t
			}(),
			want: "1月1日（水）09:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.japan.DateTimeFormatWithWeekday(tt.time); got != tt.want {
				t.Errorf("DateTimeFormatWithWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearDateTimeFormatWithWeekday(t *testing.T) {
	tests := []struct {
		name  string
		japan *Japan
		time  time.Time
		want  string
	}{
		{
			name:  "success",
			japan: NewJapan(),
			time: func() time.Time {
				t, _ := time.Parse(DateTimeLayout, "2020-01-01 00:00:00")
				return t
			}(),
			want: "2020年1月1日（水）09:00",
		},
		{
			name:  "success slash",
			japan: NewJapan(),
			time: func() time.Time {
				t, _ := time.Parse(DateTimeSlashLayout, "2020/01/01 00:00:00")
				return t
			}(),
			want: "2020年1月1日（水）09:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.japan.YearDateTimeFormatWithWeekday(tt.time); got != tt.want {
				t.Errorf("YearDateTimeFormatWithWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}
