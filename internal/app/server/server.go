package server

import (
    "bufio"
    "fmt"
    "github.com/iluxaorlov/termchat/internal/app/model"
    "github.com/iluxaorlov/termchat/internal/app/store"
    "net"
    "strings"
)

type Server struct {
    store *store.Store
}

func New(store *store.Store) *Server {
    return &Server{
        store: store,
    }
}

func (s *Server) Handle(c net.Conn) {
    r := bufio.NewReader(c)
    w := bufio.NewWriter(c)

    l, err := r.ReadString('\n')
    if err != nil {
        return
    }

    u := strings.Trim(l, "\n")

    client := &model.Client{
        Username:   u,
        Connection: c,
        Reader:     r,
        Writer:     w,
    }

    s.add(client)

    s.listen(client)
}

func (s *Server) listen(c *model.Client) {
    for {
        l, err := c.Reader.ReadString('\n')
        if err != nil {
            break
        }

        m := fmt.Sprintf("\u001b[37m%s:\u001b[0m %s", c.Username, l)

        for client := range s.store.All() {
            if client == c {
                continue
            }

            go s.write(client, m)
        }
    }

    c.Connection.Close()

    s.del(c)
}

func (s *Server) write(c *model.Client, m string) {
    c.Writer.WriteString(m)
    c.Writer.Flush()
}

func (s *Server) add(c *model.Client) {
    s.store.Add(c)
    s.notify(c, "\u001b[37m" + c.Username + ": connected\u001b[0m\n")
}

func (s *Server) del(c *model.Client) {
    s.store.Del(c)
    s.notify(c, "\u001b[37m" + c.Username + ": disconnected\u001b[0m\n")
}

func (s *Server) notify(c *model.Client, m string) {
    for client := range s.store.All() {
        if client == c {
            continue
        }

        go s.write(client, m)
    }
}
