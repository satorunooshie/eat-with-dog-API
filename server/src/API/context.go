package API

import (
	"context"
	"strconv"
)

type (
	deviceVersion int
	userID        int
	//me            int
	sessionToken  int
	clientIP      int
	clientUA      int
	requestedTime int
	debug         int
)

var (
	deviceVersionKey deviceVersion
	userIDKey        userID
	//meKey            me
	sessionTokenKey sessionToken
	clientIPKey     clientIP
	clientUAKey     clientUA
	requestTimeKey  requestedTime
	debugKey        debug
)

func SafeUserIDString(ctx context.Context) string {
	userID, ok := ctx.Value(userIDKey).(int64)
	if !ok || userID == 0 {
		return "-"
	}
	return strconv.FormatInt(userID, 10)
}
