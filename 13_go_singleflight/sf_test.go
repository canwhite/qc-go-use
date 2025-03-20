package gosingleflight

import (
	"strconv"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)
// singleflight 是 Go 语言中的一个并发控制工具，主要用于解决缓存击穿问题。
// 它的核心思想是：对于同一个 key 的并发请求，只执行一次操作，其他并发请求共享这个结果。
// 这样可以避免在缓存失效时，大量请求同时穿透到数据库导致数据库压力过大。
// 
// 主要特点：
// 1. 使用 Do 方法执行回调函数，保证相同 key 的并发调用只会执行一次
// 2. 返回值会共享给所有等待的调用者
// 3. 适合用于缓存查询、数据库查询等场景
// 
// 例如在缓存失效时，多个请求同时来查询同一个 key 的数据，
// singleflight 可以保证只有一个请求真正去数据库查询，
// 其他请求则等待这个结果，从而提高系统性能。


type user1 struct{}

func getuserByID(sg *singleflight.Group, id int) user1 {
	// 使用 id 作为 key
	// 第二个参数是个回调
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
			defer wg.Done()  //释放资源用
			//这里虽败呢给了一个id
			_ = getuserByID(&sg, 1024)
		}()
	}

	wg.Wait()
}
