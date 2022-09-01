// Application : Cmd Summary

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	coreUrl "webfetcher/core/url"
)

func init() {
	rootCmd.AddCommand(cobraCmd)
}

var cobraCmd = &cobra.Command{
	Use:   "summary",
	Short: "Print a summary from a web",
	Long:  `Print a summary from a web`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("invalid URL specified")
			return
		}

		url := args[0]

		// Create Url Instance
		u, err := coreUrl.NewUrl(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Create UrlService Instance
		us := coreUrl.NewUrlService(u)

		// Fetch summary of page
		re, err := us.FetchSummary()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Show Summary
		fmt.Println(re)
	},
}