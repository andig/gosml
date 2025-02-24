package sml

import (
	"bufio"
	"bytes"

	"errors"
)

var (
	// EscSeq is the SML escape sequence (mark begin and end of a message)
	EscSeq = []byte{0x1b, 0x1b, 0x1b, 0x1b}

	// MaxFileSize is the maximum number of bytes in a file (this is the chunk
	// of bytes which will be read from the meter between the escape sequences)
	// Some meter can be set to deliver an "extended dataset", which will be
	// larger than the default 512 bytes. If you set the extended dataset on
	// your meter, you should increase this value to 1024.
	MaxFileSize = 512
)

func ReadChunk(r *bufio.Reader, buf []byte) error {
	bytesRead, err := r.Read(buf)
	if err != nil {
		return err
	}

	if bytesRead < len(buf) {
		return errors.New("premature EOF")
	}

	return nil
}

func TransportRead(r *bufio.Reader) ([]byte, error) {
	buf := make([]byte, MaxFileSize)

	var l int
	var err error

	// find escape sequence/begin 1B 1B 1B 1B 01 01 01 01
	for l < 8 {
		if buf[l], err = r.ReadByte(); err != nil {
			return nil, err
		}

		if (buf[l] == 0x1b && l < 4) || (buf[l] == 0x01 && l >= 4) {
			l++
		} else {
			l = 0
		}
	}

	// found start sequence
	for l+8 < MaxFileSize {
		if err = ReadChunk(r, buf[l:l+4]); err != nil {
			return nil, err
		}

		// find escape sequence
		if bytes.Equal(buf[l:l+4], EscSeq) {
			l += 4

			// read end sequence
			if err = ReadChunk(r, buf[l:l+4]); err != nil {
				return nil, err
			}

			if buf[l] == 0x1a {
				// found end sequence
				l += 4
				return buf[:l], nil
			}

			// don't read other escaped sequences yet
			return nil, errors.New("unrecognized sequence")
		}

		// continue reading
		l += 4
	}

	return nil, errors.New("read buffer exceeded")
}
