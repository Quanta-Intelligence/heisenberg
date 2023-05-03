package store

import (
	"fmt"
	"heisenberg/internal"
	"math/rand"
)

// Interface for handling replication
type replica struct {
	clients map[string]*StoreClient
}

func (r *replica) addReplica(c *StoreClient, id string) error {
	if c == nil {
		return internal.NilClientError()
	}
	r.clients[id] = c
	return nil
}

// Choose random replicas to distribute read requests evenly
func (r *replica) choose() (*StoreClient, error) {
	size := len(r.clients)
	if size == 0 {
		return nil, fmt.Errorf("no replicas")
	}
	idx := rand.Intn(size)
	var replica *StoreClient
	for _, r := range r.clients {
		if idx == 0 {
			replica = r
			break
		}
		idx--
	}
	return replica, nil
}
