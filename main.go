package main

import (
	"fmt"
	"github.com/cybercongress/cyberd-wiki-index/cyber"
	"github.com/cybercongress/cyberd-wiki-index/ipfs"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	cmd := &cobra.Command{
		Use:   "cbr-wiki",
		Short: "Index Wiki to the cyber protocol",
	}

	cmd.AddCommand(cyber.SubmitLinksToCyberCmd())
	cmd.AddCommand(ipfs.UploadDurasToIpfsCmd())

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
