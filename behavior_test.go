package behaviortree

import (
	"context"
	"testing"
)

func fail() Behavior {
	return Action(func(ctx context.Context) Status {
		return Failure
	})
}

func succeed() Behavior {
	return Action(func(ctx context.Context) Status {
		return Success
	})
}

func run() Behavior {
	return Action(func(ctx context.Context) Status {
		return Running
	})
}

func TestBehaviors(t *testing.T) {
	tests := []struct {
		behavior Behavior
		expected Status
	}{
		{
			behavior: BehaviorTree(),
			expected: Invalid,
		},
		{
			behavior: BehaviorTree(Sequence()),
			expected: Success,
		},
		{
			behavior: BehaviorTree(Selector()),
			expected: Failure,
		},
		{
			behavior: BehaviorTree(Sequence(fail(), succeed())),
			expected: Failure,
		},
		{
			behavior: BehaviorTree(Sequence(succeed(), fail())),
			expected: Failure,
		},
		{
			behavior: BehaviorTree(Selector(fail(), succeed())),
			expected: Success,
		},
		{
			behavior: BehaviorTree(Selector(succeed(), fail())),
			expected: Success,
		},
	}

	for _, test := range tests {
		got := test.behavior.Execute(context.Background())
		if got != test.expected {
			t.Errorf("Got: %s, Expected: %s", got, test.expected)
		}
	}
}
