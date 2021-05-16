package cmd

import (
	"fmt"
	"strings"

	"github.com/jkulton/pls/internal/collection"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(upcaseCmd)
}

var upcaseCmd = &cobra.Command{
	Use:   "upcase",
	Short: "Upcase a string",
	Run: func(cmd *cobra.Command, args []string) {
		upper := collection.MapStrings(args, strings.ToUpper)
		s := strings.Join(upper, " ")

		fmt.Println(s)
	},
}
