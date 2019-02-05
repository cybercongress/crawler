package cyber

import (
	"github.com/spf13/cobra"
)

// Command should persist progress(state) to deal with restarts
// Add option to push to ipfs data with submitting
// Add option to skip link, if already created by others
func SubmitLinksToCyberCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "submit-links-to-cyber",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	return &cmd
}
