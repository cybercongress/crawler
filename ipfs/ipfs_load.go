package ipfs

import (
	"github.com/cybercongress/cyberd-wiki-index/reader"
	"github.com/ipfs/go-ipfs-files"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func UploadDurasToIpfsCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "upload-duras-to-ipfs <path-to-wiki-titles-file>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			chunkSize := 100000

			ipfs := Open()
			wikiReader, err := reader.Open(args[0])
			if err != nil {
				return err
			}

			hashes := []string{}
			for {

				log.Println("Loading duras")
				duras := make([]string, 0, chunkSize)
				titles, err, hasMore := wikiReader.NextTitles(chunkSize)

				if err != nil {
					return err
				}

				for _, title := range titles {
					duras = append(duras, title+".wiki")
				}

				log.Println("Creating virtual dirs")
				rootDir := getAsDirHierarchy(duras, 500)
				mfReader := files.NewMultiFileReader(rootDir, true)

				log.Println("Sending request")
				hash, err := ipfs.AddDirectory(mfReader)
				if err != nil {
					return err
				}

				log.Println("Getting hash: " + hash)
				hashes = append(hashes, hash)
				if hasMore == false {
					break
				}
			}

			dirPath := "/wiki-duras/"
			err = ipfs.CreateDir(dirPath)
			if err != nil {
				return err
			}

			for i, hash := range hashes {
				err = ipfs.AddFileToDir(hash, dirPath+strconv.Itoa(i))
				if err != nil {
					return err
				}
			}

			stat, err := ipfs.DirStat(dirPath)
			log.Println("Final stats:\n" + stat)
			return err
		},
	}
	return &cmd
}
