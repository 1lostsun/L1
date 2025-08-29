package main

import "fmt"

//применимость паттерна : если в проекте есть код, который устарел или закрыт для изменения,
//к примеру, сторонняя библиотека,
//для него пишется адаптер, который будет скрывать работу этого кода
//вызывая "под капотом" тот самый код, но без видимости для клиента.

// Плюсы : 1. не нужно переписывать старый код; 2. клиент не знает об использовании устаревшего кода

// Минусы : 1. архитектура проекта засоряется лишними адаптерами по сути скрывая проблему,
//а не решая ее; 2. повышение расходов на доп вызовы, тк в коде появляется косвенность

// Пример использования : Адаптация скачивания файлов с устройств с различными ОС.
// Клиент -> адаптер ОС -> скачивание и передача из Linux/Windows -> адаптер ОС -> Клиент

func main() {
	oldWorker := &legacyWorker{}
	adapter := &workerAdapter{oldWorker}
	adapter.modernWorking()
}

type modernWorker interface {
	modernWorking()
}

type legacyWorker struct{}

func (w *legacyWorker) work() {
	fmt.Println("some worker is working")
}

type workerAdapter struct {
	lw *legacyWorker
}

func (a *workerAdapter) modernWorking() {
	a.lw.work()
}
