# Redis Sessions

Gin middleware for Redis  session management with multi-backend support.
This implementation was forked from [gin-contrib/sessions](https://github.com/gin-contrib/sessions) and modified for specific engineering problems. We don't recommend using it in your projects because We don’t offer maintenance and breaking changes can be introduced in the future.

## Usage

### Start using it

Download and install it:

```bash
go get github.com/craftions/gin-redis-session
```

Import it in your code:

```go
import "github.com/craftions/gin-redis-session"
```

## Basic Examples

### Redis

```go
package main

import (
  sessions "github.com/craftions/gin-redis-session"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", 4096, []byte("secret"))
  r.Use(sessions.Sessions("mysession", store))

  r.GET("/incr", func(c *gin.Context) {
    session := sessions.Default(c)
    var count int
    v := session.Get("count")
    if v == nil {
      count = 0
    } else {
      count = v.(int)
      count++
    }
    session.Set("count", count)
    session.Save()
    c.JSON(200, gin.H{"count": count})
  })
  r.Run(":8000")
}
```
