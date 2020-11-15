package keepalive

import (
	mid "github.com/rickedb/gopenprotocol"
	parser "github.com/rickedb/gopenprotocol/parser"
)

type Mid9999 struct {
	Header mid.Header
}

func (m *Mid9999) Parse(pack string) error {
	header, err := parser.ParseHeader(pack)
	m.Header = header
	return err
}

func (m *Mid9999) Pack() (string, error) {
	pack := parser.PackHeader(m.Header)
	return pack, nil
}
