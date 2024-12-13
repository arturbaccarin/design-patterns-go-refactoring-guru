package main

/*
https://refactoring.guru/design-patterns/facade

Facade is a structural design pattern that provides a simplified interface to a library,
a framework, or any other complex set of classes.

A facade is a class that provides a simple interface to a complex subsystem which contains
lots of moving parts.
A facade might provide limited functionality in comparison to working with the subsystem directly.
However, it includes only those features that clients really care about.

Having a facade is handy when you need to integrate your app with a sophisticated
library that has dozens of features, but you just need a tiny bit of its functionality.

For instance, an app that uploads short funny videos with cats to social media could potentially
use a professional video conversion library. However, all that it really needs is a
class with the single method encode(filename, format).
After creating such a class and connecting it with the video conversion library,
you’ll have your first facade.

Facade is a structural design pattern that provides a simplified (but limited)
interface to a complex system of classes, library or framework.

While Facade decreases the overall complexity of the application,
it also helps to move unwanted dependencies to one place.
*/

/*
It’s easy to underestimate the complexities that happen behind the scenes when you
order a pizza using your credit card. There are dozens of subsystems that are acting in this process.
Here’s just a shortlist of them:

    1. Check account
    2. Check security PIN
    3. Credit/debit balance
    4. Make ledger entry
    5. Send notification

In a complex system like this, it’s easy to get lost and easy to break stuff if you’re doing something wrong.
That’s why there’s a concept of the Facade pattern:

a thing that lets the client work with dozens of components using a simple interface.

The client only needs to enter the card details, the security pin, the amount to pay, and the operation type.
The Facade directs further communications with various components without exposing the client to internal complexities.
*/
