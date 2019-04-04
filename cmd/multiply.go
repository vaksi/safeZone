package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vaksi/safeZone/services"
)

var Multiply = &cobra.Command{
	Use:   "multiply",
	Short: "multiply of 2 integer number",
	Long: `given 2 integer number will print output of multiply`,
	Run: multiplyCmdRun,
}

func init() {
	Multiply.PersistentFlags().Int64("x", 0, "initiate variable x")
	Multiply.PersistentFlags().Int64("y", 0, "initiate variable y")
}

func multiplyCmdRun(cmd *cobra.Command, args []string) {
	x, _ := cmd.Flags().GetInt64("x")
	y, _ := cmd.Flags().GetInt64("y")

	mulSvc := &services.MultiplyService{X: x, Y: y}
	fmt.Printf("Multiply of x=%d and y=%d is %d \n", x, y, mulSvc.Multiply())
}
