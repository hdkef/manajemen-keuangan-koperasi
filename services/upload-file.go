package services

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ABS_DST string                            //filepath absolute from system A.K.A /usr/local/share/bla/bla/bla
var RELATIVE_DST string = os.Getenv("DOCDST") //filepath relative to url domain A.K.A public/doc

func init() {
	_ = godotenv.Load()
	//first create folder to save document and get destination folder
	ABS_DST = createFolderAndGetDocDST()
}

//upload file and return filepath relative
func UploadFile(c *gin.Context, formname string, uniquevalue string) (string, error) {
	file, err := c.FormFile(formname)
	if err != nil {
		return "", err
	}

	//filename is Time.Now() with uniquevalue

	fname := time.Now().String() + uniquevalue

	fpath := filepath.Join(ABS_DST, fname)

	err = c.SaveUploadedFile(file, fpath)
	if err != nil {
		return "", err
	}
	return filepath.Join(RELATIVE_DST, fname), nil
}

func createFolderAndGetDocDST() string {

	absPWD, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dst := filepath.Join(absPWD, RELATIVE_DST)

	err = os.MkdirAll(dst, 0755)
	if err != nil {
		panic(err)
	}
	return dst
}

func RemoveFile(relativefpath string) error {

	absPWD, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dst := filepath.Join(absPWD, relativefpath)

	return os.Remove(dst)
}
