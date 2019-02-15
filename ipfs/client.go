package ipfs

import (
	"context"
	"encoding/json"
	"github.com/cybercongress/cyberd-wiki-index/util"
	"github.com/ipfs/go-ipfs-api"
	"github.com/ipfs/go-ipfs-files"
	"io/ioutil"
	"strings"
)

type addResponse struct {
	Hash string
}

type lsResponse struct {
	Entries []lsEntryResponse
}

type lsEntryResponse struct {
	Name string
	Type int
	Size int
	Hash string
}

type Client struct {
	ipfs *shell.Shell
}

func Open() Client {
	return Client{
		ipfs: shell.NewShell("localhost:5001"),
	}
}

func (c Client) AddDirectoryWithRetryOnError(mfr *files.MultiFileReader) string {
	var err error
	var dirHash string
	util.RetryUntilOk(func() error {
		dirHash, err = c.AddDirectory(mfr)
		return err
	}, "Error uploading dir to ipfs")
	return dirHash
}

// returns root hash of added dir/file
func (c Client) AddDirectory(mfr *files.MultiFileReader) (string, error) {

	resp, err := c.ipfs.
		Request("add").
		Body(mfr).
		Option("recursive", true).
		Option("raw-leaves", true).
		Option("wrap-with-directory", true).
		Send(context.Background())

	if err != nil {
		return "", err
	}

	out, err := ioutil.ReadAll(resp.Output)
	if err != nil {
		return "", err
	}

	return getRootHash(out)
}

func (c Client) GetUnixfsContentHashWithRetryOnError(content string) string {
	var err error
	var fromCid string
	util.RetryUntilOk(func() error {
		fromCid, err = c.UnixfsContentHash(content)
		return err
	}, "Error obtains keyword hash")
	return fromCid
}

func (c Client) UnixfsContentHash(content string) (string, error) {

	dir := ConstructDir([]string{content})
	mfr := files.NewMultiFileReader(dir, true)
	resp, err := c.ipfs.
		Request("add").
		Body(mfr).
		Option("only-hash", true).
		Option("wrap-with-directory", false).
		Send(context.Background())

	if err != nil {
		return "", err
	}

	out, err := ioutil.ReadAll(resp.Output)
	if err != nil {
		return "", err
	}
	return getRootHash(out)
}

func (c Client) AddFiles(mfr *files.MultiFileReader) error {

	return c.ipfs.
		Request("add").
		Body(mfr).
		Option("recursive", true).
		Option("wrap-with-directory", true).
		Option("pin", true).
		Exec(context.Background(), nil)
}

func (c Client) CreateDir(path string) error {
	return c.ipfs.
		Request("files/mkdir").
		Arguments(path).
		Option("parents", true).
		Exec(context.Background(), nil)
}

func (c Client) GetDirEntriesName(path string) ([]string, error) {

	var entries lsResponse
	var entriesNames = make([]string, 0)

	err := c.ipfs.
		Request("files/ls").
		Arguments(path).
		Option("parents", true).
		Exec(context.Background(), entries)

	if err != nil {
		for _, entry := range entries.Entries {
			entriesNames = append(entriesNames, entry.Name)
		}
	}

	return entriesNames, err
}

func (c Client) AddFileToDirWithRetryOnError(fileHash string, dirPath string) {
	util.RetryUntilOk(func() error {
		err := c.AddFileToDir(fileHash, dirPath)
		return err
	}, "Error adding file '"+fileHash+"'to dir '"+dirPath+"' to ipfs")
}

func (c Client) AddFileToDir(fileHash string, dirPath string) error {
	return c.ipfs.
		Request("files/cp").
		Arguments("/ipfs/"+fileHash, dirPath).
		Exec(context.Background(), nil)
}

func (c Client) DirStat(path string) (string, error) {
	resp, err := c.ipfs.
		Request("files/stat").
		Arguments(path).
		Option("parents", true).
		Send(context.Background())

	out, err := ioutil.ReadAll(resp.Output)
	return string(out), err
}

func getRootHash(resp []byte) (string, error) {

	//add command returns multiple json object separated by '\n' and empty last line
	//the last one entry is
	//{"foo": "bar"}
	//{"foo": "baz"}
	//
	hashes := strings.Split(string(resp), "\n")
	rootHashObj := hashes[len(hashes)-2]

	var addResp addResponse
	decoder := json.NewDecoder(strings.NewReader(rootHashObj))
	err := decoder.Decode(&addResp)
	return addResp.Hash, err
}
