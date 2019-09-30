package pixivgogo

import (
	"encoding/json"
)

type AuthError struct {
	HasError bool              `json:"has_error"`
	Details  *AuthErrorDetails `json:"errors"`
}

type AuthErrorDetails struct {
	SystemError *SystemError `json:"system,omitempty"`
}

type SystemError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (a *AuthError) Error() string {
	jsonBytes, err := json.Marshal(a)
	if err == nil {
		return ""
	}
	return string(jsonBytes)
}

type APIError struct {
	Details *APIErrorDetails `json:"error,omitempty"`
}

type APIErrorDetails struct {
	UserMessage string `json:"user_message"`
	Message     string `json:"message"`
	Reason      string `json:"reason"`
	// TODO Not sure how it would look like because I always see it being "{}"
	UserMessageDetails struct{} `json:"user_message_details"`
}

func (a *APIError) Error() string {
	jsonBytes, err := json.Marshal(a)
	if err == nil {
		return ""
	}
	return string(jsonBytes)
}
