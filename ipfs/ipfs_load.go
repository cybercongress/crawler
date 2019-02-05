package ipfs

import (
	"github.com/spf13/cobra"
)

/*
   Command should make chunks of 100k titles, and than final recursive batch of merges by mv commands.

	()   () ()   ()
     \   /   \   /
      ( )     ( )
        \     /
        (     )

*/
// Command should persist progress(state) to deal with restarts
func UploadDurasToIpfsCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "upload-duras-to-ipfs",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	return &cmd
}
