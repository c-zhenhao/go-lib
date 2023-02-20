package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kelchy/go-lib/redis"
)

func main() {

	uri := "<redis uri here>"

	// cert string is format sensitive, remove indentation
	clientCert := `-----BEGIN CERTIFICATE-----
< your cert here >
-----END CERTIFICATE-----`
	clientKey := `-----BEGIN RSA PRIVATE KEY-----
< your key here >
-----END RSA PRIVATE KEY-----`

	// use redis.New if TLS connection is not required
	redisclient, e := redis.NewSecure(uri, []byte(clientCert), []byte(clientKey))
	if e != nil {
		fmt.Println(e)
		return
	}

	// inputting nil will cause an error, context.TODO() is preferred
	keys, _ := redisclient.Keys(context.TODO(), "*")
	fmt.Println("keys", keys)

	res, _ := redisclient.Set(context.TODO(), "key", "value2", 10*time.Second)
	fmt.Println("result", res)

	resi, _ := redisclient.Del(context.TODO(), "key")
	fmt.Println("result int", resi)

	resb, _ := redisclient.SetNX(context.TODO(), "key", "value2", 10*time.Second)
	fmt.Println("result bool", resb)

	val, _ := redisclient.Get(context.TODO(), "key")
	fmt.Println("result value", val)

	lock, _ := redisclient.Lock(context.TODO(), "locktest", 20*time.Second)
	fmt.Println("result lock", lock)

	unlock, _ := redisclient.Unlock(context.TODO(), "locktest")
	fmt.Println("result unlock", unlock)
}
