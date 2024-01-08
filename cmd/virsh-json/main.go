package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/a-h/virshjson"
)

const help = `virsh-json parses virsh output tables into a JSON structure.

Usage:

	virsh list | virsh-json`

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println(help)
		return
	}
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println(help)
		return
	}
	data, err := virshjson.Convert(os.Stdin)
	if err != nil {
		fmt.Printf("failed to convert: %v", err)
		os.Exit(1)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	err = enc.Encode(data)
	if err != nil {
		fmt.Printf("failed to write JSON: %v", err)
		os.Exit(1)
	}
}
