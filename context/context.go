package context

import (
	"github.com/go-redis/redis"
	"github.com/leyle/ginbase/dbandmq"
	"github.com/leyle/testproduct/config"
	"github.com/leyle/userandrole/api"
)

type Context struct {
	Ds *dbandmq.Ds
	R *redis.Client
	Conf *config.Config
	Au *api.UserOption
}
