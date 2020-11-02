package midparser

import (
	"errors"
	"strconv"

	mid "github.com/rickedb/gopenprotocol"
)

func parseHeader(pack string) (mid.Header, error) {
	header := mid.Header{}
	if len(pack) < 20 {
		return header, errors.New("package does not have minimum length of 20 characters")
	}

	var err error
	header.Length, err = strconv.Atoi(pack[0:4])
	header.Mid, err = strconv.Atoi(pack[4:8])
	if err != nil {
		return header, err
	}

	header.Revision, err = strconv.Atoi(pack[8:11])
	if err != nil {
		header.Revision = 1
	}

	intNoAckFlag, err := strconv.Atoi(pack[11:12])
	header.NoAckFlag = err == nil && intNoAckFlag != 0

	header.StationID, err = strconv.Atoi(pack[12:14])
	if err != nil {
		header.StationID = 1
	}

	header.SpindleID, err = strconv.Atoi(pack[14:16])
	if err != nil {
		header.SpindleID = 1
	}

	sequenceNumber, err := strconv.Atoi(pack[16:18])
	if err != nil && sequenceNumber > 0 {
		header.SequenceNumber = &sequenceNumber
	}

	numberOfMessageParts, err := strconv.Atoi(pack[18:19])
	if err != nil && numberOfMessageParts > 0 {
		header.NumberOfMessageParts = &numberOfMessageParts
	}

	messagePartNumber, err := strconv.Atoi(pack[19:20])
	if err != nil && numberOfMessageParts > 0 {
		header.NumberOfMessageParts = &messagePartNumber
	}

	return header, nil
}
