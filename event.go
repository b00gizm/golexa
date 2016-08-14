package golexa

import (
	"time"
)

const (
	REQUEST_TYPE_LAUNCH = "LaunchRequest"
	REQUEST_TYPE_INTENT = "IntentRequest"
	REQUEST_TYPE_ENDED  = "SessionEndedRequest"

	REASON_USER_INITIATED         = "USER_INITIATED"
	REASON_ERROR                  = "ERROR"
	REASON_EXCEEDED_MAX_REPROMPTS = "EXCEEDED_MAX_REPROMPTS"
)

type Event struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Request Request `json:"request"`
}

type Session struct {
	Id          string            `json:"sessionId"`
	IsNew       bool              `json:"new"`
	Attributes  SessionAttributes `json:"attributes"`
	Application application       `json:"application"`
	User        user              `json:"user"`
}

type SessionAttributes map[string]interface{}

type application struct {
	Id string `json:"applicationId"`
}

type user struct {
	Id          string `json:"userId"`
	AccessToken string `json:"accessToken"`
}

type Request struct {
	Id        string    `json:"requestId"`
	Type      string    `json:"type"`
	Timestamp timestamp `json:"timestamp"`
	Intent    Intent    `json:"intent,omitempty"`
	Reason    string    `json:"reason,omitempty"`
}

type timestamp time.Time

func (t *timestamp) MarshalJSON() ([]byte, error) {
	s := time.Time(*t).Format(time.RFC3339Nano)

	return []byte(s), nil
}

func (t *timestamp) UnmarshalJSON(b []byte) error {
	time, err := time.Parse(`"`+time.RFC3339Nano+`"`, string(b))
	if err != nil {
		return err
	}

	*t = timestamp(time)

	return nil
}

func (t *timestamp) ToTime() time.Time {
	return time.Time(*t)
}

type Intent struct {
	Name string `json:"name"`
	Slots map[string]Slot `json:"slots"`
}

type Slot struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
