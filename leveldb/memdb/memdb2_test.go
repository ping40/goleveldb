package memdb

import (
	"fmt"
	"testing"

	"github.com/syndtr/goleveldb/leveldb/comparer"
)

func TestRandHeight(t *testing.T) {
	db := New(comparer.DefaultComparer, 0)
	fmt.Println("begin")
	m := make(map[int]int)
	for i := 0; i < 20000000; i++ {
		h := db.randHeight()
		m[h] = m[h] + 1

	}
	for k, v := range m {
		fmt.Printf("%d: %d\n", k, v)
	}
}

func TestPut(t *testing.T) {
	db := New(comparer.DefaultComparer, 0)
	p(db)
	db.Put([]byte("a9"), []byte("v"))
	p(db)
	db.Put([]byte("a8"), []byte("v"))
	p(db)
	db.Put([]byte("a9"), []byte("v12345"))
	p(db)
	db.Delete([]byte("a9"))
	p(db)

	/*
		db.Put([]byte("a2"), []byte("v2"))
		db.Put([]byte("a3"), []byte("v3"))
		db.Put([]byte("a4"), []byte("v4"))
		db.Put([]byte("a5"), []byte("v5"))*/

	v, err := db.Get([]byte("a80"))
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Printf("kaka: %v", v)
	}
}

func p(db *DB) {
	fmt.Printf("====\n prevNode: %v \n", db.prevNode)
	fmt.Printf("nodeData: len: %v,   [index]  ", len(db.nodeData))
	for k, _ := range db.nodeData {
		fmt.Printf("%02x ", k)
	}
	fmt.Printf("\nnodeData: len: %v,  [content]%v\n", len(db.nodeData), showData(db.nodeData))

	fmt.Printf("vData: len: %v,   [index]  ", len(db.kvData))
	for k, _ := range db.kvData {
		fmt.Printf("%02x ", k)
	}
	fmt.Printf("\nData: len: %v,  [content] %v\n", len(db.kvData), showData2(db.kvData))

}

func showData(b []int) string {
	s := ""
	for _, v := range b {
		s = fmt.Sprintf("%s %02x", s, v)
	}
	return s
}

func showData2(b []byte) string {
	s := ""
	for _, v := range b {
		s = fmt.Sprintf("%s %02x", s, v)
	}
	return s
}
