// Package behaviortree provides primitives for creating and nesting behavior trees
package behaviortree

import (
	"context"
	"fmt"
)

// Status is the return status of a behavior
type Status int

const (
	// Invalid is the default value for the status.
	Invalid Status = iota
	// Running should be returned the Behavior is already running (in the background).
	Running
	// Success should be returned when Behavior was executed successfully.
	Success
	// Failure should be returned when a Behavior execution failed.
	Failure
)

// String prints the return state of a behavior as a string.
func (s Status) String() string {
	return [...]string{"Invalid", "Running", "Success", "Failure"}[s]
}

// Behavior is a function that can be executed and that returns a status.
type Behavior interface {
	Execute(ctx context.Context) Status
}

// Action is used to cast anonymous functions to Actions, which automatically implement the Behavior interface.
type Action func(ctx context.Context) Status

// Execute is the implementation of the Behavior interface.
func (a Action) Execute(ctx context.Context) Status {
	return a(ctx)
}

// Sequence works like AND.
func Sequence(behaviors ...Behavior) Behavior {
	// We are casting an anonymous function to an Action,
	// which in turn implements the Behavior interface
	// and therefore we are returning a Behavior here.
	// Don't overuse this. It is complicated code and you
	// won't remember how it works in a couple of weeks.
	// But in some rare cases it can be quite handy when
	// you want to implement single function interfaces.
	// I would consider this bad engineering in regular code.
	// Except maybe in a small library, that doesn't change much
	// and which is well tested. And when there is really no
	// better way to do this.
	return Action(func(ctx context.Context) Status {
		ctx = updateContextIfDebugOn(ctx, "Sequence")
		var status Status
		for _, fn := range behaviors {
			status = fn.Execute(ctx)
			if status != Success {
				printTraceIfDebugOn(ctx, status)
				return status
			}
		}
		printTraceIfDebugOn(ctx, Success)
		return Success
	})
}

// Selector works like OR.
func Selector(behaviors ...Behavior) Behavior {
	return Action(func(ctx context.Context) Status {
		ctx = updateContextIfDebugOn(ctx, "Selector")
		var status Status
		for _, fn := range behaviors {
			status = fn.Execute(ctx)
			if status != Failure {
				printTraceIfDebugOn(ctx, status)
				return status
			}
		}
		printTraceIfDebugOn(ctx, Failure)
		return Failure
	})
}

// Condition works like IF.
func Condition(condition Behavior) Behavior {
	return Action(func(ctx context.Context) Status {
		ctx = updateContextIfDebugOn(ctx, "Condition")
		status := condition.Execute(ctx)
		if status != Success {
			printTraceIfDebugOn(ctx, status)
			return Failure
		}
		printTraceIfDebugOn(ctx, Success)
		return Success
	})
}

// Invert returns the inverse status of a behavior. It works like a negation.
func Invert(behavior Behavior) Behavior {
	return Action(func(ctx context.Context) Status {
		ctx = updateContextIfDebugOn(ctx, "Invert")
		status := behavior.Execute(ctx)
		if status != Success {
			printTraceIfDebugOn(ctx, Success)
			return Success
		}
		printTraceIfDebugOn(ctx, Failure)
		return Failure
	})
}

// Repeat will repeat the passed behavior for <int> times.
func Repeat(behavior Behavior, times int) Behavior {
	return Action(func(ctx context.Context) Status {
		ctx = updateContextIfDebugOn(ctx, "Repeat")
		var status Status
		for i := 0; i < times; i++ {
			status = behavior.Execute(ctx)
		}
		printTraceIfDebugOn(ctx, status)
		return Success
	})
}

// BehaviorTree creates a new Behaviortree. This is your starting point.
func BehaviorTree(behaviors ...Behavior) Behavior {
	return Action(func(ctx context.Context) Status {
		var status Status
		for _, fn := range behaviors {
			status = fn.Execute(ctx)
		}
		printTraceIfDebugOn(ctx, status)
		return status
	})
}

// Println is a simple behavior for debugging. It prints a msg string.
func Println(msg string) Behavior {
	return Action(func(ctx context.Context) Status {
		ctx = updateContextIfDebugOn(ctx, "Action:Println")
		fmt.Println(msg)
		printTraceIfDebugOn(ctx, Success)
		return Success
	})
}
