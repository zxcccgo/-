package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"

	"github.com/bwmarrin/snowflake"
)

var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(errors.New("snowflake node init failed"))
	}
	var Store Datas

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			id := node.Generate()
			name := strconv.FormatInt(int64(i), 10)
			lock.Lock()
			Store = append(Store, &Data{int64(id), name})
			lock.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(len(Store))
	for _, v := range Store {
		fmt.Println(v)
	}
	fmt.Println("------------------")
	sort.Sort(Store)

	for _, v := range Store {
		fmt.Println(v)
	}

}
//使用sort.Sort必须对 []*data  实现三个功能Len() Swap() Less(<顺序 >倒叙) 
type Data struct {
	Id   int64
	Name string
}
type Datas []*Data

func (d Datas) Len() int {
	return len(d)
}
func (d Datas) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d Datas) Less(i, j int) bool {
	return d[i].Id > d[j].Id
}
