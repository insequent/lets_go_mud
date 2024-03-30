package telnet

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

// Loosely modeled off github.com/reiver/go-telnet

type Client struct {
	Addr string
	Conn *Conn
}

func NewClient(host string, port int) (*Client, error) {
	return &Client{
		// TODO: Add URL parsing?
		Addr: fmt.Sprintf("%s:%d", host, port),
	}, nil
}

func (c *Client) Dial() error {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		return err
	}

	c.Conn = NewConn(conn)
	return nil
}

// ReadLoop is the primary read loop for the connection
func (c *Client) ReadLoop(wg *sync.WaitGroup) {
	wg.Add(1)

	go func(reader io.Reader, writer io.Writer) {
		done := false

		for {
			p := make([]byte, 1)
			n, err := reader.Read(p)

			switch {
			case err != nil && errors.Is(err, io.EOF):
				log.Printf("Connection to %s closed. Shutting down...", c.Addr)
				done = true
			case err != nil:
				log.Panicf("Error from reading from connection: %v", err)
			case n <= 0:
				continue
			default:
				writer.Write(p)
			}

			if done {
				break
			}
		}

		wg.Done()
	}(c.Conn.Reader, os.Stdout)
}

func (c *Client) StartAndListen() {
	wg := &sync.WaitGroup{}

	c.ReadLoop(wg)
	c.WriteLoop(wg)

	wg.Wait()
}

// WriteLoop is the primary writing loop for the connection
func (c *Client) WriteLoop(wg *sync.WaitGroup) {
	wg.Add(1)

	go func(reader io.Reader, writer io.Writer) {
		scanner := bufio.NewScanner(reader)
		var buffer bytes.Buffer

		for scanner.Scan() {
			buffer.Write(scanner.Bytes())
			buffer.Write(CLRF)

			b := buffer.Bytes()
			n, err := writer.Write(b)
			if err != nil {
				log.Panicf("Error writing to connection: %v", err)
			}

			if n != len(b) {
				log.Panicf("Failed to send full message (only send %d chars): '%s'", len(b), b)
			}

			buffer.Reset()
			time.Sleep(3 * time.Microsecond)
		}
	}(os.Stdin, c.Conn.Writer)

	wg.Done()
}
