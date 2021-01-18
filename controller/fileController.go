package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv20/pkg/file"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/liuhongdi/digv20/service"
	"strconv"
)

type FileController struct{}

func NewFileController() FileController {
	return FileController{}
}

//下载一个txt文件
func (a *FileController) DownTxt(c *gin.Context) {
	filepath:="/data/logs/phplogs/php_errors.log"
	if (file.FileExist(filepath) == true) {
		fmt.Println("file exist")
	} else {
		fmt.Println("file not exist")
	}

	filename := "php_errors.log"
	file.FileDown(c,filepath,filename)
}

//下载一个txt文件
func (a *FileController) DownExcel(c *gin.Context) {

	//创建新excel文件
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	//new sheet
	 f.NewSheet("Sheet3")

	//设置第1行行高
	err_h := f.SetRowHeight("Sheet3", 1, 80)
	if (err_h != nil) {
		fmt.Println(err_h)
	}
	//设置第2行行高
	err_h2 := f.SetRowHeight("Sheet3", 2, 40)
	if (err_h2 != nil) {
		fmt.Println(err_h2)
	}
	//合并单元格A1到E1
	err3 := f.MergeCell("Sheet3", "A1", "E1")
    if (err3 != nil) {
		fmt.Println(err3)
	}

	//设置B列列宽column width
	errwb := f.SetColWidth("Sheet3", "B", "B", 30)
	if (errwb != nil) {
		fmt.Println(errwb)
	}
	//设置C列列宽
	errwc := f.SetColWidth("Sheet3", "C", "C", 40)
	if (errwc != nil) {
		fmt.Println(errwc)
	}
	//设置标题的样式
	f.SetCellValue("Sheet3", "A1", "商品信息表")
	style, err := f.NewStyle(`{
   "fill":{"type":"pattern","color":["#E0EBF5"],"pattern":1},
    "font":
    {
        "bold": true,
        "italic": false,
        "family": "Times New Roman",
        "size": 36,
        "color": "#777777"
    },
"alignment":
    {
        "horizontal": "center",
        "ident": 1,
        "justify_last_line": true,
        "reading_order": 0,
        "relative_indent": 1,
        "shrink_to_fit": true,
        "vertical": "center",
        "wrap_text": true
    }
}`)

	if err != nil {
		fmt.Println(err)
	}
	errStyle := f.SetCellStyle("Sheet3", "A1", "E1", style)
	if errStyle != nil {
		fmt.Println(errStyle)
	}

	//为表头设置样式,第2行
	style2, errst := f.NewStyle(`{
    "font":
    {
        "bold": false,
        "italic": false,
        "family": "Times New Roman",
        "size": 12,
        "color": "#0000ff"
    },
"alignment":
    {
        "horizontal": "center",
        "ident": 1,
        "justify_last_line": true,
        "reading_order": 0,
        "relative_indent": 1,
        "shrink_to_fit": true,
        "vertical": "center",
        "wrap_text": true
    }
}`)
	if errst != nil {
		fmt.Println(errst)
	}
	errStyle2 := f.SetCellStyle("Sheet3", "A2", "E2", style2)
	if errStyle2 != nil {
		fmt.Println(errStyle2)
	}

    //设置表头的值
	f.SetCellValue("Sheet3", "A2", "id")
	f.SetCellValue("Sheet3", "B2", "商品名称")
	f.SetCellValue("Sheet3", "C2", "商品描述")
	f.SetCellValue("Sheet3", "D2", "商品价格")
	f.SetCellValue("Sheet3", "E2", "库存数量")
	//写入数据
	//得到商品列表
	goods,err := service.GetGoodsList()
    fmt.Println(goods)

    //遍历商品
    for i,v := range goods {
		fmt.Println(i)
    	fmt.Println(v)
		curIndex := i+3
		strIndex := strconv.Itoa(curIndex)
		f.SetCellValue("Sheet3", "A"+strIndex, v.GoodsId)
		f.SetCellValue("Sheet3", "B"+strIndex, v.GoodsName)
		f.SetCellValue("Sheet3", "C"+strIndex, v.Subject)
		f.SetCellValue("Sheet3", "D"+strIndex, v.Price)
		f.SetCellValue("Sheet3", "E"+strIndex, v.Stock)
	}

	//保存成文件
	filepath:="/data/temp/Book1.xlsx"
	filename:="Book1.xlsx"
	err5 := f.SaveAs(filepath);
	// Save spreadsheet by the given path.
	if  err5 != nil {
		fmt.Println(err5)
	} else {
		//下载
		file.FileExcelDown(c,filepath,filename)
	}

}

