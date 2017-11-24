package toy8

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
	"unsafe"
)

// 临时对象池

// 过多 goroutine 同时创建对象时
// 并发大 - ram - GC 缓慢 - 并发能力降低 ... => sync.Pool

// 对象池 每个 goroutine 不再自己单独创建对象 而是从对象池中取出一个对象（if exists)
// GET PUT NEW

// gomemcache
// keyBufPool returns []byte buffers for use by PickServer's call to
// crc32.ChecksumIEEE to avoid allocations (but doesn't avoid the copies,
// which at least are bounded in size and small)
// var keyBufPool = sync.Pool{
// 	New: func() interface{} {
// 		b := make([]byte, 256)
// 		return &b
// 	},
// }
// func (ss *ServerList) PickServer(key string) (net.Addr, error) {
// 	ss.mu.RLock()
// 	defer ss.mu.RUnlock()
// 	if len(ss.addrs) == 0 {
// 		return nil, ErrNoServers
// 	}
// 	if len(ss.addrs) == 1 {
// 		return ss.addrs[0], nil
// 	}
// 	bufp := keyBufPool.Get().(*[]byte)
// 	n := copy(*bufp, key)
// 	cs := crc32.ChecksumIEEE((*bufp)[:n])
// 	keyBufPool.Put(bufp)

// 	return ss.addrs[cs%uint32(len(ss.addrs))], nil
// }

// sync.Pool
// Get Put
var pool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func DemoSyncPool1(w io.Writer, key, val string) {
	b := pool.Get().(*bytes.Buffer)

	// sure the addr of all bs are the same
	ptr := fmt.Sprintf("%v", unsafe.Pointer(b))

	b.Reset()

	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteString(" " + ptr + " ")
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	b.WriteByte('\n')
	w.Write(b.Bytes())

	pool.Put(b)
}

// 需要频繁 申请 和 GC 的对象
// sync.Pool其实不适合用来做持久保存的对象池（比如连接池）
// 它更适合用来做临时对象池，目的是为了降低GC的压力。

// sync.Once
