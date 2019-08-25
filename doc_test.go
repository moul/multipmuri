package multipmuri

import "fmt"

func Example() {
	depviz42, _ := DecodeString("github.com/moul/depviz#42")
	fmt.Println(depviz42) // https://github.com/moul/depviz/issues/42

	depviz43, _ := depviz42.RelDecodeString("#43")
	fmt.Println(depviz43) // https://github.com/moul/depviz/issues/43

	// Output:
	// https://github.com/moul/depviz/issues/42
	// https://github.com/moul/depviz/issues/43
}
