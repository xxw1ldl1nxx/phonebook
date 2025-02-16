package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete contact.",
	Long:  `Delete contact with number.`,
	Run: func(cmd *cobra.Command, args []string) {
		telephone, _ := cmd.Flags().GetString("telephone")
		if telephone == "" {
			fmt.Println("empty telephone")
			return
		}
		telForm := strings.ReplaceAll(telephone, "-", "")
		if !matchTel(telForm) {
			fmt.Println("not a valid telephone format")
			return
		}
		if err := deleteEntry(telForm); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("telephone", "t", "", "telephone of user to delete")
}

func deleteEntry(tel string) error {
	idx, ok := index[tel]
	if ok {
		return fmt.Errorf("telephone %s doesn't exists", tel)
	}
	data = append(data[:idx], data[idx+1:]...)
	if err := saveJSONFile(JSONFILE); err != nil {
		return err
	}
	return nil
}
