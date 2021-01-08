package main

import (
	"time"

	"go.etcd.io/etcd/clientv3"
)

//Store is the interface that the Client for the store must implement
type Store interface {
}

/*EtcdCtl  is the etcd client that will implement the Store Interface and will be used for
interaction with the etcd store at the backend.*/
type EtcdCtl struct {
	Client *clientv3.Client
}

//NewStore method return store interface to the user.
func NewStore() (Store, error) {
	cli, err := NewClient()
	client := EtcdCtl{Client: cli}
	return client, err
}

//NewClient returns instance of the etcd client in go and allows for direct access to the etcd cluster.
func NewClient() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
}
