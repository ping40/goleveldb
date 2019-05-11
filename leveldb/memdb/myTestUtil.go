package memdb

import "fmt"

func PrintMemDB(db *DB) {
	fmt.Printf("====\n prevNode: %v \n", db.prevNode)
	fmt.Printf("nodeData: len: %v,   [index]  ", len(db.nodeData))
	for k, _ := range db.nodeData {
		fmt.Printf("%03x ", k)
	}
	fmt.Printf("\nnodeData: len: %v,  [content]%v\n", len(db.nodeData), showData1(db.nodeData))

	fmt.Printf("vData: len: %v,   [index]  ", len(db.kvData))
	for k, _ := range db.kvData {
		fmt.Printf("%03x ", k)
	}
	fmt.Printf("\nvData: len: %v, [content] %v\n", len(db.kvData), showData21(db.kvData))

}

func showData1(b []int) string {
	s := ""
	for _, v := range b {
		s = fmt.Sprintf("%s %03x", s, v)
	}
	return s
}

func showData21(b []byte) string {
	s := ""
	for _, v := range b {
		s = fmt.Sprintf("%s %03x", s, v)
	}
	return s
}
