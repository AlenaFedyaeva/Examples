/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "A brief description of even",
	Long:  `A long descrioption of even`,
	Run: func(cmd *cobra.Command, args []string) {
		var evenSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 == 0 {
				evenSum = evenSum + itemp
			}
		}
		fmt.Printf("The even addition of %s is %d", args, evenSum)
	},
}

func init() {
	arrargsCmd.AddCommand(evenCmd)

	// evenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// evenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
