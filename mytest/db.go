package main

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {

	db, err := leveldb.OpenFile("/tmp/goleveldb", nil)
	if err != nil {
		fmt.Printf("openfile err: %v", err)
		return
	}
	defer db.Close()

	//	samplePut(db)
	//gets(db)
	//puts(db)
	for i := 0; i < 100; i++ {
		fmt.Println("batch i= ", i)
		batches(db, time.Now().Unix())
		time.Sleep(time.Second)
	}

	/*wg := sync.WaitGroup{}
	for kk := 0; kk < 10; kk++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				fmt.Printf("i: %v , %v \n", i, time.Now())
				if err := batches(db); err != nil {
					fmt.Printf("put err: %v", err)
				}
			}
			wg.Done()
		}()
	}

	fmt.Printf("before wait")
	wg.Wait()
	fmt.Printf("after wait")*/
	/*if err = db.CompactRange(util.Range{}); err != nil {
		fmt.Printf("CompactRange err: %v", err)

	}*/
}

func batches(db *leveldb.DB, t int64) error {
	batch := new(leveldb.Batch)

	for i := t; i < t+50; i++ {
		batch.Put([]byte(fmt.Sprintf("k3%d", i)), GenRandomBytes(10240))
	}
	err := db.Write(batch, nil)
	if err != nil {
		fmt.Printf("put err: %v", err)
		return err
	}
	return nil
}

func GenRandomBytes(size int) (blk []byte) {
	blk = make([]byte, size)
	rand.Read(blk)
	return
}

func puts(db *leveldb.DB) error {

	for i := 0; i < 10; i++ {
		if i%100 == 0 {
			fmt.Println("i = ", i, time.Now())
		}
		err := db.Put([]byte(fmt.Sprintf("m%v", i)), []byte(fmt.Sprintf("go%v", i)), nil)
		if err != nil {
			fmt.Printf("put err: %v", err)
			return err
		}

	}
	err := db.Put([]byte("k2"), []byte("v2222222"), nil)
	if err != nil {
		fmt.Printf("put err: %v", err)
		return err
	}
	err = db.Put([]byte("k3"), []byte("v3333333"), nil)
	if err != nil {
		fmt.Printf("put err: %v", err)
		return err
	}

	if v, err := db.Get([]byte("abccc"), nil); err != nil {
		fmt.Printf("put err: %v", err)
		return err
	} else {
		fmt.Printf("gut value: %v", v)
	}
	return nil
}

func gets(db *leveldb.DB) error {
	/*
		for i := 0; i < 300; i++ {
			err := db.Put([]byte("k1"), []byte(fmt.Sprintf("v%d", i)), nil)
			if err != nil {
				fmt.Printf("put err: %v", err)
				return err
			}
		}*/

	if v, err := db.Get([]byte("k1-423423"), nil); err != nil {
		fmt.Printf("get err: %v", err)
		return err
	} else {
		fmt.Printf("gut value: %v", v)
	}

	return nil
}

func samplePut(db *leveldb.DB) error {

	err := db.Put([]byte("a1"), []byte("v%d"), nil)
	if err != nil {
		fmt.Printf("put err: %v", err)
		return err
	}
	err = db.Put([]byte("k1"), []byte("v%d"), nil)
	if err != nil {
		fmt.Printf("put err: %v", err)
		return err
	}
	err = db.Put([]byte("k2"), []byte("v%d"), nil)
	if err != nil {
		fmt.Printf("put err: %v", err)
		return err
	}

	if v, err := db.Get([]byte("k1-1234214"), nil); err != nil {
		fmt.Printf("get err: %v", err)
		return err
	} else {
		fmt.Printf("gut value: %v", v)
	}

	return nil
}
