package main

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/syndtr/goleveldb/leveldb/util"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {

	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	db, err := leveldb.OpenFile("/home/huangping/temp/goleveldb", o)
	if err != nil {
		fmt.Printf("openfile err: %v", err)
		return
	}
	//defer db.Close()

	//	samplePut(db)
	gets(db)
	//puts(db)
	/*for i := 0; i < 50; i++ {
		fmt.Println("batch i= ", i)
		batches(db, time.Now().Unix())
		time.Sleep(time.Second)
	}*/

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
	if err = db.CompactRange(util.Range{}); err != nil {
		fmt.Printf("CompactRange err: %v", err)

	}

	s := &leveldb.DBStats{}
	if err := db.Stats(s); err != nil {
		fmt.Printf("db.Stats error, %v", err)
	} else {
		fmt.Printf("db.Stats:%+v", s)
	}
}

func batches(db *leveldb.DB, t int64) error {
	batch := new(leveldb.Batch)

	for i := t; i < t+20; i++ {
		batch.Put([]byte(fmt.Sprintf("k3%d", i)), GenRandomBytes(1023))
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

	for i := 0; i < 2560; i++ {
		key := make([]byte, 1)
		key[0] = byte(i)
		if i%100 == 0 {
			fmt.Println("i = ", i, time.Now())
		}
		err := db.Put(key, []byte(fmt.Sprintf("go%v", i)), nil)
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

	if v, err := db.Get([]byte("x1"), nil); err != nil {
		fmt.Printf("\n get err: %v", err)
		return err
	} else {
		fmt.Printf("\n got value: %v", v)
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
