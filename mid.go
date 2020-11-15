package mid

type Mid interface {
	Parse(pack string) error
	Pack() (string, error)
}

type Header struct {
	Length               int
	Mid                  int
	Revision             int
	NoAckFlag            bool
	StationID            int
	SpindleID            int
	SequenceNumber       *int
	NumberOfMessageParts *int
	MessagePartNumber    *int
}

type DataField struct {
}
