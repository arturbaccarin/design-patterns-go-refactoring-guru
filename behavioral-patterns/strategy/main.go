package main

import "fmt"

/*
https://refactoring.guru/design-patterns/strategy

Strategy is a behavioral design pattern that lets you define a family of algorithms,
put each of them into a separate class, and make their objects interchangeable.

One day you decided to create a navigation app for casual travelers. The app was
centered around a beautiful map which helped users quickly orient themselves in any city.

One of the most requested features for the app was automatic route planning.
A user should be able to enter an address and see the fastest route to that
destination displayed on the map.

The first version of the app could only build the routes over roads.
People who traveled by car were bursting with joy. But apparently,
not everybody likes to drive on their vacation. So with the next update,
you added an option to build walking routes. Right after that, you added
another option to let people use public transport in their routes.

Each time you added a new routing algorithm, the main class of the navigator
doubled in size. At some point, the beast became too hard to maintain.

The Strategy pattern suggests that you take a class that does something
specific in a lot of different ways and extract all of these algorithms
into separate classes called strategies.

The original class, called context, must have a field for storing a
reference to one of the strategies. The context delegates the work to
a linked strategy object instead of executing it on its own.

The context isn’t responsible for selecting an appropriate algorithm for the job.
Instead, the client passes the desired strategy to the context. In fact, the context
doesn’t know much about strategies. It works with all strategies through the same
generic interface, which only exposes a single method for triggering the algorithm
encapsulated within the selected strategy.

This way the context becomes independent of concrete strategies, so you can add new
algorithms or modify existing ones without changing the code of the context or other strategies.

In our navigation app, each routing algorithm can be extracted to its own class with a single
buildRoute method. The method accepts an origin and destination and returns a collection of the route’s checkpoints.

the main navigator class doesn’t really care which algorithm is
selected since its primary job is to render a set of checkpoints on the map.

Imagine that you have to get to the airport. You can catch a bus, order a cab,
or get on your bicycle. These are your transportation strategies. You can pick
one of the strategies depending on factors such as budget or time constraints.
*/

/*
Code example:

Suppose you are building an In-Memory-Cache. Since it’s in memory, it has a limited size.
Whenever it reaches its maximum size, some entries have to be evicted to free-up space.
This can happen via several algorithms. Some of the popular algorithms are:

1. Least Recently Used (LRU): remove an entry that has been used least recently.
2. First In, First Out (FIFO): remove an entry that was created first.
3. Least Frequently Used (LFU): remove an entry that was least frequently used.

The problem is how to decouple our cache class from these algorithms so that we can change the
algorithm at run time. Also, the cache class should not change when a new algorithm is being added.

This is where Strategy pattern comes into the picture. It suggests creating a family of the
algorithm with each algorithm having its own class. Each of these classes follows the same interface,
and this makes the algorithm interchangeable within the family. Let’s say the common interface name is evictionAlgo.
*/

// Strategy interface
type EvictionAlgo interface {
	evict(c *Cache)
}

// Concrete strategy
type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}

// Context
type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")
}
