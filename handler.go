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
