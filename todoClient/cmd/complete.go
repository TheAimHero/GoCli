package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var completeCmd = &cobra.Command{
	Use:          "complete <id>",
	Short:        "Marks an item as completed",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")
		return completeAction(os.Stdout, apiRoot, args[0])
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func completeAction(out io.Writer, apiRoot, arg string) error {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return fmt.Errorf("%w: Item id must be a number", ErrNotNumber)
	}
	if err := completeItem(apiRoot, id); err != nil {
		return err
	}
	return printComplete(out, id)
}

func printComplete(out io.Writer, id int) error {
	_, err := fmt.Fprintf(out, "Item number %d marked as completed.\n", id)
	return err
}
