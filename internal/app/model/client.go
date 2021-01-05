package model

import (
    "bufio"
    "net"
)

type Client struct {
    Username   string
    Connection net.Conn
    Reader     *bufio.Reader
    Writer     *bufio.Writer
}
