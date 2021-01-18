package dao

import (
	"github.com/liuhongdi/digv20/global"
	"github.com/liuhongdi/digv20/model"
)

func SelectAllGoods() ([]*model.Goods, error) {
	//fields := []string{"articleId", "subject", "url"}

	var goods []*model.Goods

	global.DBLink.Find(&goods)

	/*
	if err != nil {
		fmt.Println("sql is error:")
		fmt.Println(err)
		return nil, err
	}
	var goods []*model.Goods

	 */
	//defer rows.Close()
	/*
	for rows.Next() {
		fmt.Println("rows.next:")
		r := &model.Goods{}
		if err := rows.Scan(&r.ArticleId, &r.Subject, &r.Url); err != nil {
			fmt.Println("rows.next:")
			fmt.Println(err)
			return nil, err
		}
		goods = append(goods, r)
	}
	*/
	return goods, nil
}