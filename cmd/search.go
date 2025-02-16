package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search entry.",
	Long: `Search saved entry by his phone number.`,
	Run: func(cmd *cobra.Command, args []string) {
		telephone, _ := cmd.Flags().GetString("key")
		if telephone == "" {
			fmt.Println("empty telephone")
			return
		}
		telForm := strings.ReplaceAll(telephone, "-", "")
		if !matchTel(telForm) {
			fmt.Println("not a valid telephone format")
			return
		}
		if err := search(telForm); err != nil {
			fmt.Println(err)
			return 
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("key", "k", "", "telephone for search")
}

func search(tel string) error {
	idx, ok := index[tel]
	if ok {
		return fmt.Errorf("telephone %s doesn't exists", tel)
	}
	entry := data[idx]
	printOne(entry, -1)
	return nil
}