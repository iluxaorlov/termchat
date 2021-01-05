package main

import (
    "github.com/iluxaorlov/termchat/internal/app/server"
    "github.com/iluxaorlov/termchat/internal/app/store"
    "github.com/joho/godotenv"
    "log"
    "net"
    "os"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal(err)
    }

    a := os.Getenv("SERVER_ADDRESS")

    l, err := net.Listen("tcp", a)
    if err != nil {
        log.Fatal(err)
    }

    defer l.Close()

    s := server.New(store.New())

    for {
        c, err := l.Accept()
        if err != nil {
            continue
        }

        go s.Handle(c)
    }
}
