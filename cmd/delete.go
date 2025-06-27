package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete not needed task",
	Long: `Delete not needed task
	For example: todo delete [id]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Println("Write id task which you want to delete: todo delete 2")
			return 
		}
		
		var tasks []Task

		id, err := strconv.Atoi(args[0])
		if err != nil{
			fmt.Println("Write id number")
			return
		}

		data, err := os.ReadFile("tasks.json")
		if err != nil{
			fmt.Println("Error when read the file", err)
			return
		}

		err = json.Unmarshal(data, &tasks)
		if err != nil{
			fmt.Println("Error", err)
			return
		}
		
		if len(tasks) >= id{
			id = id - 1
			tasks = append(tasks[:id], tasks[id+1:]...)
			tasks = changeId(tasks, id)
		}else{
			fmt.Println("The list doesn not have a task with this ID")
			return
		}

		SaveTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func changeId(tasks []Task, id int,) []Task{
	for i := id; i < len(tasks); i++{
		tasks[i]=Task{
			Id: i+1,
			Task: tasks[i].Task,
			State: tasks[i].State,
		}
	}

	return tasks
}
