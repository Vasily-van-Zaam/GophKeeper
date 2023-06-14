package repository

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
)

type Repository interface {
	Local() localStore
	Remote() remoteStore
}

type repository struct {
	config config.Config
	remote remoteStore
	local  localStore
}

// Local implements Repository.
func (r *repository) Local() localStore {
	return r.local
}

// Remote implements Repository.
func (r *repository) Remote() remoteStore {
	return r.remote
}

func New(conf config.Config, local localStore) Repository {
	return &repository{
		remote: NewRemote(conf),
		config: conf,
		local:  local,
	}
}
