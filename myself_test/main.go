package main

import "fmt"

type Idle struct {
	Name  string
	Level int
}

func (p *Idle) attack() {
	// 这里可以添加攻击逻辑
	fmt.Printf("%s-Attacking...\n", p.Name)
}

func (p *Idle) defend() {
	// 这里可以添加防御逻辑
	fmt.Printf("%s-Defending...\n", p.Name)
}

func main() {

	var count int = 10

	//创建两个角色
	var role1 = Idle{
		Name:  "角色1",
		Level: 1,
	}

	var role2 = Idle{
		Name:  "角色2",
		Level: 1,
	}
	//使用channel 10个回合依次攻击和防御
	var ch = make(chan func(), count)

	go func() {
		for i := 0; i < count; i++ {
			ch <- role1.attack
			ch <- role2.defend
		}

		close(ch)
	}()

	for task := range ch {
		task()
	}

}
