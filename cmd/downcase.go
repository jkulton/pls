package cmd

import (
	"fmt"
	"strings"

	"github.com/jkulton/pls/internal/collection"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(downcaseCmd)
}

var downcaseCmd = &cobra.Command{
	Use:   "downcase",
	Short: "Downcase a string",
	Run: func(cmd *cobra.Command, args []string) {
		lower := collection.MapStrings(args, strings.ToLower)
		s := strings.Join(lower, " ")

		fmt.Println(s)
	},
}
