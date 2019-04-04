package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vaksi/safeZone/services"
)

var Prime = &cobra.Command{
	Use:   "prime",
	Short: "prime of n integer number",
	Long: `given n integer number will prints output of prime`,
	Run: primeCmdRun,
}

func init() {
	Prime.PersistentFlags().Int64("n", 0, "initiate variable n")
}

func primeCmdRun(cmd *cobra.Command, args []string) {
	n, _ := cmd.Flags().GetInt64("n")

	primeSvc := &services.PrimeService{N: n}
	fmt.Printf("Prime of n=%d is %d \n", n, primeSvc.Prime())
}
