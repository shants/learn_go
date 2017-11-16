package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"sync"
)

//
// Common between client and server
//



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


//////////////////////////////////////////////////////////

func (kv *KV) Get(args *GetArgs, reply *GetReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	reply.Err = "OK"
	val, ok := kv.keyvalue[args.Key]
	if ok {
		reply.Err = OK
		reply.Value = val
	} else {
		reply.Err = ErrNoKey
		reply.Value = ""
	}
	return nil
}

func (kv *KV) Put(args *PutArgs, reply *PutReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.keyvalue[args.Key] = args.Value
	reply.Err = OK
	return nil
}


func (kv *KV) List(args *ListArgs, reply *ListReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	lr := ListReply{}
	count :=0
	reply.Err = "OK"
	for _, v := range kv.keyvalue {
		if c, e := strconv.Atoi(v); e == nil || count > c{
			break
		}else{
			fmt.Println("\n appending %v", v)
			lr.Value = append(lr.Value, v)
			count++
		}
	}
	fmt.Println("\n returning %v", lr.Value)
	return nil
}

func server() {
	kv := new(KV)

	kv.keyvalue = map[string]string{}
	rpcs := rpc.NewServer()
	rpcs.Register(kv)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
//	go func() {
		for {
			conn, err := l.Accept()
			if err == nil {
				go rpcs.ServeConn(conn)
			} else {
				break
			}
		}
		l.Close()
		fmt.Printf("Server done\n")
//	}()
}


//
// main
//

func main() {
	server()

}
