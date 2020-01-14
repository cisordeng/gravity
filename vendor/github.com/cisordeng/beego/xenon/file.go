package xenon

import (
	"io"
	"mime/multipart"
	"os"
)

func SaveToFile(file multipart.File, tofile string) {
	defer file.Close()
	f, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	PanicNotNilError(err, "rest:upload failed", "上传失败")
	defer f.Close()
	io.Copy(f, file)
}