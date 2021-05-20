package main

import (

		"context"
		//"strconv"
		"fmt"
		"github.com/go-redis/redis/v8"
		//"math/rand"
		"sync"


    )
var ctx = context.Background()
var wg sync.WaitGroup

func callRedis() {
	defer wg.Done()
	rdb := redis.NewClient(&redis.Options{
        Addr:     "192.168.82.223:6379",
        Password: "123456", // no password set
        DB:       0,  // use default DB
    })
	//rdb.Decr(ctx, "hongbaonum")

	keys := []string{"hongbaos"}
	// cmd := `
	// local a = redis.call('get', KEYS[1])
	// a = a-2
	// return redis.call('set', KEYS[1],a)
	// `

	cmd2 := `
	if redis.call("llen", KEYS[1]) == 0 then 
		return 0
	end

	local a = redis.call("lpop", KEYS[1])
    	return a
	
	`


	res, err := rdb.Eval(ctx, cmd2, keys).Result()
	if err != nil {
		fmt.Println(err)
	}
	if res != "0" {
		fmt.Println("=====",res)
	}
	


	//fmt.Println(res)

	//fmt.Println(keys)
	// num -=1
	// rdb.Set(ctx, "hongbaonum", num, 0)

}
func main() {
	//rand.Seed(86)
	// 发红包
	rdb := redis.NewClient(&redis.Options{
        Addr:     "192.168.82.223:6379",
        Password: "123456", // no password set
        DB:       0,  // use default DB
    })
	if r, err := rdb.LPush(ctx, "hongbaos", 6,10,8,2,4).Result();err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
	// 10个人抢红包
	for i:=0;i<10;i++ {
		wg.Add(1)
		go callRedis()
		
	}
	wg.Wait()
}