package utils

import "net/http"

func RespondError(w http.ResponseWriter, err error) {
	appErr := HandleError(err)

	Error(w, appErr.Status, appErr.Code, appErr.Message)
}
