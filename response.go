package golexa

const (
	OUTPUT_SPEECH_TYPE_PLAIN = "PlainText"
	OUTPUT_SPEECH_TYPE_SSML  = "SSML"

	CARD_TYPE_SIMPLE       = "Simple"
	CARD_TYPE_STANDARD     = "Standard"
	CARD_TYPE_LINK_ACCOUNT = "LinkAccount"
)

func newEmptyResponse() *Response {
	return &Response{
		Version: "1.0",
		Response: innerResponse{
			ShouldEndSession: true,
			OutputSpeech:     nil,
		},
	}
}

type Response struct {
	Version           string            `json:"version"`
	SessionAttributes SessionAttributes `json:"sessionAttributes,omitempty"`
	Response          innerResponse     `json:"response"`
}

type innerResponse struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	ShouldEndSession bool          `json:"shouldEndSession"`
}

type OutputSpeech struct {
	Type  string `json:"type"`
	Text  string `json:"text,omitempty"`
	SSML  string `json:"ssml,omitempty"`
}

type Card struct {
	Type      string     `json:"type"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content,omitempty"`
	Text      string     `json:"text,omitempty"`
	ImageUrls *imageUrls `json:"image,omitempty"`
}

type imageUrls struct {
	SmallImageUrl string `json:"smallImageUrl"`
	LargeImageUrl string `json:"largeImageUrl"`
}

type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech,omitempty"`
}

func (r *Response) AddPlainTextSpeech(text string) *Response {
	r.Response.OutputSpeech = &OutputSpeech{
		Type: OUTPUT_SPEECH_TYPE_PLAIN,
		Text: text,
	}

	return r
}

func (r *Response) AddSSMLSpeech(ssml string) *Response {
	r.Response.OutputSpeech = &OutputSpeech{
		Type: OUTPUT_SPEECH_TYPE_SSML,
		SSML: ssml,
	}

	return r
}

func (r *Response) AddSimpleCard(title, content string) *Response {
	r.Response.Card = &Card{
		Type:    CARD_TYPE_SIMPLE,
		Title:   title,
		Content: content,
	}

	return r
}

func (r *Response) AddStandardCard(title, text, smallImageUrl, largeImageUrl string) *Response {
	r.Response.Card = &Card{
		Type:  CARD_TYPE_STANDARD,
		Title: title,
		Text:  text,
		ImageUrls: &imageUrls{
			SmallImageUrl: smallImageUrl,
			LargeImageUrl: largeImageUrl,
		},
	}

	return r
}

func (r *Response) AddLinkAccountCard() *Response {
	r.Response.Card = &Card{
		Type: CARD_TYPE_LINK_ACCOUNT,
	}

	return r
}

func (r *Response) AddPlainTextReprompt(text string) *Response {
	r.Response.Reprompt = &Reprompt{
		OutputSpeech: OutputSpeech{
			Type: OUTPUT_SPEECH_TYPE_PLAIN,
			Text: text,
		},
	}

	return r
}

func (r *Response) AddSSMLReprompt(ssml string) *Response {
	r.Response.Reprompt = &Reprompt{
		OutputSpeech: OutputSpeech{
			Type: OUTPUT_SPEECH_TYPE_SSML,
			SSML: ssml,
		},
	}

	return r
}

func (r *Response) AddSessionAttributes(attributes SessionAttributes) *Response {
	r.SessionAttributes = attributes

	return r
}

func (r *Response) KeepSessionAlive() *Response {
	r.Response.ShouldEndSession = false

	return r
}
