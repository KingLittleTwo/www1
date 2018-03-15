package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	title           string
	content         string    `orm:"index"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Author          string
	Views           int64     `orm:"index"`
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastuUerId int64
}

func Init() {
	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}
