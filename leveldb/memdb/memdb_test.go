// Copyright (c) 2014, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package memdb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/syndtr/goleveldb/leveldb/comparer"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/testutil"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func (p *DB) TestFindLT(key []byte) (rkey, value []byte, err error) {
	p.mu.RLock()
	if node := p.findLT(key); node != 0 {
		n := p.nodeData[node]
		m := n + p.nodeData[node+nKey]
		rkey = p.kvData[n:m]
		value = p.kvData[m : m+p.nodeData[node+nVal]]
	} else {
		err = ErrNotFound
	}
	p.mu.RUnlock()
	return
}

func (p *DB) TestFindLast() (rkey, value []byte, err error) {
	p.mu.RLock()
	if node := p.findLast(); node != 0 {
		n := p.nodeData[node]
		m := n + p.nodeData[node+nKey]
		rkey = p.kvData[n:m]
		value = p.kvData[m : m+p.nodeData[node+nVal]]
	} else {
		err = ErrNotFound
	}
	p.mu.RUnlock()
	return
}

func (p *DB) TestPut(key []byte, value []byte) error {
	p.Put(key, value)
	return nil
}

func (p *DB) TestDelete(key []byte) error {
	p.Delete(key)
	return nil
}

func (p *DB) TestFind(key []byte) (rkey, rvalue []byte, err error) {
	return p.Find(key)
}

func (p *DB) TestGet(key []byte) (value []byte, err error) {
	return p.Get(key)
}

func (p *DB) TestNewIterator(slice *util.Range) iterator.Iterator {
	return p.NewIterator(slice)
}

var _ = testutil.Defer(func() {
	Describe("Memdb", func() {
		Describe("write test", func() {
			It("should do write correctly", func() {
				db := New(comparer.DefaultComparer, 0)
				t := testutil.DBTesting{
					DB:      db,
					Deleted: testutil.KeyValue_Generate(nil, 1000, 1, 1, 30, 5, 5).Clone(),
					PostFn: func(t *testutil.DBTesting) {
						Expect(db.Len()).Should(Equal(t.Present.Len()))
						Expect(db.Size()).Should(Equal(t.Present.Size()))
						switch t.Act {
						case testutil.DBPut, testutil.DBOverwrite:
							Expect(db.Contains(t.ActKey)).Should(BeTrue())
						default:
							Expect(db.Contains(t.ActKey)).Should(BeFalse())
						}
					},
				}
				testutil.DoDBTesting(&t)
			})
		})

	})
})
