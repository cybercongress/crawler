package ipfs

import (
	"context"
	"encoding/json"
	"github.com/ipfs/go-ipfs-api"
	"github.com/ipfs/go-ipfs-files"
	"io/ioutil"
	"strings"
)

type addResponse struct {
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

// returns root hash of added dir/file
func (c Client) AddMultiFile(mfr *files.MultiFileReader) (string, error) {

	resp, err := c.ipfs.
		Request("add").
		Body(mfr).
		Option("recursive", true).
		//Option("only-hash", true).
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

func (c Client) CreateDir(path string) error {
	return c.ipfs.
		Request("files/mkdir").
		Arguments(path).
		Option("parents", true).
		Exec(context.Background(), nil)
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
