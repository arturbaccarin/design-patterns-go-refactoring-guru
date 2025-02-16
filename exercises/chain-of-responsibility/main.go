package main

import "fmt"

/*
You are building a customer support system where multiple support
agents handle customer issues based on their severity. The customer
issues come with varying levels of priority, and the support agents
specialize in different levels of severity. Your task is to implement
a chain of responsibility to route these issues to the appropriate support agent.

1.Create a chain of support agents, where each agent can handle a certain level of severity.

2. If an agent can't handle the severity of the issue, the issue is passed to the next agent in the chain.

3. Support agents should have the following severity levels:
Low: Basic issues that require quick resolution.
Medium: Intermediate issues that require more time and expertise.
High: Critical issues that need immediate attention.
*/

type SupportAgent interface {
	Handle(issue *Issue) bool
}

type Agent struct {
	Severity string
	Next     SupportAgent
}

func (a *Agent) Handle(issue *Issue) bool {
	if a.Severity == issue.Severity {
		fmt.Printf("Agent %s handles issue: %s\n", a.Severity, issue.Description)
		return true
	}
	if a.Next != nil {
		return a.Next.Handle(issue)
	}
	return false
}

func NewAgent(severity string, next SupportAgent) *Agent {
	return &Agent{Severity: severity, Next: next}
}

type Issue struct {
	Severity    string
	Description string
}

func main() {
	high := NewAgent("high", nil)
	medium := NewAgent("medium", high)
	low := NewAgent("low", medium)

	low.Handle(&Issue{Severity: "high", Description: "Issue 1"})
	low.Handle(&Issue{Severity: "medium", Description: "Issue 2"})
	low.Handle(&Issue{Severity: "low", Description: "Issue 3"})
}
