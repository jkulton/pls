package cmd

import (
	"fmt"
	"strings"

	"crypto/sha256"
	"encoding/hex"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(SHA256Cmd)
}

var SHA256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "Get hex SHA256 representation of string.",
	Run: func(cmd *cobra.Command, args []string) {
		s := strings.Join(args, " ")
		h := sha256.New()
		h.Write([]byte(s))
		hs := hex.EncodeToString(h.Sum(nil))

		fmt.Println(hs)
	},
}
