package sml

import (
	"bufio"
	"bytes"
	"errors"
)

const (
	SML_MAX_FILE_SIZE = 512
)

var (
	escSeq = []byte{0x1b, 0x1b, 0x1b, 0x1b}
	endSeq = []byte{0x1b, 0x1b, 0x1b, 0x1b, 0x1a}
)

func sml_read_chunk(r *bufio.Reader, buf []byte) error {
	bytes, err := r.Read(buf)
	if err != nil {
		return err
	}

	if bytes < len(buf) {
		// fmt.Printf("sml_read_chunk %d -> %d\n", len(buf), bytes)
		return errors.New("sml: premature eof")
	}

	// success - no error
	return nil
}

func sml_transport_read(r *bufio.Reader) ([]byte, error) {
/*
	unsigned char buf[max_len];
	memset(buf, 0, max_len);
	unsigned int len = 0;

	if (max_len < 8) {
		// prevent buffer overflow
		fprintf(stderr, "libsml: error: sml_transport_read buffer overflow\n");
		return 0;
	}

	while (len < 8) {
		if (sml_read(fd, &readfds, &(buf[len]), 1) == 0) {
			return 0;
		}

		if ((buf[len] == 0x1b && len < 4) || (buf[len] == 0x01 && len >= 4)) {
			len++;
		}
		else {
			len = 0;
		}
	}

	// found start sequence
	while ((len+8) < max_len) {
		if (sml_read(fd, &readfds, &(buf[len]), 4) == 0) {
			return 0;
		}

		if (memcmp(&buf[len], esc_seq, 4) == 0) {
			// found esc sequence
			len += 4;
			if (sml_read(fd, &readfds, &(buf[len]), 4) == 0) {
				return 0;
			}

			if (buf[len] == 0x1a) {
				// found end sequence
				len += 4;
				memcpy(buffer, &(buf[0]), len);
				return len;
			}
			else {
				// don't read other escaped sequences yet
				fprintf(stderr,"libsml: error: unrecognized sequence\n");
				return 0;
			}
		}
		len += 4;
	}

	return 0;
*/
	buf := make([]byte, SML_MAX_FILE_SIZE)

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
	for len+8 < SML_MAX_FILE_SIZE {
		if err = sml_read_chunk(r, buf[len:len+4]); err != nil {
			return nil, err
		}

		// find escape sequence
		if bytes.Equal(buf[len:len+4], escSeq) {
			len += 4

			// read end sequence
			if err = sml_read_chunk(r, buf[len:len+4]); err != nil {
				return nil, err
			}

			if buf[len] == 0x1a {
				// found end sequence
				len += 4
				return buf[:len], nil
			} else {
				// don't read other escaped sequences yet
				return nil, errors.New("sml: unrecognized sequence")
			}
		}

		// continue reading
		len += 4
	}

	return nil, errors.New("sml: read buffer exceeded")
}
