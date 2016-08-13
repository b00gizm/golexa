package golexa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponse_AddPlainTextSpeech(t *testing.T) {
	r := newEmptyResponse().AddPlainTextSpeech("hello world")

	assert.Equal(t, "1.0", r.Version, "should have correct version")
	assert.Equal(t, true, r.Response.ShouldEndSession, "should end session by default")

	assert.Equal(t, OUTPUT_SPEECH_TYPE_PLAIN, r.Response.OutputSpeech.Type, "should have correct type")
	assert.Equal(t, "hello world", r.Response.OutputSpeech.Text, "should have correct text")
	assert.Empty(t, r.Response.OutputSpeech.SSML, "should not have SSML set")
}

func TestResponse_AddSSMLSpeech(t *testing.T) {
	r := newEmptyResponse().AddSSMLSpeech("<speak>hello world.</speak>")

	assert.Equal(t, "1.0", r.Version, "should have correct version")
	assert.Equal(t, true, r.Response.ShouldEndSession, "should end session by default")

	assert.Equal(t, OUTPUT_SPEECH_TYPE_SSML, r.Response.OutputSpeech.Type, "should have correct type")
	assert.Equal(t, "<speak>hello world.</speak>", r.Response.OutputSpeech.SSML, "should have correct SSML")
	assert.Empty(t, r.Response.OutputSpeech.Text, "should not have text set")
}

func TestResponse_AddSimpleCard(t *testing.T) {
	r := newEmptyResponse().AddSimpleCard("greeting", "hello world")

	assert.Equal(t, "1.0", r.Version, "should have correct version")
	assert.Equal(t, true, r.Response.ShouldEndSession, "should end session by default")

	assert.Equal(t, CARD_TYPE_SIMPLE, r.Response.Card.Type, "should have correct type")
	assert.Equal(t, "greeting", r.Response.Card.Title, "should have correct title")
	assert.Equal(t, "hello world", r.Response.Card.Content, "should have correct content")
	assert.Empty(t, r.Response.Card.Text, "should not have text set")
	assert.Empty(t, r.Response.Card.ImageUrls, "should not have image set")
}

func TestResponse_AddStandardCard(t *testing.T) {
	r := newEmptyResponse().AddStandardCard(
		"greeting",
		"hello world",
		"http://cdn.example.org/image_small.jpg",
		"http://cdn.example.org/image_large.jpg",
	)

	assert.Equal(t, "1.0", r.Version, "should have correct version")
	assert.Equal(t, true, r.Response.ShouldEndSession, "should end session by default")

	assert.Equal(t, CARD_TYPE_STANDARD, r.Response.Card.Type, "should have correct type")
	assert.Equal(t, "greeting", r.Response.Card.Title, "should have correct title")
	assert.Equal(t, "hello world", r.Response.Card.Text, "should have correct text")
	assert.Equal(
		t,
		"http://cdn.example.org/image_small.jpg",
		r.Response.Card.ImageUrls.SmallImageUrl,
		"should have correct small image URL",
	)
	assert.Equal(
		t,
		"http://cdn.example.org/image_large.jpg",
		r.Response.Card.ImageUrls.LargeImageUrl,
		"should have correct large image URL",
	)
	assert.Empty(t, r.Response.Card.Content, "should not have content set")
}

func TestResponse_AddLinkAccountCard(t *testing.T) {
	r := newEmptyResponse().AddLinkAccountCard()

	assert.Equal(t, "1.0", r.Version, "should have correct version")
	assert.Equal(t, true, r.Response.ShouldEndSession, "should end session by default")

	assert.Equal(t, CARD_TYPE_LINK_ACCOUNT, r.Response.Card.Type, "should have correct type")
	assert.Empty(t, r.Response.Card.Title, "should not have title set")
	assert.Empty(t, r.Response.Card.Content, "should not have content set")
	assert.Empty(t, r.Response.Card.Text, "should not have text set")
	assert.Empty(t, r.Response.Card.ImageUrls, "should not have image set")
}

func TestResponse_AddPlainTextReprompt(t *testing.T) {
	r := newEmptyResponse().AddPlainTextReprompt("hello world")

	assert.Equal(t, "1.0", r.Version, "should have correct version")
	assert.Equal(t, true, r.Response.ShouldEndSession, "should end session by default")

	assert.Equal(t, OUTPUT_SPEECH_TYPE_PLAIN, r.Response.Reprompt.OutputSpeech.Type, "should have correct type")
	assert.Equal(t, "hello world", r.Response.Reprompt.OutputSpeech.Text, "should have correct text")
	assert.Empty(t, r.Response.Reprompt.OutputSpeech.SSML, "should not have SSML set")
}

func TestResponse_AddSSMLReprompt(t *testing.T) {
	r := newEmptyResponse().AddSSMLReprompt("<speak>hello world.</speak>")

	assert.Equal(t, "1.0", r.Version, "should have correct version")
	assert.Equal(t, true, r.Response.ShouldEndSession, "should end session by default")

	assert.Equal(t, OUTPUT_SPEECH_TYPE_SSML, r.Response.Reprompt.OutputSpeech.Type, "should have correct type")
	assert.Equal(t, "<speak>hello world.</speak>", r.Response.Reprompt.OutputSpeech.SSML, "should have correct SSML")
	assert.Empty(t, r.Response.Reprompt.OutputSpeech.Text, "should not have text set")
}

func TestResponse_AddSessionAttributes(t *testing.T) {
	attributes := SessionAttributes{
		"foo": map[string]string{
			"bar": "baz",
		},
	}

	r := newEmptyResponse().AddSessionAttributes(attributes)

	assert.Equal(t, "1.0", r.Version, "should have correct version")
	assert.Equal(t, true, r.Response.ShouldEndSession, "should end session by default")
	assert.Equal(t, attributes, r.SessionAttributes, "should have correct session attributes")
}

func TestResponse_KeepSessionAlive(t *testing.T) {
	r := newEmptyResponse().KeepSessionAlive()

	assert.Equal(t, "1.0", r.Version, "should have correct version")
	assert.Equal(t, false, r.Response.ShouldEndSession, "should have session set to not be kept alive")
}
