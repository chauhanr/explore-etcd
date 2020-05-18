package main

import (
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("Error connecting to etcd. %s\n", err)
		return
	}
	defer cli.Close()
	fmt.Printf("Connected to etcd instance.\n")

}
