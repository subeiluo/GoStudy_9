package data_layer

import (
	"9_12/mysql/model"
	"database/sql"
	"fmt"
)

func InsertBook(book *model.Book)(err error) {
	if book == nil {
		err = fmt.Errorf("invalid book parameter")
	}
	sqlstr := "select book_id from booklist where book_id=?"
	var bookid string
	err = Db.Get(&bookid, sqlstr, book.BookID)
	if err == sql.ErrNoRows {
		//如果查询失败 则数据不存在 插入数据
		sqlstr = "insert into booklist(" +
			"author," +
			"book_name," +
			"publish_time," +
			"stock_num," +
			"book_id" +
			")values(" +
			"?,?,?,?,?)"
		_, err = Db.Exec(sqlstr,
			book.Author,
			book.Bookname,
			book.Publishtime,
			book.Stocknum,
			book.BookID)
		if err != nil {
			return
		}
		return
	}
	if err!=nil {
		return
	}
	err= fmt.Errorf("book_id:%s is already exists", bookid)
	return
}