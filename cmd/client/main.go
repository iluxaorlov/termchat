package main

import (
    "bufio"
    "github.com/joho/godotenv"
    "log"
    "net"
    "os"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal(err)
    }

    a := os.Getenv("CLIENT_ADDRESS")

    c, err := net.Dial("tcp", a)
    if err != nil {
        log.Fatal(err)
    }

    go write(c)

    r := bufio.NewReader(c)
    w := bufio.NewWriter(os.Stdout)

    w.WriteString("\u001b[37mEnter your username:\u001b[0m ")
    w.Flush()

    for {
        l, err := r.ReadString('\n')
        if err != nil {
            w.WriteString("\u001b[37mConnection broken\n")
            w.Flush()
            return
        }

        w.WriteString(l)
        w.Flush()
    }
}

func write(c net.Conn) {
    r := bufio.NewReader(os.Stdin)
    w := bufio.NewWriter(c)

    for {
        t, err := r.ReadString('\n')
        if err != nil {
            continue
        }

        w.WriteString(t)
        w.Flush()
    }
}
