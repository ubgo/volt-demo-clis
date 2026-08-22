// Command snap is one of two CLIs in the monorepo demo. Its releases are
// tagged snap/vX.Y.Z — the BINARY name, wherever this folder lives.
package main

import (
	"flag"
	"fmt"

	"github.com/khanakia/volt-demo-clis/internal/shared"
)

var version = "dev" // stamped via cmd/snap/.volt.yml

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("snap %s\n", version)
		return
	}
	fmt.Println(shared.Banner("snap"), "— snap!")
}
