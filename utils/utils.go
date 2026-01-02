package utils

import (
	"fmt"
	"os"

	"github.com/untrustedmodders/go-plugify"
)

func CreateManifest(name, version, author string, dependencies []string) {
	err := plugify.Generate("./...", "./build/", name, version, "", author, "", "", make([]string, 0), dependencies, make([]string, 0), name, "main")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating plugin manifest: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Manifet and autoexports successfully generated!")

	os.Exit(0)
}
