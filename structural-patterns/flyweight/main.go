package main

import "fmt"

/*
https://refactoring.guru/design-patterns/flyweight

Flyweight is a structural design pattern that lets you fit more objects into the
available amount of RAM by sharing common parts of state between multiple objects
instead of keeping all of the data in each object.

The actual problem was related to your particle system.
Each particle, such as a bullet, a missile or a piece of shrapnel
was represented by a separate object containing plenty of data.
At some point, when the carnage on a player’s screen reached its climax,
newly created particles no longer fit into the remaining RAM, so the program crashed.

On closer inspection of the Particle class, you may notice that the color and
sprite fields consume a lot more memory than other fields.
What’s worse is that these two fields store almost identical data across all particles.
For example, all bullets have the same color and sprite.

Other parts of a particle’s state, such as coordinates, movement vector and speed,
are unique to each particle. After all, the values of these fields change over time.
This data represents the always changing context in which the particle exists,
while the color and sprite remain constant for each particle.

Intrinsic state: constant data of an object. It lives within the object; other objects can only read it, not change it.

Extrinsic state: the rest of the object’s state, often altered “from the outside” by other objects.

The Flyweight pattern suggests that you stop storing the extrinsic state inside the object.
Instead, you should pass this state to specific methods which rely on it.

Only the intrinsic state stays within the object, letting you reuse it in different contexts.

Where does the extrinsic state move to?
In most cases, it gets moved to the container object, which aggregates objects before we apply the pattern.

To move the extrinsic state into this class, you need to create several array fields for storing coordinates,
vectors, and speed of each individual particle. But that’s not all.
You need another array for storing references to a specific flyweight that represents a particle.
These arrays must be in sync so that you can access all data of a particle using the same index.

Won’t we need to have as many of these contextual objects as we had at the very beginning?
Technically, yes. But the thing is, these objects are much smaller than before.
The most memory-consuming fields have been moved to just a few flyweight objects.
Now, a thousand small contextual objects can reuse a single heavy flyweight object instead of storing a thousand copies of its data.

Since the same flyweight object can be used in different contexts, you have to make sure that its state can’t be modified.
A flyweight should initialize its state just once, via constructor parameters.
It shouldn’t expose any setters or public fields to other objects.

For more convenient access to various flyweights, you can create a factory method that manages
a pool of existing flyweight objects.
The method accepts the intrinsic state of the desired flyweight from a client,
looks for an existing flyweight object matching this state, and returns it if it was found.
If not, it creates a new flyweight and adds it to the pool.

There are several options where this method could be placed. The most obvious place is a flyweight container.
Alternatively, you could create a new factory class. Or you could make the factory method static and put it
inside an actual flyweight class.

Flyweight is a structural design pattern that allows programs to support vast quantities of objects by keeping their
memory consumption low.
*/

/*
In a game of Counter-Strike, Terrorist and Counter-Terrorist have a different type of dress.
For simplicity, let’s assume that both Terrorist and Counter-Terrorists have one dress type each.
The dress object is embedded in the player object as below.

Let’s say there are 5 Terrorists and 5 Counter-Terrorist, so a total of 10 players. Now there are two options concerning dress.

1. Each of the 10 player objects creates a different dress object and embeds them. A total of 10 dress objects will be created.

We create two dress objects:

1. Single Terrorist Dress Object: This will be shared across 5 Terrorists.
2. Single Counter-Terrorist Dress Object: This will be shared across 5 Counter-Terrorist.

As you can see in Approach 1, a total of 10 dress objects are created while in approach 2 only 2 dress objects are created.
The second approach is what we follow in the Flyweight design pattern. The two dress objects which we created are called
the flyweight objects.

The Flyweight pattern takes out the common parts and creates flyweight objects.
These flyweight objects (dress) can then be shared among multiple objects (player).
This drastically reduces the number of dress objects, and the good part is that even if
you create more players, only two dress objects will be sufficient.

Intrinsic State: Dress in the intrinsic state as it can be shared across multiple Terrorist and Counter-Terrorist objects.

Extrinsic State: Player location and the player weapon are an extrinsic state as they are different for every object.
*/

// flyweight interface
type Dress interface {
	getColor() string
}

// flyweight factory
const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerroristDressType terrorist dress type
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}

// Concrete flyweight object
type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}

type TerroristDress struct {
	color string
}

func (t *TerroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

// Context
type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

// Client code
type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (c *game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}

func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
