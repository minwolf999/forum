package posts

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func FormatImg(Path multipart.File, handler *multipart.FileHeader) (ImagePath string) {
	//Create the path who is stored the image
	dst, err := os.Create(filepath.Join("./front/static/imgs/", filepath.Base(handler.Filename)))
	if err != nil {
		log.Println(err)
	}
	defer dst.Close()

	//Copy the image in the folder
	if _, err := io.Copy(dst, Path); err != nil {
		log.Println("Err Copy")
	}
	//Stock the path in the var
	ImagePath = ("/front/static/imgs/" + handler.Filename)
	//This value is returned and sending after in the database
	return ImagePath
}
