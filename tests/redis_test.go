package tests

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"testing"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     testConfig.Database.Redis,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		t.Fatal(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err != redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
