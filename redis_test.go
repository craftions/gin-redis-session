package sessions_test

import (
	"testing"

	sessions "github.com/craftions/gin-redis-session"
	"github.com/craftions/gin-redis-session/tester"
)

const redisTestServer = "localhost:6379"

var newRedisStore = func(_ *testing.T) sessions.Store {
	store, err := sessions.NewStore(10, "tcp", redisTestServer, "", 4096, []byte("secret"))
	if err != nil {
		panic(err)
	}
	return store
}

func TestRedis_SessionGetSet(t *testing.T) {
	tester.GetSet(t, newRedisStore)
}

func TestRedis_SessionDeleteKey(t *testing.T) {
	tester.DeleteKey(t, newRedisStore)
}

func TestRedis_SessionFlashes(t *testing.T) {
	tester.Flashes(t, newRedisStore)
}

func TestRedis_SessionClear(t *testing.T) {
	tester.Clear(t, newRedisStore)
}

func TestRedis_SessionOptions(t *testing.T) {
	tester.Options(t, newRedisStore)
}

func TestGetRedisStore(t *testing.T) {
	t.Run("unmatched type", func(t *testing.T) {
		type store struct{ sessions.Store }
		err, rediStore := sessions.GetRedisStore(store{})
		if err == nil || rediStore != nil {
			t.Fail()
		}
	})
}
