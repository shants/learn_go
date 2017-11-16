package main

import (
	"fmt"
	"log"
	"sync"
	"net/rpc"
)
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

func list(no int) []string {
	fmt.Println("here1")
	client := Dial()

	args := &ListArgs{Key:string(no) }
	reply := ListReply{Value:[]string{}}
	fmt.Println("call rpc kv.list")
	err := client.Call("KV.List", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println(" kv.list return ", reply.Value)
	client.Close()
	return reply.Value
}

func main() {

	Put("subject", "6.824")
	Put("subject1", "6.001")
	Put("subject2", "6.002")
	Put("subject3", "6.003")
	fmt.Printf("Put subjet done\n")

	fmt.Printf("Get %s\n", Get("subject2"))
	fmt.Printf("Get %s\n", Get("subject3"))

	list(3)
	//fmt.Printf("list done\n %v %v", v[0])
	/*
	for i,n := range v {
		fmt.Printf("\n %v  %v", i,n)
	}
	*/

		fmt.Printf("Get %s\n", Get("subject"))


}
