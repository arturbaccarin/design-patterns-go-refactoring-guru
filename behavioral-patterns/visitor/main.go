package main

/*
https://refactoring.guru/design-patterns/visitor

the last one

Visitor is a behavioral design pattern that lets you separate
algorithms from the objects on which they operate.

Imagine that your team develops an app which works with
geographic information structured as one colossal graph.
Each node of the graph may represent a complex entity such as a city,
but also more granular things like industries, sightseeing areas, etc.

The nodes are connected with others if there’s a road between
the real objects that they represent. Under the hood, each node
type is represented by its own class, while each specific node is an object.

At some point, you got a task to implement exporting the graph into XML format.
At first, the job seemed pretty straightforward. You planned to add an export
method to each node class and then leverage recursion to go over each node of
the graph, executing the export method. The solution was simple and elegant:
thanks to polymorphism, you weren’t coupling the code which called the export
method to concrete classes of nodes.

At some point, you got a task to implement exporting the graph into XML format.
At first, the job seemed pretty straightforward. You planned to add an export
method to each node class and then leverage recursion to go over each node of the graph,
executing the export method. The solution was simple and elegant: thanks to polymorphism,
you weren’t coupling the code which called the export method to concrete classes of nodes.
*/
