package main

import (
	"LearningGolang/toy8"
	"os"
)

func main() {
	toy8.DemoSyncPool1(os.Stdout, "path", "/search?q=flowers")
	toy8.DemoSyncPool1(os.Stdout, "path", "/search?q=fruits")
	toy8.DemoSyncPool1(os.Stdout, "path", "/search?q=vagetables")
	toy8.DemoSyncPool1(os.Stdout, "path", "/search?q=sugurs")
}
