package main

/*
https://refactoring.guru/design-patterns/template-method

Template Method is a behavioral design pattern that defines
the skeleton of an algorithm in the superclass but lets subclasses
override specific steps of the algorithm without changing its structure.

Imagine that you’re creating a data mining application that analyzes
corporate documents. Users feed the app documents in various formats
(PDF, DOC, CSV), and it tries to extract meaningful data from these
docs in a uniform format.

The first version of the app could work only with DOC files.
In the following version, it was able to support CSV files.
A month later, you “taught” it to extract data from PDF files.

While the code for dealing with various data formats was entirely
different in all classes, the code for data processing and analysis
is almost identical. Wouldn’t it be great to get rid of the code duplication,
leaving the algorithm structure intact?

The Template Method pattern suggests that you break down an algorithm into
a series of steps, turn these steps into methods, and put a series of calls
to these methods inside a single template method. The steps may either be abstract,
or have some default implementation.

As you can see, we’ve got two types of steps:

- abstract steps must be implemented by every subclass
- optional steps already have some default implementation,
but still can be overridden if needed

There’s another type of step, called hooks.
A hook is an optional step with an empty body.
A template method would work even if a hook isn’t overridden.
Usually, hooks are placed before and after crucial steps of algorithms,
providing subclasses with additional extension points for an algorithm.
*/
