package climacell

import "net/url"

type Request interface {
	ToQuery(*RequestDefaults) *url.Values
}
