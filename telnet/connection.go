package telnet

import (
	"io"
	"net"
)

const (
	SE = 240
	SB = 250

	WILL = 251
	WONT = 252
	DO   = 253
	DONT = 254

	IAC = 255
)

var CLRF = []byte{'\r', '\n'}

type Conn struct {
	Reader io.Reader // buffered
	Writer io.Writer

	conn net.Conn
}

func NewConn(conn net.Conn) *Conn {
	return &Conn{
		//Reader: bufio.NewReader(conn),
		Reader: conn,
		Writer: conn,

		conn: conn,
	}
}
