package helper

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func UploadFile(request *http.Request, field string, path string) (string, error) {
	var filename string
	file, header, err := request.FormFile(field)
	if err != nil {
		filename = "Empty"
		return filename, errors.New("Tidak Ada Gambar")
	} else {
		filename = header.Filename
		out, err := os.Create("public/" + path + "/" + filename)
		defer out.Close()
		PanicIfError(err)
		_, err = io.Copy(out, file)
		PanicIfError(err)
		return filename, nil
	}
}
