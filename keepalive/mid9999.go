package keepalive

import (
	mid "github.com/rickedb/gopenprotocol"
)

type Mid9999 struct {
	Header mid.Header
}

func (mid Mid9999) parse(pack string) {
	mid.Header = parseHeader()

	return mid
}

func (mid Mid9999) pack() string {
	return ""
}
