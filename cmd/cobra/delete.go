package cobra

import (
	"fmt"

	"secret"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a secret from your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key := args[0]
		err := v.Delete(key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s deleted successfully!\n", key)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
