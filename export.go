package sml

import (
	"bufio"
)

/*
 * Exports
 */

func TransportRead(r *bufio.Reader) ([]byte, error) {
	return sml_transport_read(r)
}

func FileParse(buf []byte) ([]sml_message, error) {
	return sml_file_parse(buf)
}

func MessageParse(buf *sml_buffer) (sml_message, error) {
	return sml_message_parse(buf)
}
