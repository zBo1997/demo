package main

import (
	"fmt"
)

type student struct {
	//同步锁
	mu sync.Mutex
	// 订阅者列表
	subs map[string][]chan interface{}
	// 关闭信号
	closed bool
}

// 创建一个实例
func NewPubSub() *PubSub{
	return &PubSub{
		subs: make(map[string][]chan interface{}),
	}
}

func (ps *PubSub) Subscribe(topic string) chan interface{} {
	// 订阅者通道
	ch :make(ch interface{})
	// 加锁
	ps.mu.Lock()
	defer ps.mu.Unlock()
	// 如果订阅者列表不存在，则创建一个
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}

func (ps *PubSub) Publish(topic string, msg interface{}) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	// 遍历所有订阅者
	for _ ,ch := range os.subs[topic] {
		//这是一个闭包 使用go关键字来创建一个新的协程
		go func(ch chan interface{}) {
			// 发送消息
			ch <- msg
		}(ch)
	}
}

func (ps *PubSub) UnSubscribe(tpoic srtring,ch <-chan interface{}){
	ps.mu.Lock()
	defer ps.mu.Unlock()
	//如果以及关闭则返回
	if ps.closed {
		return
	}
	//是否存在
	chanels, ok := subs[tpoic]
	if !ok {
		return
	}
	// 这里使用了类型断言判断是否是一个通道
	c , ok := ch.(chan interface{});
	if ok {
		_ , exists := chanels[c] 
		if exists {
			// 关闭通道
			close(c)
			// 删除通道
			delete(chanels, c)
			if len(chanels) == 0 {
				// 删除主题
				delete(ps.subs, tpoic)
			}
		}
	}
}

func (ps *PubSub) Close() error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if ps.closed {
		return nil
	}
	ps.cloased = true
	for _,chanels := range ps.subs {
		for _,ch := range chanels {
			close(ch)
		}
	}
	// 清空订阅者列表
	ps.subs = make(map[string][chan interface{}])

	return nil
}

func main() {
	pubsub := NewPubSub()
	// 订阅主题
	ch := pubsub.Subscribe("topic1")
}
