package design_model

import (
	"fmt"
)

/**
观察者模式
1. 定义
	定义了对象间的一种一对多的依赖关系，以便一个对象的状态发生变化时，所有依赖于它的对象都得到通知并自动刷新。
2. 实现步骤
	1. 定义触发事件 Event
	2. 定义通知人 Notifier
	3. 定义观察者 Observer
	4. 注册通知列表，当触发事件时，进行通知
3. 使用场景
	一个事件会触发多个业务，比如在一个系统中，账户的余额变动会触发数据库的余额更改、增加流水记录、是否满足活动要求等。进入活动观察者，如果活动过期就删除此观察者
4. 优点
	观察者模式将观察者和被观察的对象分离开,体现了面向对象设计中一个对象只做一件事情的原则，提高了应用程序的可维护性和重用性。
 */

type (
	// Event defines an indication of a point-in-time occurrence.
	Event struct {
		// Data in this case is a simple int, but the actual
		// implementation would depend on the application.
		Data int64
	}

	// Observer defines a standard interface for instances that wish to list for
	// the occurrence of a specific event.
	Observer interface {
		// OnNotify allows an event to be "published" to interface implementations.
		// In the "real world", error handling would likely be implemented.
		OnNotify(Event)
	}

	// Notifier is the instance being observed. Publisher is perhaps another decent
	// name, but naming things is hard.
	Notifier interface {
		// Register allows an instance to register itself to listen/observe
		// events.
		Register(Observer)
		// Deregister allows an instance to remove itself from the collection
		// of observers/listeners.
		Deregister(Observer)
		// Notify publishes new events to listeners. The method is not
		// absolutely necessary, as each implementation could define this itself
		// without losing functionality.
		Notify(Event)
	}
)

type (
	eventObserver struct{
		id int
	}

	eventNotifier struct{
		// Using a map with an empty struct allows us to keep the observers
		// unique while still keeping memory usage relatively low.
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

func (p *eventNotifier) Notify(e Event) {
	for o := range p.observers {
		o.OnNotify(e)
	}
}