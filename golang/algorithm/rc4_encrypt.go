package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const RC4_MAX_SIZE = 256
const ENCRYPT_READ_BUFFER = 5 * 1024 * 1024

type RC4Encrypt struct {
	sbox [RC4_MAX_SIZE]byte
	key  [RC4_MAX_SIZE]byte
}

func (this *RC4Encrypt) InitKey(encryptKey []byte) {
	eKey := encryptKey
	eKeyLen := len(encryptKey)

	//init the sequence
	for m := 0; m < RC4_MAX_SIZE; m++ {
		this.key[m] = eKey[m%eKeyLen]
		this.sbox[m] = byte(m)
	}

	n := 0
	for m := 0; m < RC4_MAX_SIZE; m++ {
		n = (n + int(this.sbox[m]) + int(this.key[m])) & 0xff

		this.sbox[m] ^= this.sbox[n]
		this.sbox[n] ^= this.sbox[m]
		this.sbox[m] ^= this.sbox[n]
	}
}

func (this *RC4Encrypt) DoEncrypt(dataToEncrypt []byte, dataToEncryptLen int, offsetInFile int64) (encryptedData []byte) {
	encryptedData = make([]byte, 0, dataToEncryptLen)

	dataMaxOffset := offsetInFile + int64(dataToEncryptLen)

	for i := offsetInFile; i < dataMaxOffset; i++ {
		var h int = (int(i) + 8) & 0xff
		var j int = (h + int(this.key[h]) + int(this.sbox[h])) & 0xff
		var k = this.sbox[(this.sbox[h]+this.sbox[j])&0xff]

		newByte := dataToEncrypt[i-offsetInFile] ^ k
		encryptedData = append(encryptedData, newByte)
	}
	return
}

func main() {
	var encryptKey string
	var srcFile string
	var destFile string

	flag.StringVar(&srcFile, "src", "", "src file to encrypt/decrypt")
	flag.StringVar(&destFile, "dest", "", "encrypt/decrypt dest file")
	flag.StringVar(&encryptKey, "key", "", "encrypt/decrypt key")

	flag.Usage = func() {
		fmt.Println(`Usage of ./encrypt:
  -src="": src file to encrypt/decrypt
  -dest="": encrypted/decrypted dest file
  -key="": encrypt/decrypt key`)
	}

	flag.Parse()

	//check params
	if srcFile == "" {
		fmt.Println("no src file specified")
		return
	}

	if destFile == "" {
		fmt.Println("not dest file specified")
		return
	}

	if encryptKey == "" {
		fmt.Println("no key specified")
		return
	}

	_, err := os.Stat(srcFile)
	if os.IsNotExist(err) {
		fmt.Println("src file specified not exist")
		return
	}

	if err := encrypt(srcFile, destFile, encryptKey); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done!")
	}

}

func encrypt(srcFile, destFile, encryptKey string) (err error) {
	srcFp, openErr := os.Open(srcFile)
	if openErr != nil {
		err = errors.New(fmt.Sprintf("open src file error, %s", openErr.Error()))
		return
	}
	defer srcFp.Close()

	destFp, openErr := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if openErr != nil {
		err = errors.New(fmt.Sprintf("open dest file error, %s", openErr.Error()))
		return
	}
	defer destFp.Close()

	buffer := make([]byte, ENCRYPT_READ_BUFFER)

	bReader := bufio.NewReader(srcFp)
	bWriter := bufio.NewWriter(destFp)

	var offset int64 = 0
	rc4 := RC4Encrypt{}
	rc4.InitKey([]byte(encryptKey))

	for {

		rCnt, rErr := bReader.Read(buffer)

		if rErr == io.EOF {
			break
		} else {
			//encrypt
			encryptedData := rc4.DoEncrypt(buffer, rCnt, offset)

			_, wErr := bWriter.Write(encryptedData)
			if wErr != nil {
				err = errors.New(fmt.Sprintf("write to dest file error, %s", wErr.Error()))
				return
			}

			//update offset
			offset += int64(rCnt)
		}
	}

	fErr := bWriter.Flush()
	if fErr != nil {
		err = errors.New(fmt.Sprintf("flush to dest file error, %s", fErr.Error()))
		return
	}

	return
}
