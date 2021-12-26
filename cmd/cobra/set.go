package cobra

import (
	"fmt"
	"github.com/spf13/cobra"
	"secret"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s set successfully!\n", key)
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
