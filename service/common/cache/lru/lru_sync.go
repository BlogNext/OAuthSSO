package lru

import (
	"container/list"
	"sync"
	"time"
)

//key 是接口，可以是任意类型
type Key interface{}

//一个 entry 包含一个 key 和一个 value，都是任意类型
type entry struct {
	key    Key
	value  interface{}
	expire *time.Time //过期时间
}

//判断是否过期
func (e *entry) IsExpire() bool {
	if e.expire != nil {
		nowTime := time.Now()
		if nowTime.After(*e.expire) {
			//过期了
			return true
		}
	}
	//没有过期
	return false

}

const (
	DefaultMaxEntries = 30
)

//lrucache 是一个 LRU Cache.线程安全
//lrucache算法的数据结构是哈希双向链表
type LruCache struct {
	//MaxEntries 是 Cache 中实体的最大数量，0 表示没有限制,修改了下，默认30个
	MaxEntries int

	//OnEvicted 是一个可选的回调函数，当一个实体从 Cache 中被移除时执行
	OnEvicted func(key Key, value interface{})

	//ll是一个双向链表指针，执行一个 container/list 包中的双向链表
	ll *list.List

	//cache 是一个 map，存放具体的 k/v 对，value 是双向链表中的具体元素，也就是 *Element
	cache map[interface{}]*list.Element

	//读写锁
	rwMutex sync.RWMutex
}

//创建一个 LRU Cache。maxEntries 为 0 表示缓存没有大小限制
func NewLruCache(maxEntries int) *LruCache {

	if maxEntries <= 0 {
		maxEntries = DefaultMaxEntries
	}

	return &LruCache{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

//传入一个 key，返回一个是否有该 key 以及对应 value
//value 值，
//ok ，true有值，false没有值
func (c *LruCache) Get(key Key) (value interface{}, ok bool) {
	//添加读锁
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	if ele, hit := c.cache[key]; hit {

		entry := ele.Value.(*entry)
		if entry.IsExpire() {
			//到期了,不做任何事情,在新增的时候判断是否过期，然后删除，这里不删除，不然会死锁
			return nil, false
		}

		c.ll.MoveToFront(ele)

		return entry.value, true
	}

	return nil, false
}

//向cache中插入一个K=>V,
func (c *LruCache) Add(key Key, value interface{}, expires time.Duration) {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	if ee, ok := c.cache[key]; ok {
		entry := ee.Value.(*entry)
		//若果元素已经存在，把它更新到表头
		c.ll.MoveToFront(ee)
		entry.value = value
		//延长更新时间
		expiresTime := time.Now().Add(expires)
		entry.expire = &expiresTime
		return
	}

	//元素第一次访问,进入缓存
	var ele *list.Element
	if expires != 0 {
		//有过期时间，设计过期时间
		expiresTime := time.Now().Add(expires)
		ele = c.ll.PushFront(&entry{key, value, &expiresTime})
	} else {
		ele = c.ll.PushFront(&entry{key: key, value: value, expire: nil})
	}

	c.cache[key] = ele

	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		//从 Cache 中删除最久未被访问的数据
		c.RemoveOldest()
	}

}

//从 Cache 中删除最久未被访问的数据
func (c *LruCache) RemoveOldest() {
	//加入写锁
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	//返回双向列表中最后一个元素
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

//从 Cache 中删除一个元素，供内部调用
func (c *LruCache) removeElement(e *list.Element) {
	//加入写锁
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	//先从 list 中删除
	c.ll.Remove(e)

	kv := e.Value.(*entry)

	//再从 map 中删除
	delete(c.cache, kv.key)

	//如果回调函数不为空则调用
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

//获取 Cache 当前的元素个数
func (c *LruCache) Len() int {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	return c.ll.Len()
}

//清空 Cache
func (c *LruCache) Clear() {

	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.ll = nil
	c.cache = nil
}

//定时删除过期的元素
func (c *LruCache) TimingRemoveExpire() {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	for key, e := range c.cache {
		entry := e.Value.(*entry)
		if entry.IsExpire() {
			c.removeElement(c.cache[key])
		}
	}
}
