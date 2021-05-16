package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uuidCmd)
}

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Print a random UUID v4",
	Run: func(cmd *cobra.Command, args []string) {
		newUUID, err := uuid.NewRandom()

		if err != nil {
			fmt.Println("Something went wrong.")
		}

		fmt.Println(newUUID)
	},
}
