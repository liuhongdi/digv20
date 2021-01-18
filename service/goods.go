package service

import (
	"github.com/liuhongdi/digv20/dao"
	"github.com/liuhongdi/digv20/model"
)

//得到多篇文章，按分页返回
func GetGoodsList() ([]*model.Goods,error) {
	goods, err := dao.SelectAllGoods()
    return goods,err
}
