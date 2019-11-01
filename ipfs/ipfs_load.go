package ipfs

import (
	"github.com/cybercongress/crawler/wiki"
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
			wikiReader, err := wiki.OpenTitlesReader(args[0])
			if err != nil {
				return err
			}

			dirPath := "/wiki-duras/"
			err = ipfs.CreateDir(dirPath)
			if err != nil {
				return err
			}

			hashCount := 0
			for {

				log.Println("Loading duras")
				duras := make([]string, 0, chunkSize)
				titles, err, hasMore := wikiReader.NextTitles(chunkSize)

				if err != nil {
					return err
				}

				for _, title := range titles {
					duras = append(duras, wiki.Dura(title))
				}

				log.Println("Creating virtual dirs")
				rootDir := getAsDirHierarchy(duras, 500)
				mfReader := files.NewMultiFileReader(rootDir, true)

				log.Println("Uploading dir to ipfs")
				hash := ipfs.AddDirectoryWithRetryOnError(mfReader)

				log.Println("Produced hash: " + hash)
				ipfs.AddFileToDirWithRetryOnError(hash, dirPath+strconv.Itoa(hashCount))

				if hasMore == false {
					break
				}
				hashCount++
			}

			stat, err := ipfs.DirStat(dirPath)
			log.Println("Final stats:\n" + stat)
			return err
		},
	}
	return &cmd
}
