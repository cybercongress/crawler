package main

import (
	"fmt"
	"github.com/cybercongress/cyberd-wiki-index/cyber"
	"github.com/cybercongress/cyberd-wiki-index/ipfs"
	"github.com/cybercongress/cyberd-wiki-index/state"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	homeDir, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	indexationState, err := state.Load(homeDir + "/.cyber-wiki")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "cbr-wiki",
		Short: "Index Wiki to the cyber protocol",
	}

	cmd.AddCommand(cyber.SubmitLinksToCyberCmd(indexationState))
	cmd.AddCommand(ipfs.UploadDurasToIpfsCmd())

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
