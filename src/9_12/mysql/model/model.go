package model

import "time"

//书籍对象
type Book struct {
	Id          int64     `db:"id"`
	Author      string    `db:"author"`
	Bookname    string    `db:"book_name"`
	Publishtime time.Time `db:"publish_time"`
	Stocknum    uint      `db:"stock_num"`
	BookID      string    `db:"book_id"`
}