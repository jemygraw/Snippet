package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/x/rpc.v7"
)

const (
	BLOCK_SIZE int64 = 1 << 22
)

const (
	EXIT_OK  = 0
	EXIT_ERR = 1
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

func FormatFsize(fsize int64) (result string) {
	if fsize > TB {
		result = fmt.Sprintf("%.2f TB", float64(fsize)/float64(TB))
	} else if fsize > GB {
		result = fmt.Sprintf("%.2f GB", float64(fsize)/float64(GB))
	} else if fsize > MB {
		result = fmt.Sprintf("%.2f MB", float64(fsize)/float64(MB))
	} else if fsize > KB {
		result = fmt.Sprintf("%.2f KB", float64(fsize)/float64(KB))
	} else {
		result = fmt.Sprintf("%d B", fsize)
	}

	return
}

type PutRet struct {
	Key      string `json:"key"`
	Hash     string `json:"hash"`
	Fsize    int64  `json:"fsize"`
	MimeType string `json:"mimeType"`
}

func main() {
	var bucket string
	var accessKey string
	var secretKey string
	var key string
	var localFile string
	var mimeType string
	var upHost string
	var fileType int
	var deleteAfterDays int
	var overwrite bool

	flag.StringVar(&bucket, "bucket", "", "bucket to save")
	flag.StringVar(&accessKey, "ak", "", "access key")
	flag.StringVar(&secretKey, "sk", "", "secret key")
	flag.StringVar(&key, "key", "", "file key")
	flag.StringVar(&localFile, "file", "", "local file path")
	flag.StringVar(&mimeType, "mime", "", "file mime type")
	flag.StringVar(&upHost, "host", "", "upload host")
	flag.IntVar(&fileType, "type", 0, "file storage type")
	flag.BoolVar(&overwrite, "overwrite", false, "overwrite existing file")
	flag.IntVar(&deleteAfterDays, "expire", 0, "delete after days")
	flag.Parse()

	if bucket == "" || accessKey == "" || secretKey == "" || localFile == "" {
		fmt.Fprintln(os.Stderr, "Usage: print -h to see help")
		os.Exit(EXIT_ERR)
	}

	fileInfo, statErr := os.Stat(localFile)
	if statErr != nil {
		fmt.Fprintf(os.Stderr, "stat local file error, %s\n", statErr.Error())
		os.Exit(EXIT_ERR)
	}

	config := &storage.Config{}
	var zone *storage.Zone
	if upHost != "" {
		//check up host
		upHostURI, pErr := url.Parse(upHost)
		if pErr != nil {
			fmt.Fprintf(os.Stderr, "invalid up host `%s` error", upHost)
			os.Exit(EXIT_ERR)
		}

		scheme := "http"
		if upHostURI.Scheme != "" {
			scheme = upHostURI.Scheme
		}

		useHTTPS := false
		if scheme == "https" {
			useHTTPS = true
		}

		upDomain := upHostURI.Host
		if upDomain != "" {
			zone = &storage.Zone{
				SrcUpHosts: []string{upDomain},
			}
			config.UseHTTPS = useHTTPS
			config.Zone = zone
		}
	}

	fileSize := fileInfo.Size()
	ctx := context.Background()
	putPolicy := storage.PutPolicy{}

	if overwrite && key != "" {
		putPolicy.Scope = fmt.Sprintf("%s:%s", bucket, key)
	} else {
		putPolicy.Scope = bucket
	}
	putPolicy.ReturnBody = `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"mimeType":"$(mimeType)"}`
	//set deadline
	putPolicy.Expires = 7 * 24 * 3600 //7days

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	putRet := PutRet{}

	startTime := time.Now()
	var putErr error
	if fileSize < BLOCK_SIZE {
		//form upload
		formUploader := storage.NewFormUploader(config)
		putExtra := storage.PutExtra{}
		if mimeType != "" {
			putExtra.MimeType = mimeType
		}
		if key != "" {
			putErr = formUploader.PutFile(ctx, &putRet, upToken, key, localFile, &putExtra)
		} else {
			putErr = formUploader.PutFileWithoutKey(ctx, &putRet, upToken, localFile, &putExtra)
		}
	} else {
		//resume upload
		resumeUploader := storage.NewResumeUploader(config)
		putExtra := storage.RputExtra{}
		if mimeType != "" {
			putExtra.MimeType = mimeType
		}
		if key != "" {
			putErr = resumeUploader.PutFile(ctx, &putRet, upToken, key, localFile, &putExtra)
		} else {
			putErr = resumeUploader.PutFileWithoutKey(ctx, &putRet, upToken, localFile, &putExtra)
		}
	}

	if putErr != nil {
		if v, ok := putErr.(*rpc.ErrorInfo); ok {
			fmt.Fprintf(os.Stderr, "Put file error, %d %s, Reqid: %s\n", v.Code, v.Err, v.Reqid)
		} else {
			fmt.Fprintln(os.Stderr, "Put file error,", putErr)
		}

		os.Exit(EXIT_ERR)
	} else {
		fmt.Println("Put file", localFile, "=>", bucket, ":", putRet.Key, "success!")
		fmt.Println("Hash:", putRet.Hash)
		fmt.Println("Fsize:", putRet.Fsize, "(", FormatFsize(putRet.Fsize), ")")
		fmt.Println("MimeType:", putRet.MimeType)

		lastNano := time.Now().UnixNano() - startTime.UnixNano()
		lastTime := fmt.Sprintf("%.2f", float32(lastNano)/1e9)
		avgSpeed := fmt.Sprintf("%.1f", float32(fileSize)*1e6/float32(lastNano))
		fmt.Println("Last time:", lastTime, "s, Average Speed:", avgSpeed, "KB/s")

		os.Exit(EXIT_OK)
	}
}
