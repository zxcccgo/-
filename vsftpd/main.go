package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"encoding/hex"

	"gopkg.in/dutchcoders/goftp.v1"
)

func main() {
    var err error
    var ftp *goftp.FTP

    // For debug messages: goftp.ConnectDbg("ftp.server.com:21")
    if ftp, err = goftp.Connect("127.0.0.1:21"); err != nil {
        panic(err)
    }

    defer ftp.Close()
    fmt.Println("Successfully connected to", "127.0.0.1")

    // TLS client authentication
    // config := tls.Config{
    //     InsecureSkipVerify: true,
    //     ClientAuth:         tls.RequestClientCert,
    // }

    // if err = ftp.AuthTLS(&config); err != nil {
	
    //     panic(err)
    // }

    // Username / password authentication
    if err = ftp.Login("zx", "123456"); err != nil {
        panic(err)
    }

    if err = ftp.Cwd("/"); err != nil {
        panic(err)
    }

    var curpath string
    if curpath, err = ftp.Pwd(); err != nil {
        panic(err)
    }

    fmt.Printf("Current path: %s", curpath)

    // Get directory listing
    var files []string
    if files, err = ftp.List(""); err != nil {
        panic(err)
    }
    fmt.Println("Directory listing:", files)

    // Upload a file
    var file *os.File
    if file, err = os.Open("./ftp.txt"); err != nil {
        panic(err)
    }

    if err := ftp.Stor("/zx", file); err != nil {
		log.Println("68")
        panic(err)
    }

    // Download each file into local memory, and calculate it's sha256 hash
  
        _, err = ftp.Retr("hello", func(r io.Reader) error {
            var hasher = sha256.New()
            if _, err = io.Copy(hasher, r); err != nil {
                return err
            }

            hash := fmt.Sprintf("%s %x", "hello", hex.EncodeToString(hasher.Sum(nil)))
            fmt.Println(hash)

            return err
        })

       
    
}

