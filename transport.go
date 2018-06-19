package sml

import (
	"bufio"
	"bytes"

	"github.com/pkg/errors"
)

const (
	MAXFILESIZE = 512
)

var (
	EscSeq = []byte{0x1b, 0x1b, 0x1b, 0x1b}
	EndSeq = []byte{0x1b, 0x1b, 0x1b, 0x1b, 0x1a}
)

func ReadChunk(r *bufio.Reader, buf []byte) error {
	bytes, err := r.Read(buf)
	if err != nil {
		return err
	}

	if bytes < len(buf) {
		// fmt.Printf("ReadChunk %d -> %d\n", len(buf), bytes)
		return errors.New("premature eof")
	}

	// success - no error
	return nil
}

func TransportRead(r *bufio.Reader) ([]byte, error) {
	buf := make([]byte, MAXFILESIZE)

	var len int
	var err error

	// find escape sequence/begin 1B 1B 1B 1B 01 01 01 01
	for len < 8 {
		if buf[len], err = r.ReadByte(); err != nil {
			return nil, err
		}

		if (buf[len] == 0x1b && len < 4) || (buf[len] == 0x01 && len >= 4) {
			len++
		} else {
			len = 0
		}
	}

	// found start sequence
	for len+8 < MAXFILESIZE {
		if err = ReadChunk(r, buf[len:len+4]); err != nil {
			return nil, err
		}

		// find escape sequence
		if bytes.Equal(buf[len:len+4], EscSeq) {
			len += 4

			// read end sequence
			if err = ReadChunk(r, buf[len:len+4]); err != nil {
				return nil, err
			}

			if buf[len] == 0x1a {
				// found end sequence
				len += 4
				return buf[:len], nil
			}

			// don't read other escaped sequences yet
			return nil, errors.New("unrecognized sequence")
		}

		// continue reading
		len += 4
	}

	return nil, errors.New("read buffer exceeded")
}
