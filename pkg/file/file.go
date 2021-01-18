package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func FileDown(c *gin.Context,filepath string,filename string) {
	//filename := "php_errors.log"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filepath)
}
