package main

import (
	"fmt"
	"sync"
)

type PubSub struct {
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
	ch := make(chan interface{})
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
	if _, ok := ps.subs[topic]; !ok {
		panic("topic not found")
	}
	// 遍历所有订阅者
	for _ ,ch := range ps.subs[topic] {
		//这是一个闭包 使用go关键字来创建一个新的协程
		go func(ch chan interface{}) {
			// 发送消息
			ch <- msg
		}(ch)
	}
	
}

func (ps *PubSub) UnSubscribe(topic string,ch <-chan interface{}){
	ps.mu.Lock()
	defer ps.mu.Unlock()
	//如果以及关闭则返回
	if ps.closed {
		return
	}
	//是否存在
	chanels, ok := ps.subs[topic]
	if !ok {
		return
	}
	// 遍历找到匹配的通道
	for i, c := range chanels {
		if c == ch {
			// 关闭通道
			close(c)
			// 删除通道 从i开始到最后一个进行覆盖到新的通道来进行删除
			chanels = append(chanels[:i], chanels[i+1:]...)
			if len(chanels) == 0 {
				// 删除主题
				delete(ps.subs, topic)
			} else {
				ps.subs[topic] = chanels
			}
			break
		}
	}
}

func (ps *PubSub) Close() error {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if ps.closed {
		return nil
	}
	ps.closed = true
	for _,chanels := range ps.subs {
		for _,ch := range chanels {
			close(ch)
		}
	}
	// 清空订阅者列表
	ps.subs = make(map[string][]chan interface{})
	return nil
}

func main() {
	pubsub := NewPubSub()
	// 订阅主题
	ch := pubsub.Subscribe("topic1")
	// 发布消息 使用recover来捕获异常
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	pubsub.Publish("topic", "Hello World")

	// 接收消息
	msg := <-ch
	fmt.Println("Received message:", msg)
}
