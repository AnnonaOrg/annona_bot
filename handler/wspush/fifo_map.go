package wspush

// import (
// 	"sync"
// )

// var fifoMap *FIFOMap

// func init() {
// 	fifoMap = NewFIFOMap()
// }

// // 先进先出（FIFO）的 map 实现
// type FIFOMap struct {
// 	mu    sync.Mutex
// 	keys  []string
// 	items map[string]interface{}
// }

// func NewFIFOMap() *FIFOMap {
// 	return &FIFOMap{
// 		items: make(map[string]interface{}),
// 	}
// }

// func (m *FIFOMap) Set(key string, value interface{}) {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	// 将 key 加入切片尾部，表示最近访问
// 	m.keys = append(m.keys, key)
// 	m.items[key] = value
// }

// func (m *FIFOMap) Get(key string) (interface{}, bool) {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	value, ok := m.items[key]
// 	return value, ok
// }

// func (m *FIFOMap) RemoveOldest() {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	if len(m.keys) == 0 {
// 		return
// 	}

// 	oldestKey := m.keys[0]
// 	// 移除切片头部，表示最老的访问
// 	m.keys = m.keys[1:]
// 	delete(m.items, oldestKey)
// }

// func (m *FIFOMap) Count() int {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	return len(m.keys)
// }
