package time

import (
	"time"

	"github.com/satorunooshie/eat-with-dog-API/server/errors"
)

var (
	appLocation *time.Location
	appTimeZone string
	timezones   = map[string]string{
		"jp": "Asia/Tokyo",
	}
)

func Init(countryCode string) error {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return errors.Errorf("failed to load UTC: %v", err)
	}
	time.Local = location

	if _, ok := timezones[countryCode]; !ok {
		return errors.Errorf("invalid country code: %s", countryCode)
	}
	appTimeZone = timezones[countryCode]

	tmp, err := time.LoadLocation(appTimeZone)
	if err != nil {
		return errors.Errorf("failed to load app timezone location: %v", err)
	}
	if tmp == nil {
		return errors.New("the configuration of timezone might be wrong")
	}

	appLocation = tmp
	return nil
}

func AppLocation() *time.Location {
	return appLocation
}

func AppTimeZone() string {
	return appTimeZone
}
