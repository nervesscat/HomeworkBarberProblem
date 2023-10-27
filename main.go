package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	isBarberSleep bool = true
	queue              = Queue{}
	maxClients         = 5
)

type Queue []string

func (q *Queue) Enqueue(value string) {
	if len(*q) == maxClients {
		fmt.Println("La barbería está llena, no se puede atender al cliente")
		return
	}
	*q = append(*q, value)
}

func (q *Queue) Dequeue() (string, error) {
	if len(*q) == 0 {
		return "", fmt.Errorf("queue is empty")
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, nil
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) IsFull() bool {
	return len(*q) == cap(*q)
}

func attendingClients() {
	for {
		time.Sleep(10 * time.Second)
		if !queue.IsEmpty() {
			value, _ := queue.Dequeue()
			fmt.Println("Atendiendo cliente:", value)
		} else {
			fmt.Println("No hay clientes que atender")
			isBarberSleep = true
			fmt.Println("El barbero se durmió")
		}
	}
}

func creatingClients() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Ingresa el nombre del cliente: ")
		text, _ := reader.ReadString('\n')
		if isBarberSleep {
			fmt.Println("El barbero se despertó")
			isBarberSleep = false
		}
		queue.Enqueue(text)
	}
}

func main() {
	go attendingClients()
	creatingClients()
}
