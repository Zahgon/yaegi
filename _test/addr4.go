package main

const jsonData = `[
  "foo",
  "bar"
]`

const jsonData2 = `[
  {"foo": "foo"},
  {"bar": "bar"}
]`

const jsonData3 = `{
  "foo": "foo",
  "bar": "bar"
}`

func fromSlice() { _ = "STUB: not implemented"; return }

func fromEmpty() { _ = "STUB: not implemented"; return }

func sliceOfObjects() { _ = "STUB: not implemented"; return }

func intoMap() { _ = "STUB: not implemented"; return }

func main() {
	fromSlice()
	fromEmpty()
	sliceOfObjects()
	intoMap()
}
