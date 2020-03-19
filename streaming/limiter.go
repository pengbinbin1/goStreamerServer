package main

import (
	"log"
)

type ConnLimter struct {
	Capacity int
	Bucket   chan int
}

func NewConnLimeter(n int) *ConnLimter {
	cl := &ConnLimter{}
	cl.Capacity = n
	cl.Bucket = make(chan int, n)
	return cl
}

func (cl *ConnLimter) ReleaseConn() {
	log.Println("releas a connection")
	<-cl.Bucket
}

func (cl *ConnLimter) GetConn() bool {
	if len(cl.Bucket) >= cl.Capacity {
		log.Println("bucket is full")
		return false
	}
	cl.Bucket <- 1
	return true
}
