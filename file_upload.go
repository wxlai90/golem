package golem

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type FileUploadConfig struct {
	Filename    string
	Destination string
}

func FileUpload(config FileUploadConfig) Middleware {
	return func(req *Request, res *Response, next Next) {
		file, fileHeaders, err := req.FormFile(config.Filename)
		if err != nil {
			switch err {
			case http.ErrNotMultipart:
				next()
				return
			case http.ErrMissingFile:
				next()
				return
			default:
				log.Printf("err: %s\n", err)
				panic(err)
			}
		}
		defer file.Close()

		destinationName := fileHeaders.Filename
		fullPath := fmt.Sprintf("%s/%s", config.Destination, destinationName)
		dest, err := os.Create(fullPath)
		if err != nil {
			panic(err)
		}
		defer dest.Close()

		_, err = io.Copy(dest, file)
		if err != nil {
			panic(err)
		}

		next()
	}
}
