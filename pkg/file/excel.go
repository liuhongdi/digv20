package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func FileExcelDown(c *gin.Context,filepath string,filename string) {
	//filename := "php_errors.log"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/msexcel")
	c.File(filepath)
}
