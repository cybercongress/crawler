package cyber

import (
	"fmt"
	"github.com/cybercongress/cyberd-wiki-index/reader"
	"github.com/cybercongress/cyberd/client"
	"github.com/cybercongress/cyberd/x/link/types"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
	"time"
)

// Command should persist progress(state) to deal with restarts
// Add option to push to ipfs data with submitting
// Add option to skip link, if already created by others
func SubmitLinksToCyberCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "submit-links-to-cyber",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			chunkSize := 1000

			wikiReader, err := reader.Open(args[0])
			cbdClient := client.NewHttpCyberdClient(
				viper.GetString(client.FlagNode),
				viper.GetString(client.FlagPassphrase),
				viper.GetString(client.FlagAddress),
			)
			if err != nil {
				return err
			}

			counter := int64(0)
			links := make([]types.Link, 0, chunkSize)
			for {

				title, keywords, err := wikiReader.NextTitleWithKeywords()
				if err != nil {
					if err == io.EOF {
						break
					}
					return err
				}

				page := title + ".wiki"
				for _, keyword := range keywords {
					links = append(links, types.Link{From: reader.Cid(keyword), To: reader.Cid(page)})
					counter++
				}

				if len(links) >= chunkSize {
					fmt.Printf("%d %s\n", counter, title)
					printAccBandwidth(cbdClient)
					time.Sleep(10 * time.Second)

					err = cbdClient.SubmitLinksSync(links)
					if err != nil {
						return err
					}

					links = make([]types.Link, 0, chunkSize)
				}

			}
			return nil
		},
	}

	homeDir, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Flags().String(client.FlagAddress, "", "Account to sign transactions")
	cmd.Flags().String(client.FlagPassphrase, "", "Passphrase of account")
	cmd.Flags().String(client.FlagNode, "127.0.0.1:26657", "Url of node communicate with")
	cmd.Flags().String(client.FlagHome, homeDir+"/.cyberdcli", "Cyberd CLI home folder")

	_ = viper.BindPFlag(client.FlagPassphrase, cmd.Flags().Lookup(client.FlagPassphrase))
	_ = viper.BindPFlag(client.FlagAddress, cmd.Flags().Lookup(client.FlagAddress))
	_ = viper.BindPFlag(client.FlagNode, cmd.Flags().Lookup(client.FlagNode))
	_ = viper.BindPFlag(client.FlagHome, cmd.Flags().Lookup(client.FlagHome))

	return &cmd
}

func printAccBandwidth(cbdClient client.CyberdClient) {
	accBw, err := cbdClient.GetAccountBandwidth()
	if err == nil {
		per := int64(100 * float64(accBw.RemainedValue) / float64(accBw.MaxValue))
		fmt.Printf("Remaining acc bw: %d %v%%\n", accBw.RemainedValue, per)
	}
}
