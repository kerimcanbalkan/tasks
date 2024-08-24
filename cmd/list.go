package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the all tasks that have been added",
	Long:  `list lists the tasks in table format.It does not take any arguments, Example usage, tasks list`,
	Run:   list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
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
	writer := tabwriter.NewWriter(
		os.Stdout, 0, 2, 4, ' ', 0,
	)

	if _, err := writer.Write([]byte("ID\tDescription\tCreatedAt\tIsComplete\n")); err != nil {
		log.Fatal("Tabwriter had an error", err)
	}
	for _, line := range lines {
		if line[0] == "ID" {
			continue
		}
		parsedTime, err := time.Parse(time.RFC3339, line[2])
		if err != nil {
			log.Fatal("Failed to parse time:", err)
		}
		timeDifference := timediff.TimeDiff(parsedTime)
		formattedString := fmt.Sprintf("%s\t%s\t%s\t%s\n", line[0], line[1], timeDifference, line[3])
		if _, err := writer.Write([]byte(formattedString)); err != nil {
			log.Fatal("Tabwriter had an error", err)
		}

	}
	writer.Flush()
	defer file.Close()
}
