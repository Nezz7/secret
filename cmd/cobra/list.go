package cobra

import (
	"fmt"

	"secret"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all secret keys in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		keys, err := v.List()
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(keys) == 0 {
			fmt.Println("No keys are stored.")
			return
		}
		fmt.Println("Stored keys :")
		for _, k := range keys {
			fmt.Println(k)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
