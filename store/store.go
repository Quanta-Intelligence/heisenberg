package store

import (
	"heisenberg/model"
	"heisenberg/state"
	"heisenberg/worker"
	"log/slog"
)

type Store struct {
	manager   worker.Manager
	compacter compacter
	disk      state.Disk
	logger    *slog.Logger
}

func New() Store {
	return Store{}
}

func (s *Store) Receive(msg any) {
	switch msg := msg.(type) {
	case model.PutRequest:
		s.logger.Debug(string(msg.Key))
	}
}

func process() {

}
