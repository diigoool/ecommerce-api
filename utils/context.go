package utils

import (
	"context"
)

func GetUserID(ctx context.Context) (uint, bool) {
	id, ok := ctx.Value("user_id").(int)

	return uint(id), ok
}

func GetRequestID(ctx context.Context) string {

	requestID, ok := ctx.Value(
		RequestIDKey,
	).(string)

	if !ok {
		return ""
	}

	return requestID

}
