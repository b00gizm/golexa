# golexa

A little Go library to easily handle Alexa custom skill requests, conforming to the official [Alexa Skill Kit JSON reference](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/alexa-skills-kit-interface-reference#outputspeech-object). golexa handles all the heavy lifting, so you can concentrate on building great apps within minutes.

## Install

```bash
go get github.com/b00giZm/golexa
```

or, if you prefer [Glide](https://github.com/Masterminds/glide)

```bash
glide get github.com/b00giZm/golexa
```

## Examples

### Quickstart

Your first Alexa app in less than 10 lines of code.

```go
import "github.com/b00giZm/golexa"

app := golexa.Default()
app.OnLaunch(func(a *golexa.Alexa, req *golexa.Request, session *golexa.Session) *golexa.Response {
        return a.Response().AddPlainTextSpeech("Welcome to my awesome app")
})

response := app.Process(json)
```

You can attach several handlers to deal with the different Alexa request types:

#### LaunchRequest handler
```go
app.OnLaunch(func(a *Alexa, req *Request, session *Session) *Response) { ... })
```

#### IntentRequest handler
```go
app.OnIntent(func(a *Alexa, intent *Intent, req *Request, session *Session) *Response) { ... })
```

#### SessionEndedRequest handler
```go
app.OnSessionEnded(func(a *Alexa, req *Request, session *Session) *Response) { ... })
```

Every handler has a `*golexa.Alexa` pointer as first parameter, which gives you access to several convenience methods for quickly building response objects.

## Your Own Handler Types

When building bigger apps, using the built-in callback handlers can be a bit cumbersome, especially if you have lots of different intents. For this special case, golexa provides several interfaces to implement for your own handlers:

```go
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
```

With these interfaces and the `Init()` method of `golexa.Alexa`, you can have one handler for each intent, which is a much cleaner solution than having a big handler with lots of `if else` statements:

```go
myHandlers := []golexa.Handler{
        myFooIntentHandler,
        myBarIntentHandler,
        ...
}

app := golexa.Init(myHandlers)
response := app.Process(json)
```

Those handlers are processed in the order in which they were provided. The processing goes on until one of the handlers return a valid `*golexa.Reponse`. If a handler can't handle a specific request, it can just return `nil` and the next consecutive handler is called.

(by the way: You can of course provide your own handlers for `LaunchRequest` and `SessionEndedRequest` inside the `Init()` slice, but remember that there can be only one `LaunchRequest` handler and just one `SessionEndedRequest` handler, which means that if you provide several of them, all but the very last one will be discarded.)

## Apex Usage

The main motivation for this library was to have something which plays nicely with the awesome [Apex](http://apex.run/) toolchain and their Golang bridge for AWS Lambda:

```go
package main

import (
        "github.com/apex/go-apex"
        "github.com/b00giZm/golexa"
)

func main() {
        apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
                app := golexa.Default()
                app.OnLaunch(func(a *golexa.Alexa, req *golexa.Request, session *golexa.Session) *golexa.Response {
                        return a.Response().AddPlainTextSpeech("Welcome to my awesome app")
                })
        
                return app.Process(event)
        })
}
```

Boom! Done 🎉

## Development

Prerequisites:

* Go >= 1.6.2
* Glide >= 0.10.2

Fork and clone this repository

```bash
cd $GOATH/src/github.com/b00giZm
git clone https://github.com/b00giZm/golexa.git
```

Install dependencies via [Glide](https://github.com/Masterminds/glide)

```bash
glide install
```

Run tests

```bash
go test
```

## Maintainer

Pascal Cremer

* Email: <hello@codenugget.co>
* Twitter: [@b00gizm](https://twitter.com/b00gizm)
* Web: [http://codenugget.co](http://codenugget.co)

## License

> The MIT License (MIT)
>
> Copyright (c) 2016 Pascal Cremer
>
>Permission is hereby granted, free of charge, to any person obtaining a copy
>of this software and associated documentation files (the "Software"), to deal
>in the Software without restriction, including without limitation the rights
>to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
>copies of the Software, and to permit persons to whom the Software is
>furnished to do so, subject to the following conditions:
>
>The above copyright notice and this permission notice shall be included in all
>copies or substantial portions of the Software.
>
>THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
>IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
>FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
>AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
>LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
>OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
>SOFTWARE.
