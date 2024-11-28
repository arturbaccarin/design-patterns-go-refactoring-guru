package main

/*
Singleton is a creational design pattern that lets you ensure that a class has only one instance,
while providing a global access point to this instance.

Why would anyone want to control how many instances a class has?
The most common reason for this is to control access to some shared resource—for example, a database or a file.

Remember those global variables that you (all right, me) used to store some essential objects?
While they’re very handy, they’re also very unsafe since any code can potentially overwrite the contents of those variables
and crash the app.

Just like a global variable, the Singleton pattern lets you access some object from anywhere in the program.
However, it also protects that instance from being overwritten by other code.

Singleton has almost the same pros and cons as global variables.
Although they’re super-handy, they break the modularity of your code.
*/

/*
We can create a single instance inside the init function.
This is only applicable if the early initialization of the instance is ok.
The init function is only called once per file in a package, so we can be sure that only a single instance will be created.
*/

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

// The sync.Once will only perform the operation once. See the code below
var once sync.Once

type singleWithOnce struct {
}

var singleWithOnceInstance *singleWithOnce

func getOnceInstance() *singleWithOnce {
	if singleWithOnceInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating singleWithOnce instance now.")
				singleWithOnceInstance = &singleWithOnce{}
			})
	} else {
		fmt.Println("SingleWithOnce instance already created.")
	}

	return singleWithOnceInstance
}

func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}

	for i := 0; i < 30; i++ {
		go getOnceInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
