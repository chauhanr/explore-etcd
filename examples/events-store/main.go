package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"go.etcd.io/etcd/clientv3"
)

//key_prefix will be put to use as prefix to look for
var key_prefix = "pod"

func main() {
	cli, err := NewClient()
	if err != nil {
		panic(err)
	}
	ulidSource := NewMonotonicULIDsource(baseEntropy())

	// generate bunch of events and insert into the etcd store.
	ctx := context.Background()
	for i := 0; i <= 10; i++ {
		now := time.Now()
		id, _ := ulidSource.New(now)
		event := FakeEvent{ID: id.String(), Data: i, TimeStamp: now}
		e, err := json.Marshal(event)
		if err != nil {
			fmt.Printf("Error marshalling the fake event %s\n", err)
			return
		}
		es := string(e)
		key := key_prefix + id.String()
		_, err = cli.KV.Put(ctx, key, es)
		if err != nil {
			fmt.Printf("error inserting key %s\n", id.String())
		}
	}
	// retirive it and look at the order
	res, err := cli.KV.Get(ctx, key_prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		fmt.Printf("Failed to retrive keys %s", key_prefix)
		return
	}
	for _, ev := range res.Kvs {
		fmt.Printf("%s: %s\n", ev.Key, ev.Value)
	}

	// clean up the keys

}

func baseEntropy() *rand.Rand {
	t := time.Unix(1000000, 0)
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return entropy
}
