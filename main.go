package main

import (
	"fmt"
	_ "go-snowflake-service/cache"
	_ "go-snowflake-service/conf"
	"go-snowflake-service/snowflake"
	"log"
)

/**
func main() {
	listen, err := net.Listen("tcp", conf.ConfigObject.Addr)

	if err != nil {
		log.Println(err)
	}

	server := grpc.NewServer()
	pb.RegisterIdGeneratorServiceServer(server, &services.IdGeneratorService{})

	reflection.Register(server)

	log.Printf("[Snowflake] Server started on %s", conf.ConfigObject.Addr)

	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
**/

func main() {
	ch := make(chan snowflake.ID)
	go func() {
		count := 2000
		// 并发 count 个 goroutine 进行 snowflake ID 生成
		for i := 0; i < count; i++ {
			go func() {
				id, _ := workers.GenerateId(1)
				log.Println("1", id)
				ch <- id
			}()
		}
	}()

	m := make(map[snowflake.ID]int)
	for i := 0; i < 2000; i++ {
		id := <-ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		if ok {
			fmt.Printf("ID is not unique!\n")
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i
	}
	// 成功生成 snowflake ID
	fmt.Println("All  snowflake ID generate successed!\n")
}
