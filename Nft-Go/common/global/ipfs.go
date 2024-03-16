package global

import (
	"bytes"
	"context"
	"github.com/dubbogo/gost/log/logger"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
)

var IpfsClient *Ipfs

type Ipfs struct {
	Url string
	Sh  *shell.Shell
}

func InitIpfs(url string) {
	sh := shell.NewShell(url)
	IpfsClient = &Ipfs{
		Url: url,
		Sh:  sh,
	}
	logger.Info("ipfs connect success")
}

func GetIpfs() Ipfs {
	return *IpfsClient
}

// UploadIPFS 上传数据到ipfs
func (i *Ipfs) UploadIPFS(data []byte) (hash string, err error) {
	hash, err = i.Sh.Add(bytes.NewReader(data), shell.Pin(true))
	if err != nil {
		return
	}
	return
}

// UnPinIPFS 从ipfs上删除数据
func (i *Ipfs) UnPinIPFS(hash string) (err error) {
	err = i.Sh.Unpin(hash)
	if err != nil {
		return
	}

	err = i.Sh.Request("repo/gc", hash).
		Option("recursive", true).
		Exec(context.Background(), nil)
	if err != nil {
		return
	}

	return nil
}

// CatIPFS 从ipfs下载数据
func (i *Ipfs) CatIPFS(hash string) (string, error) {
	read, err := i.Sh.Cat(hash)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(read)

	return string(body), nil
}
