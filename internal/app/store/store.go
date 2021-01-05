package store

import (
    "github.com/iluxaorlov/termchat/internal/app/model"
    "sync"
)

type Store struct {
    clients map[*model.Client]*model.Client
    mutex   *sync.Mutex
}

func New() *Store {
    return &Store{
        clients: make(map[*model.Client]*model.Client),
        mutex:   &sync.Mutex{},
    }
}

func (s *Store) Add(c *model.Client) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.clients[c] = c
}

func (s *Store) Del(c *model.Client) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    delete(s.clients, c)
}

func (s *Store) All() map[*model.Client]*model.Client {
    return s.clients
}
