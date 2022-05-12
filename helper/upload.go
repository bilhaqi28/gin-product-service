package helper

import (
	"io"
	"net/http"
	"os"
)

func UploadFile(request *http.Request, field string, path string) string {
	file, header, err := request.FormFile(field)
	PanicIfError(err)
	filename := header.Filename
	out, err := os.Create("public/" + path + "/" + filename)
	defer out.Close()
	PanicIfError(err)
	_, err = io.Copy(out, file)
	PanicIfError(err)
	return filename
}
