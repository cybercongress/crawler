package ipfs

import (
	links "github.com/cybercongress/cyberd/x/link"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

var pref = cid.Prefix{
	Version:  1,
	Codec:    cid.Raw,
	MhType:   multihash.SHA2_256,
	MhLength: -1, // default length
}

func RawContentHash(data string) links.Cid {
	result, _ := pref.Sum([]byte(data))
	return links.Cid(result.String())
}
