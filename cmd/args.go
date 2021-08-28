package cmd

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/madpin/gololo/sum"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(argsCmd)
}

var argsCmd = &cobra.Command{
	Use:   "args",
	Short: "Print the version number of Gololo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Inside subCmd Run with args: %v\n", args)

		first_arg, _ := strconv.Atoi(args[0])
		second_arg, _ := strconv.Atoi(args[1])
		fmt.Println(first_arg, reflect.TypeOf(first_arg))
		fmt.Println("Sum result: %d\n", sum.Sum(first_arg, second_arg))
	},
}
