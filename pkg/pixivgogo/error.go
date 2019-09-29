package pixivgogo

import (
	"encoding/json"
)

type PixivError struct {
	HasError bool          `json:"has_error"`
	Details  *ErrorDetails `json:"errors"`
}

type ErrorDetails struct {
	SystemError *SystemError `json:"system"`
}

type SystemError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *PixivError) Error() string {
	jsonBytes, err := json.Marshal(e)
	if err == nil {
		return ""
	}
	return string(jsonBytes)
}
