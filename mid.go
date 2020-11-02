package mid

type Mid interface {
	parse(pack string)
	pack() string
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
