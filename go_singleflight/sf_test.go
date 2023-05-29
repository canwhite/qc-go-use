package gosingleflight

import (
	"strconv"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

type user1 struct {
	id       int
	name     string
	password string
	email    string
	token    string
}

func getuserByID(sg *singleflight.Group, id int) user1 {
	// 使用 id 作为 key
	v, _, _ := sg.Do(strconv.Itoa(id), func() (interface{}, error) {
		// 模拟数据库查询耗时
		time.Sleep(time.Millisecond)
		return user1{}, nil
	})
	return v.(user1)
}
func BenchmarkBufferWithPool1(b *testing.B) {
	var wg sync.WaitGroup
	var sg singleflight.Group

	for n := 0; n < b.N; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = getuserByID(&sg, 1024)
		}()
	}

	wg.Wait()
}
