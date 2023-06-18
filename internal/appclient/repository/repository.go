package repository

import (
	"github.com/Vasily-van-Zaam/GophKeeper.git/internal/config"
)

type Repository interface {
	Local() Local
	Store() localStore
	Remote() remoteStore
}

type repository struct {
	config config.Config
	remote remoteStore
	local  Local
	strore localStore
}

// LocalStore implements Repository.
func (r *repository) Store() localStore {
	return r.strore
}

// Local implements Repository.
func (r *repository) Local() Local {
	return r.local
}

// Remote implements Repository.
func (r *repository) Remote() remoteStore {
	return r.remote
}

func New(conf config.Config, store localStore) Repository {
	return &repository{
		remote: NewRemote(conf),
		config: conf,
		local:  NewLocal(conf, store),
		strore: store,
	}
}
