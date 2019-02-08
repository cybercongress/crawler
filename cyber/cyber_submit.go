package cyber

import (
	"fmt"
	"github.com/cybercongress/cyberd-wiki-index/ipfs"
	"github.com/cybercongress/cyberd-wiki-index/wiki"
	"github.com/cybercongress/cyberd/client"
	"github.com/cybercongress/cyberd/x/link/types"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
)

// todo add state save
func SubmitLinksToCyberCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "submit-links-to-cyber <path-to-wiki-titles-file>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			chunkSize := viper.GetInt("chunk")

			ipfsClient := ipfs.Open()
			wikiReader, err := wiki.OpenTitlesReader(args[0])
			if err != nil {
				return err
			}

			cbdClient := client.NewHttpCyberdClient(
				viper.GetString(client.FlagNode),
				viper.GetString(client.FlagPassphrase),
				viper.GetString(client.FlagAddress),
			)

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

				page := ipfs.RawContentHash(wiki.Dura(title))
				for _, keyword := range keywords {
					fromCid := ipfsClient.GetUnixfsContentHashWithRetryOnError(keyword)
					links = append(links, types.Link{From: types.Cid(fromCid), To: page})
					counter++
				}

				if len(links) >= chunkSize {
					log.Printf("%d %s\n", counter, title)
					printAccBandwidth(cbdClient)

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

	cmd.Flags().String(client.FlagAddress, "", "Address to sign transactions")
	cmd.Flags().String(client.FlagPassphrase, "", "Passphrase of account")
	cmd.Flags().String(client.FlagNode, "127.0.0.1:26657", "Url of node communicate with")
	cmd.Flags().String(client.FlagHome, homeDir+"/.cyberdcli", "Cyberd CLI home folder")
	cmd.Flags().Int("chunk", 1000, "How many links put into single transaction")

	_ = viper.BindPFlag(client.FlagPassphrase, cmd.Flags().Lookup(client.FlagPassphrase))
	_ = viper.BindPFlag(client.FlagAddress, cmd.Flags().Lookup(client.FlagAddress))
	_ = viper.BindPFlag(client.FlagNode, cmd.Flags().Lookup(client.FlagNode))
	_ = viper.BindPFlag(client.FlagHome, cmd.Flags().Lookup(client.FlagHome))
	_ = viper.BindPFlag("chunk", cmd.Flags().Lookup("chunk"))

	return &cmd
}

func printAccBandwidth(cbdClient client.CyberdClient) {
	accBw, err := cbdClient.GetAccountBandwidth()
	if err == nil {
		per := int64(100 * float64(accBw.RemainedValue) / float64(accBw.MaxValue))
		log.Printf("Remaining acc bw: %d %v%%\n", accBw.RemainedValue, per)
	}
}
