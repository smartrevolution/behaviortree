package main

import (
	"context"
	"fmt"

	"github.com/smartrevolution/behaviortree"
)

func main() {

	//	ctx := WithDebugging(context.Background())
	ctx := context.Background()

	bt := behaviortree.BehaviorTree(
		behaviortree.Sequence(
			behaviortree.Println("Hello World will be printed"),
		),
	)
	bt.Execute(ctx)

	bt = behaviortree.BehaviorTree(
		behaviortree.Sequence(
			behaviortree.Condition(
				behaviortree.Action(func(ctx context.Context) behaviortree.Status {
					return behaviortree.Failure
				}),
			),
			behaviortree.Println("Hello World will NOT be printed"),
		),
	)
	bt.Execute(ctx)

	// IsPlayerVisible := behaviortree.Action(func(ctx context.Context) behaviortree.Status {
	// 	fmt.Println("Player is visible!")
	// 	return behaviortree.Success
	// })

	IsPlayerVisible := behaviortree.Action(func(ctx context.Context) behaviortree.Status {
		fmt.Println("Player is NOT visible!")
		return behaviortree.Failure
	})

	IsPlayerInRange := behaviortree.Action(func(ctx context.Context) behaviortree.Status {
		fmt.Println("Player is in range!")
		return behaviortree.Success
	})

	FireAtPlayer := behaviortree.Action(func(ctx context.Context) behaviortree.Status {
		fmt.Println("Peng!!!")
		return behaviortree.Success
	})

	MoveTowardsPlayer := behaviortree.Action(func(ctx context.Context) behaviortree.Status {
		fmt.Println("Move towards player!")
		return behaviortree.Success
	})

	HaveWeGotASuspectedLocation := behaviortree.Action(func(ctx context.Context) behaviortree.Status {
		fmt.Println("We have a suspected location!")
		return behaviortree.Success
	})

	MoveToPlayersLastKnownLocation := behaviortree.Action(func(ctx context.Context) behaviortree.Status {
		fmt.Println("Move to players last know location!")
		return behaviortree.Success
	})

	MoveToRandomLocation := behaviortree.Action(func(ctx context.Context) behaviortree.Status {
		fmt.Println("Move to random location!")
		return behaviortree.Success
	})

	LookAround := behaviortree.Action(func(ctx context.Context) behaviortree.Status {
		fmt.Println("Looking around!")
		return behaviortree.Success
	})

	ctx = context.Background()
	bt = behaviortree.BehaviorTree(
		behaviortree.Selector(
			behaviortree.Sequence(
				behaviortree.Condition(IsPlayerVisible),
				behaviortree.Selector(
					behaviortree.Sequence(
						behaviortree.Condition(IsPlayerInRange),
						behaviortree.Repeat(FireAtPlayer, 3),
					),
					MoveTowardsPlayer,
				),
			),
			behaviortree.Sequence(
				behaviortree.Condition(HaveWeGotASuspectedLocation),
				MoveToPlayersLastKnownLocation,
				LookAround,
			),
			behaviortree.Sequence(
				MoveToRandomLocation,
				LookAround,
			),
		),
	)
	bt.Execute(ctx)

}
