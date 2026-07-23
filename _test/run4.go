package main

type fn func(int)

func f1(i int) { _ = "STUB: not implemented"; return }

func test(f fn, v int) { _ = "STUB: not implemented"; return }

func main() { test(f1, 21) }
