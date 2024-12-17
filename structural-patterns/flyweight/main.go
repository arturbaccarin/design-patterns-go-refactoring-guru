package main

/*
https://refactoring.guru/design-patterns/flyweight

Commit of the day

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
*/
