package main

import "fmt"

//Первая заглавная буква в названии типов, разрешает использовать их в другом файле, в другом пакете. В этом же пакете необязательно заглавную.

type queue struct {
	list LinkedList
}

func (queue *queue) Enter(number int) {
	queue.list.AppEnd(number)
}
func (queue *queue) Leave() int {
	n := queue.list.getElement(0)
	queue.list.DeletFirst()
	return n
}
func main() {
	fmt.Println(`Создаём пустую очередь`)
	var myQueue queue
	fmt.Println(myQueue)

	fmt.Println(`Запускаем в очередь: `)
	myQueue.Enter(1)
	myQueue.list.Printline()

	fmt.Println(`Запускаем в очередь: `)
	myQueue.Enter(2)
	myQueue.Enter(3)
	myQueue.list.Printline()

	fmt.Println(`Заставляем покинуть очередь один элемент`)
	myQueue.Leave()
	fmt.Println(`Заставляем покинуть очередь ещё один элемент, только сейчас смотрим за элементом`, myQueue.Leave())

	fmt.Println(`Очередь: `)
	myQueue.list.Printline()
}
