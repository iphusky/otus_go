package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

var (
	ErrNotEnoughArgs        = errors.New("host and port must be set")
	ErrAddressAndPortNotSet = errors.New("connection error: address not set")
)

type TelnetClient struct {
	address string
	conn    net.Conn
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) *TelnetClient {
	return &TelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *TelnetClient) Connect() error {
	if len(c.address) == 0 {
		return ErrAddressAndPortNotSet
	}

	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	c.conn = conn

	return nil
}

func (c *TelnetClient) Close() error {
	return c.conn.Close()
}

func (c *TelnetClient) Send() error {
	_, err := io.Copy(c.conn, c.in)
	return err
}

func (c *TelnetClient) Receive() error {
	_, err := io.Copy(c.out, c.conn)
	return err
}
