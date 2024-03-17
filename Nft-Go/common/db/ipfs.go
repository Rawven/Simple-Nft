package db

import (
	"bytes"
	"context"
	"github.com/dubbogo/gost/log/logger"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
	"os"
)

var IpfsClient *Ipfs
var gateWay string

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
	//TODO gateway是啥？
	gateWay = url
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

func (i *Ipfs) UploadIPFSByPath(filePath string) (cid string, err error) {
	// 创建一个shell
	sh := shell.NewShell("localhost:5001")

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 上传文件
	cid, err = sh.Add(file)
	if err != nil {
		return "", err
	}

	return cid, nil
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

func (i *Ipfs) GetFileUrl(hash string) string {
	return gateWay + "/ipfs/" + hash
}
