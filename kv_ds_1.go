package main

import "sync"

const (
	OK       = "OK"
	ErrNoKey = "ErrNoKey"
)

type Err string

type PutArgs struct {
	Key   string
	Value string
}

type PutReply struct {
	Err Err
}

type GetArgs struct {
	Key string
}

type GetReply struct {
	Err   Err
	Value string
}

type ListArgs struct {
	Key string
}

type ListReply struct {
	Err Err
	Value []string
}
//
// Server
//

type KV struct {
	mu       sync.Mutex
	keyvalue map[string]string
}


