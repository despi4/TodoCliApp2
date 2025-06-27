package cmd

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/mergestat/timediff"
	"time"
)

const taskFile = "tasks.json"

type Task struct{
	Id int `json:"id"`
	Task string `json:"task"`
	State string `json:"state"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long: `Add command help you add new task.
	For example: todo add [task]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Println("Write Task, for example:todo add Make Homework")
			return
		}

		tasks := taskload()
		var str string

		for i := 0; i < len(args); i++{
			if i == 0{
				str += args[i]
			} else {
				str += " " + args[i]
			}
		}

		newTask := Task{
			Id: len(tasks)+1,
			Task: str,
			State: "Undone",
		}

		tasks = append(tasks, newTask)

		SaveTasks(tasks)

		fmt.Printf("Task added: [%d] %s\n", newTask.Id, newTask.Task)
		
		fmt.Println(timediff.TimeDiff(time.Now().Add(time.Hour)))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func taskload() []Task{
	var tasks []Task
	
	if _, err := os.Stat(taskFile); err == nil{
		data, err := os.ReadFile(taskFile)
		if err == nil{
			json.Unmarshal(data, &tasks)
		}
	}

	return tasks
}

func SaveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "	")
	if err != nil{
		fmt.Println("Error", err)
		return
	}

	err = os.WriteFile(taskFile, data, 0644)
	if err != nil{
		fmt.Println("Error", err)
	}
}