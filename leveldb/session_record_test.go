// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package leveldb

import (
	"bytes"
	"fmt"
	"testing"
)

func decodeEncode(v *sessionRecord) (res bool, err error) {
	b := new(bytes.Buffer)
	err = v.encode(b)
	if err != nil {
		return
	}
	v2 := &sessionRecord{}
	err = v.decode(b)
	if err != nil {
		return
	}
	b2 := new(bytes.Buffer)
	err = v2.encode(b2)
	if err != nil {
		return
	}
	return bytes.Equal(b.Bytes(), b2.Bytes()), nil
}

func TestSessionRecord_EncodeDecode(t *testing.T) {
	big := int64(1) << 50
	v := &sessionRecord{}
	i := int64(0)
	test := func() {
		res, err := decodeEncode(v)
		if err != nil {
			t.Fatalf("error when testing encode/decode sessionRecord: %v", err)
		}
		if !res {
			t.Error("encode/decode test failed at iteration:", i)
		}
	}

	for ; i < 4; i++ {
		test()
		v.addTable(3, big+300+i, big+400+i,
			makeInternalKey(nil, []byte("foo"), uint64(big+500+1), keyTypeVal),
			makeInternalKey(nil, []byte("zoo"), uint64(big+600+1), keyTypeDel))
		v.delTable(4, big+700+i)
		v.addCompPtr(int(i), makeInternalKey(nil, []byte("x"), uint64(big+900+1), keyTypeVal))
	}

	v.setComparer("foo")
	v.setJournalNum(big + 100)
	v.setPrevJournalNum(big + 99)
	v.setNextFileNum(big + 200)
	v.setSeqNum(uint64(big + 1000))
	test()
}

func TestMakeInternalKey(t *testing.T) {

	dst := makeInternalKey(nil, []byte("f"), 1, keyTypeVal)
	dst2 := makeInternalKey(nil, []byte("f"), 256, keyTypeVal)
	dst3 := makeInternalKey(nil, []byte("f"), 257, keyTypeVal)

	fmt.Printf(" \n  dst:")
	for _, b := range dst {
		fmt.Printf(" %02x", b)
	}
	fmt.Printf(" \n dst2:")
	for _, b := range dst2 {
		fmt.Printf(" %02x", b)
	}
	fmt.Printf(" \n dst3:")
	for _, b := range dst3 {
		fmt.Printf(" %02x", b)
	}

	b1 := bytes.Compare(dst, dst2)
	b2 := bytes.Compare(dst, dst3)
	fmt.Printf(" \n b1: %v, b2: %v ", b1, b2)
}
