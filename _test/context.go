package main

import "context"

func get(ctx context.Context, k string) string { _ = "STUB: not implemented"; return "" }

func main() {
	ctx := context.WithValue(context.Background(), "hello", "world")
	println(get(ctx, "hello"))
}
