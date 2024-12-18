package main

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
*/

/*
The pattern extracts the repeating intrinsic state from a main Tree class and moves it into the flyweight class TreeType.

Now instead of storing the same data in multiple objects, it’s kept in just a few flyweight objects and linked to appropriate
Tree objects which act as contexts. The client code creates new tree objects using the flyweight factory,
which encapsulates the complexity of searching for the right object and reusing it if needed.
*/
