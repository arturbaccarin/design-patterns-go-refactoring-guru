package main

/*
https://refactoring.guru/design-patterns/command

Command is a behavioral design pattern that turns a request into a stand-alone
object that contains all information about the request.

This transformation lets you pass requests as a method arguments, delay or
queue a request’s execution, and support undoable operations.

Good software design is often based on the principle of separation of concerns,
which usually results in breaking an app into layers.
The most common example: a layer for the graphical user interface and another layer for
the business logic. The GUI layer is responsible for rendering a beautiful
picture on the screen, capturing any input and showing results of what the user and the app are doing.

However, when it comes to doing something important, like calculating the trajectory
of the moon or composing an annual report, the GUI layer delegates the work to
the underlying layer of business logic.

The Command pattern suggests that GUI objects shouldn’t send these requests directly.

Instead, you should extract all of the request details, such as the object being called,
the name of the method and the list of arguments into a separate command class
with a single method that triggers this request.

Command objects serve as links between various GUI and business logic objects.

From now on, the GUI object doesn’t need to know what business logic object will
receive the request and how it’ll be processed. The GUI object just triggers the command,
which handles all the details.

The next step is to make your commands implement the same interface.
Usually it has just a single execution method that takes no parameters.

Since the command execution method doesn’t have any parameters, how would we pass the
request details to the receiver? It turns out the command should be either
pre-configured with this data, or capable of getting it on its own.

They’ll be linked to a command which gets executed when a user interacts with the GUI element.
As you’ve probably guessed by now, the elements related to the same operations will
be linked to the same commands, preventing any code duplication.

As a result, commands become a convenient middle layer that reduces coupling
between the GUI and business logic layers.

After a long walk through the city, you get to a nice restaurant and sit at
the table by the window. A friendly waiter approaches you and quickly takes your order,
writing it down on a piece of paper.
The waiter goes to the kitchen and sticks the order on the wall.
After a while, the order gets to the chef, who reads it and cooks the meal accordingly.
The cook places the meal on a tray along with the order.
The waiter discovers the tray, checks the order to make sure everything is as you wanted it,
and brings everything to your table.

The paper order serves as a command. It remains in a queue until the chef is ready to
serve it. The order contains all the relevant information required to cook the meal.
It allows the chef to start cooking right away instead of running around clarifying
the order details from you directly.

Commit of the day, because I'm studying other things now. 
*/
