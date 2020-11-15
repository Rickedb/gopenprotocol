package parser

import (
	"errors"
	"fmt"
	"strconv"

	mid "github.com/rickedb/gopenprotocol"
)

func ParseHeader(pack string) (mid.Header, error) {
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

func PackHeader(header mid.Header) string {
	pack := fmt.Sprintf("%04d", header.Length)
	pack += fmt.Sprintf("%04d", header.Mid)
	pack += fmt.Sprintf("%03d", header.Revision)
	if header.NoAckFlag {
		pack += "1"
	} else {
		pack += "0"
	}
	pack += fmt.Sprintf("%02d", header.StationID)
	pack += fmt.Sprintf("%02d", header.SpindleID)
	if header.SequenceNumber != nil {
		pack += fmt.Sprintf("%2d", *header.SequenceNumber)
	} else {
		pack += fmt.Sprintf("%2s", " ")
	}

	if header.NumberOfMessageParts != nil {
		pack += fmt.Sprintf("%2d", *header.NumberOfMessageParts)
	} else {
		pack += fmt.Sprintf("%2s", " ")
	}

	if header.MessagePartNumber != nil {
		pack += fmt.Sprintf("%2d", *header.MessagePartNumber)
	} else {
		pack += fmt.Sprintf("%2s", " ")
	}

	return pack
}
