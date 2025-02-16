package cmd

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var green = color.New(color.FgGreen).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of entry.",
	Long:  `Shows full list of recorded entrys.`,
	Run: func(cmd *cobra.Command, args []string) {
		printList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func (b phoneBook) Len() int {
	return len(b)
}

func (b phoneBook) Less(i, j int) bool {
	if b[i].Surname == b[j].Surname {
		return b[i].Name < b[j].Name
	}
	return b[i].Surname < b[j].Surname
}

func (b phoneBook) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func printList() {
	sort.Sort(data)
	for i, entry := range data {
		printOne(entry, i)
	}
}

func printOne(entry Entry, idx int) {
	coloredText := fmt.Sprintf(
		"name: %s, surname: %s, tel: %s, last access: %s",
		green(entry.Name),
		green(entry.Surname),
		green(entry.Tel),
		green(entry.LastAccess),
	)
	if idx > 0 {
		fmt.Println(cyan(idx), coloredText)
	} else {
		fmt.Println(coloredText)
	}
}
