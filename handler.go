package golexa

type Handler interface{}

type LaunchHandler interface {
	Handler
	HandleLaunch(a *Alexa, req *Request, session *Session) *Response
}

type IntentHandler interface {
	Handler
	HandleIntent(a *Alexa, intent *Intent, req *Request, session *Session) *Response
}

type SessionEndedHandler interface {
	Handler
	HandleSessionEnded(a *Alexa, req *Request, session *Session) *Response
}

type LaunchHandlerFunc func(a *Alexa, req *Request, session *Session) *Response

func (fn LaunchHandlerFunc) HandleLaunch(a *Alexa, req *Request, session *Session) *Response {
	return fn(a, req, session)
}

type IntentHandlerFunc func(a *Alexa, intent *Intent, req *Request, session *Session) *Response

func (fn IntentHandlerFunc) HandleIntent(a *Alexa, intent *Intent, req *Request, session *Session) *Response {
	return fn(a, intent, req, session)
}

type SessionEndedHandlerFunc func(a *Alexa, req *Request, session *Session) *Response

func (fn SessionEndedHandlerFunc) HandleSessionEnded(a *Alexa, req *Request, session *Session) *Response {
	return fn(a, req, session)
}
