package time

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateFormat(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{
			name: "success",
			t:    Now(),
			want: Now().Format("2006-01-02"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DateFormat(tt.t)
			assert.Equal(t, tt.want, got)
			if got := DateFormat(tt.t); got != tt.want {
				t.Errorf("DateFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateFormatForHuman(t *testing.T) {
	const (
		japan     = "Asia/Tokyo"
		today     = "2006-01-02"
		yesterday = "2006-01-01"
		tomorrow  = "2006-01-03"
	)

	defer func() { appLocation = nil }()

	tests := []struct {
		timezone string
		utc      string
		want     string
	}{
		{
			timezone: japan,
			utc:      "2006-01-01 14:59:59",
			want:     yesterday,
		},
		{
			timezone: japan,
			utc:      "2006-01-01 15:00:00",
			want:     yesterday,
		},
		{
			timezone: japan,
			utc:      "2006-01-01 19:59:59", // border
			want:     yesterday,
		},
		{
			timezone: japan,
			utc:      "2006-01-01 20:00:00",
			want:     today,
		},
		{
			timezone: japan,
			utc:      "2006-01-02 00:00:00",
			want:     today,
		},
		{
			timezone: japan,
			utc:      "2006-01-02 14:59:59",
			want:     today,
		},
		{
			timezone: japan,
			utc:      "2006-01-02 15:00:00",
			want:     today,
		},
		{
			timezone: japan,
			utc:      "2006-01-02 19:59:59", // border
			want:     today,
		},
		{
			timezone: japan,
			utc:      "2006-01-02 20:00:00",
			want:     tomorrow,
		},
	}
	for _, tt := range tests {
		target := fmt.Sprintf("%+v", tt)
		dt, err := ParseDateTimeFormat(tt.utc)
		assert.NoError(t, err, target)

		location, err := time.LoadLocation(tt.timezone)
		assert.NoError(t, err, target)

		appLocation = location

		assert.Equal(t, tt.want, DateFormatForHuman(dt), target)
	}
}
