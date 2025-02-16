package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert contact",
	Long:  `Insert contact (name, surname, telephone) to application's database.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("empty name")
			return
		}
		surname, _ := cmd.Flags().GetString("surname")
		if surname == "" {
			fmt.Println("empty surname")
			return
		}
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
		entry := NewEntry(name, surname, telForm)
		insert(entry)
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringP("name", "n", "", "contact name")
	insertCmd.Flags().StringP("surname", "s", "", "contact surname")
	insertCmd.Flags().StringP("telephone", "t", "", "contact telephone")

}

func matchTel(t string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.Match([]byte(t))
}

func insert(entry *Entry) error {
	if _, ok := index[entry.Tel]; ok {
		return fmt.Errorf("telephone %s is already exists", entry.Tel)
	}
	data = append(data, *entry)
	if err := saveJSONFile(JSONFILE); err != nil {
		return err
	}
	return nil
}
