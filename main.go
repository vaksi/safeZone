package main

import (
	"github.com/spf13/cobra"
	"github.com/vaksi/safeZone/cmd"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "safeZone",
	Short: "safeZone - safe_zone services",
	Long:  `safeZone is mini calculator services`,
	Args:  cobra.MinimumNArgs(1),
}

func main() {
	rootCmd.AddCommand(
		cmd.Sum,
		cmd.Multiply,
		cmd.Prime,
		cmd.Fibonacci,
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
		os.Exit(1)
	}
}