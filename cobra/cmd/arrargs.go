/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// arrargsCmd represents the arrargs command
var arrargsCmd = &cobra.Command{
	Use:   "arrargs",
	Short: "Use several args",
	Long:  `Several args example`,
	Run:   arrArgs,
}

func init() {
	rootCmd.AddCommand(arrargsCmd)

	// arrargsCmd.PersistentFlags().String("foo", "", "A help for foo")
	arrargsCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

func arrArgs(cmd *cobra.Command, args []string) {
	fmt.Println("arrargs called")

	fstatus, _ := cmd.Flags().GetBool("float")
	if fstatus { // if status is true, call addFloat
		addFloat(args)
	} else {
		addInt(args)
	}
}

func addInt(args []string) {
	var sum int

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}

	fmt.Printf("arrargs int %s is %d", args, sum)
}

func addFloat(args []string) {
	var sum float64
	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + ftemp
	}

	fmt.Printf("arrargs floating numbers %s is %f", args, sum)
}
