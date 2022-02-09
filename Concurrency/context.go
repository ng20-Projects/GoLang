package main

import (
	"context"
	"fmt"
)

func seedContext(ctx context.Context) context.Context {
	// ctx = context.WithValue(ctx, "one", "111")
	// ctx = context.WithValue(ctx, 1, "1111")
	ctx = context.WithValue(ctx, 1, "111")
	return ctx
}

func readCtx(ctx context.Context) {
	// value := ctx.Value("one")
	value := ctx.Value(1)
	fmt.Println(value)
}

func main() {
	ctx := context.Background()
	//seed dome data in ctx
	ctx = seedContext(ctx)
	readCtx(ctx)
}
