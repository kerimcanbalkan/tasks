package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete task with given ID",
	Long:  `Deletes the task given id by the arguments. Example usage tasks delete <ID>`,
	Run:   delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete(cmd *cobra.Command, args []string) {
	id := args[0]
	var keepers [][]string
	taskDeleted := false

	file, err := os.Open("output.csv")
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("There are no tasks added!")
		return
	} else if err != nil {
		log.Fatal("Couldn't open the file", err)
	}

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal("Couldn't read the csv")
	}

	for _, line := range lines {
		if line[0] != id {
			keepers = append(keepers, line)
		} else {
			taskDeleted = true
		}
	}

	if !taskDeleted {
		fmt.Printf("Task with given id %s not found", id)
		return
	}

	f, err := os.Create("output.csv")
	if err != nil {
		log.Fatal("Could not create the file", err)
	}
	writer := csv.NewWriter(f)

	for _, el := range keepers {
		if err := writer.Write(el); err != nil {
			log.Fatal("Failed to write record:", err)
		}

	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal("Failed to flush writer:", err)
	}

}
