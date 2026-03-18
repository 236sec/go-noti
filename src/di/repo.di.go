package di

import (
	"sync"

	"goboilerplate.com/src/repo"
)

var getUserRepo = sync.OnceValue(func() *repo.UserRepo {
	return repo.NewUserRepo(GetDB(),GetRedis())
})