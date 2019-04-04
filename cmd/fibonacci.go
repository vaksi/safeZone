package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vaksi/safeZone/services"
)

var Fibonacci = &cobra.Command{
	Use:   "fibonacci",
	Short: "fibonacci of n integer number",
	Long: `given n integer number will prints output of fibonacci number`,
	Run: fibonacciCmdRun,
}

func init() {
	Fibonacci.PersistentFlags().Int64("n", 0, "initiate variable n")
}

func fibonacciCmdRun(cmd *cobra.Command, args []string) {
	n, _ := cmd.Flags().GetInt64("n")

	fibonacciSvc := &services.FibonacciService{N: n}
	fmt.Printf("Prime of n=%d is %d \n", n, fibonacciSvc.Fibonacci())
}
