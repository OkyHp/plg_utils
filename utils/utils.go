package utils

import (
	"fmt"
	"os"

	"github.com/OkyHp/plg_utils/s2sdk"
	"github.com/untrustedmodders/go-plugify"
)

func CreateManifest(name, version, author string, dependencies []string) {
	err := plugify.Generate("./...", "", name, version, "", author, "", "", make([]string, 0), dependencies, make([]string, 0), name, "main")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating plugin manifest: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ Manifet and autoexports successfully generated!")

	os.Exit(0)
}

func Uint32ToIPv4(ip int32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24),
		byte(ip>>16),
		byte(ip>>8),
		byte(ip),
	)
}

func GetClientLanguageEx(playerSlot int32) string {
	lang := "en"

	buff := s2sdk.GetClientLanguage(playerSlot)
	if len(buff) >= 2 {
		lang = buff[:2]
	}

	return lang
}
