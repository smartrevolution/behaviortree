package behaviortree

import (
	"context"
	"fmt"
	"strings"
)

type trace struct {
	behaviorName string
	description  string
	level        int
	parent       *trace
}

type key int

const (
	debug key = iota
	traceInfo
)

// WithDebugging allows you to trace a BehaviorTree.
func WithDebugging(ctx context.Context) context.Context {
	return withTrace(context.WithValue(ctx, debug, true), newTrace())
}

func updateContextIfDebugOn(ctx context.Context, descr string) context.Context {
	if isDebugging(ctx) {
		parent := getTrace(ctx)
		trace := &trace{
			description: descr,
			level:       parent.level + 1,
		}
		ctx = withTrace(ctx, trace)
	}
	return ctx
}

func isDebugging(ctx context.Context) bool {
	debug, _ := ctx.Value(debug).(bool)
	return debug
}

func newTrace() *trace {
	return &trace{
		description: "BehaviorTree",
		level:       0,
	}
}

func withTrace(ctx context.Context, trace *trace) context.Context {
	return context.WithValue(ctx, traceInfo, trace)
}

func (t *trace) indention() string {
	var sb strings.Builder
	for i := 0; i < t.level; i++ {
		sb.WriteRune(' ')
	}
	return sb.String()
}

func (t *trace) String() string {
	return fmt.Sprintf("%s%s", t.indention(), t.description)
}

func getTrace(ctx context.Context) *trace {
	trace, _ := ctx.Value(traceInfo).(*trace)
	return trace
}

func printTraceIfDebugOn(ctx context.Context, status Status) {
	if isDebugging(ctx) {
		trace := getTrace(ctx)
		fmt.Println(trace, status)
	}
}
