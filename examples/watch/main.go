package main

import (
	"context"
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
		fmt.Errorf("Error connecting to etcd. %s\n", err)
		return
	}
	defer cli.Close()
	rch := cli.Watch(context.Background(), "/api/v1/Pod", clientv3.WithPrefix())
	for wres := range rch {
		for _, ev := range wres.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

}
