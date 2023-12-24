package cmd

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:          "list",
	Short:        "List todo items",
	SilenceUsage: true,
	RunE: func(_ *cobra.Command, _ []string) error {
		apiRoot := viper.GetString("api-root")
		return listAction(os.Stdout, apiRoot)
	},
}

func listAction(out io.Writer, apiRoot string) error {
	items, err := getAll(apiRoot)
	if err != nil {
		return err
	}
	return printAll(out, items)
}

func printAll(out io.Writer, items []item) error {
	w := tabwriter.NewWriter(out, 3, 2, 0, ' ', 0)
	for k, v := range items {
		done := "-"
		if v.Done {
			done = "X"
		}
		fmt.Fprintf(w, "%s\t%d\t%s\t\n", done, k+1, v.Task)
	}
	return w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
