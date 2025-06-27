package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "Command for call list of elements",
	Long: `Command for call list of elements`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile("tasks.json")
		if err != nil{
			fmt.Println("Error", err)
			return
		}

		var tasks []Task
		json.Unmarshal(data, &tasks)

		a, b := tableLength(tasks)
		c := createTable(a, b)

		fmt.Println(c)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func tableLength(tasks []Task) (int, int){
	tableLength := 40
	l := len(tasks)
	maxId := tasks[l-1].Id

	for i := 0;	i < len(tasks)-1; i++{
		length := len(tasks[i].Task)

		if tableLength < length{
			tableLength = length
		}
	}

	return maxId, tableLength
}

func createTable(mxid, tblen int) string{
	var(
		corner = '+'
		plafond = '-'
	)
	runes1 := make([]rune, 0)
	runes2 := make([]rune, 0)
	runes3 := make([]rune, 0)

	for i := 0; i < 10; i++{
		runes3 = append(runes3, plafond)
		if i == 9 {
			runes3 = append(runes3, corner)
		}
	}

	for i := -1; i <= tblen; i++{
		if i == 0 && mxid < 10{
			runes1 = append(runes1, corner, plafond, plafond, plafond, plafond, plafond, corner)
		} else if i == 0 && mxid < 100{
			runes1 = append(runes1, corner, plafond, plafond, plafond, plafond, plafond, plafond, corner)
		} else if i == 0 && mxid < 1000{
			runes1 = append(runes1, corner, plafond, plafond, plafond, plafond, plafond, plafond, plafond, corner)
		} else if i < tblen{
			runes2 = append(runes2, '-')
		} else {
			runes2 = append(runes2, '+')
		}
	}

	return string(runes1) + string(runes2) + string(runes3)
}