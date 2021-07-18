package log

import (
	"regexp"
)

var (
	regTokenURI = regexp.MustCompile(`token=[^&]+`)
	emptyToken  = "token="
)

func FilterTokenFromURI(uri string) string {
	return regTokenURI.ReplaceAllString(uri, emptyToken)
}
