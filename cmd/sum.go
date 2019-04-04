package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vaksi/safeZone/services"
)

var Sum = &cobra.Command{
	Use:   "sum",
	Short: "sum of 2 integer number",
	Long: `given 2 integer number will print output of sum`,
	Run: sumCmdRun,
}

func init() {
	Sum.PersistentFlags().Int64("x", 0, "initiate variable x")
	Sum.PersistentFlags().Int64("y", 0, "initiate variable y")
}

// SumCmdRun
func sumCmdRun(cmd *cobra.Command, args []string) {
	x, _ := cmd.Flags().GetInt64("x")
	y, _ := cmd.Flags().GetInt64("y")

	sumSvc := &services.SumService{X: x, Y: y}
	fmt.Printf("Sum of x=%d and y=%d is %d \n", x, y, sumSvc.Sum())
}