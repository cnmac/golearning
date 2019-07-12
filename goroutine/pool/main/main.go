package main

import (
	"github.com/cnmac/golearning/goroutine/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines  = 25
	pooledResource = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("close connection : db id :", dbConn.ID)
	return nil
}

var idCounter int32

func createConn() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("create new connection ", id)
	return &dbConnection{
		ID: id,
	}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)
	p, err := pool.New(createConn, pooledResource)
	if err != nil {
		log.Println(err)
	}
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQuerys(p, q)
			wg.Done()
		}(query)
	}
	wg.Wait()
	log.Println("shutdown program")
	p.Close()
}

func performQuerys(p *pool.Pool, i int) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", i, conn.(*dbConnection).ID)
}
