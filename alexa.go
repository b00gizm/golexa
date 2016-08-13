package golexa

import (
	"encoding/json"
)

func Default() *Alexa {
	return Init([]Handler{})
}

func Init(handlers []Handler) *Alexa {
	var launchHandler LaunchHandler
	var sessionEndedHandler SessionEndedHandler

	intentHandlers := make([]IntentHandler, 0)

	for _, handler := range handlers {
		if h, ok := handler.(LaunchHandler); ok {
			launchHandler = h
		}

		if h, ok := handler.(IntentHandler); ok {
			intentHandlers = append(intentHandlers, h)
		}

		if h, ok := handler.(SessionEndedHandler); ok {
			sessionEndedHandler = h
		}
	}

	return &Alexa{
		launchHandler:       launchHandler,
		intentHandlers:      intentHandlers,
		sessionEndedHandler: sessionEndedHandler,
	}
}

type Alexa struct {
	launchHandler       LaunchHandler
	intentHandlers      []IntentHandler
	sessionEndedHandler SessionEndedHandler
}

func (a *Alexa) Process(msg json.RawMessage) (*Response, error) {
	var event Event

	err := json.Unmarshal(msg, &event)
	if err != nil {
		return nil, err
	}

	if a.launchHandler != nil && event.Request.Type == REQUEST_TYPE_LAUNCH {
		response := a.launchHandler.HandleLaunch(
			a,
			&event.Request,
			&event.Session,
		)

		if response != nil {
			return response, nil
		}
	}

	if len(a.intentHandlers) > 0 && event.Request.Type == REQUEST_TYPE_INTENT {
		var response *Response
		for _, handler := range a.intentHandlers {
			r := handler.HandleIntent(
				a,
				&event.Request.Intent,
				&event.Request,
				&event.Session,
			)

			if r != nil {
				response = r
				break
			}
		}

		if response != nil {
			return response, nil
		}
	}

	if a.sessionEndedHandler != nil && event.Request.Type == REQUEST_TYPE_ENDED {
		response := a.sessionEndedHandler.HandleSessionEnded(
			a,
			&event.Request,
			&event.Session,
		)

		if response != nil {
			return response, nil
		}
	}

	return newEmptyResponse(), nil
}

func (a *Alexa) OnLaunch(handler func(a *Alexa, req *Request, session *Session) *Response) {
	a.launchHandler = LaunchHandlerFunc(handler)
}

func (a *Alexa) OnIntent(handler func(a *Alexa, intent *Intent, req *Request, session *Session) *Response) {
	a.intentHandlers = []IntentHandler{
		IntentHandlerFunc(handler),
	}
}

func (a *Alexa) OnSessionEnded(handler func(a *Alexa, req *Request, session *Session) *Response) {
	a.sessionEndedHandler = SessionEndedHandlerFunc(handler)
}

func (a *Alexa) Response() *Response {
	return newEmptyResponse()
}
