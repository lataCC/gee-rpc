package main

import (
	"context"
	"geerpc"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addrCh chan string) {
	//var foo Foo
	//if err := geerpc.Register(&foo); err != nil {
	//	log.Fatal("register error:", err)
	//}
	//// pick a free port
	//l, err := net.Listen("tcp", ":0")
	//if err != nil {
	//	log.Fatal("network error:", err)
	//}
	//log.Println("start rpc server on", l.Addr())
	//addr <- l.Addr().String()
	//geerpc.Accept(l)

	var foo Foo
	l, _ := net.Listen("tcp", ":9999")
	_ = geerpc.Register(&foo)
	geerpc.HandleHTTP()
	addrCh <- l.Addr().String()
	_ = http.Serve(l, nil)
}

func call(addrCh chan string) {
	client, _ := geerpc.DialHTTP("tcp", <-addrCh)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			var reply int
			if err := client.Call(context.Background(), "Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}

func main() {
	//log.SetFlags(0)
	//addr := make(chan string)
	//go startServer(addr)
	//client, _ := geerpc.Dial("tcp", <-addr)
	//defer func() { _ = client.Close() }()
	//
	//time.Sleep(time.Second)
	//
	//var wg sync.WaitGroup
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		args := &Args{Num1: i, Num2: i * i}
	//		var reply int
	//		if err := client.Call(context.Background(),"Foo.Sum", args, &reply); err != nil {
	//			log.Fatal("call Foo.Sum error:", err)
	//		}
	//		log.Printf("%d + %d = %d:", args.Num1, args.Num2, reply)
	//	}(i)
	//}
	//wg.Wait()
	log.SetFlags(0)
	ch := make(chan string)
	go call(ch)
	startServer(ch)
}
