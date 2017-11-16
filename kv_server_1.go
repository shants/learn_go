package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"golang.org/x/tools/go/ssa"
	"strconv"
)

//
// Common between client and server
//

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

func server() {
	kv := new(KV)

	kv.keyvalue = map[string]string{}
	rpcs := rpc.NewServer()
	rpcs.Register(kv)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
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
	}()
}

//
// Client
//

func Dial() *rpc.Client {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func Get(key string) string {
	client := Dial()
	args := &GetArgs{"subject"}
	reply := GetReply{"", ""}
	err := client.Call("KV.Get", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()
	return reply.Value
}

func Put(key string, val string) {
	client := Dial()
	args := &PutArgs{"subject", "6.824"}
	reply := PutReply{""}
	err := client.Call("KV.Put", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()
}


func (kv *KV) List(args *ListArgs, reply *ListReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	lr := ListReply{}
	count :=0
	reply.Err = "OK"
	for _, v := range kv.keyvalue {
		if c, e := strconv.Atoi(v); e == nil || c > count{
			break
		}else{
			lr.Value = append(lr.Value, v)
			count++
		}
	}
	return nil
}

//
// main
//

func main() {
	server()

}
