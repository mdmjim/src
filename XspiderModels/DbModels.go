package XspiderModels

import (
    //"fmt"
    //"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Xusr struct{
	Id bson.ObjectId "bson:'_id'"
	Email string
	Pwd string
	Fee float64
	Role string
	CreatedAt int64
}

type Recharge struct{ 
	Id bson.ObjectId "bson:'_id'"
	Uid bson.ObjectId
	Fee float64
	CreatedAt int64
	State int
}

type Billing struct{
	Id bson.ObjectId "bson:'_id'"
	Uid bson.ObjectId
	Tid bson.ObjectId
	Title string
	Fee float64
	CreatedAt int64
}

type App struct{ 
	Id bson.ObjectId "bson:'_id'"
	Name string
	Img string
	Desc string
	Schemas string
}

type Keyword struct{ 
	Id bson.ObjectId "bson:'_id'"
	Uid bson.ObjectId
	Aid bson.ObjectId
	Txt string
}

type Source struct{
	Id  bson.ObjectId "bson:'_id'"
	Aid bson.ObjectId
	Name string
	Img string
	Desc string
	SnapQry string
	DetailQry string
	PosterQry string
	PxPerData float64
}

type Task struct{
	Id  bson.ObjectId "bson:'_id'"
	Aid bson.ObjectId
	Uid bson.ObjectId
	Kid bson.ObjectId
	Sid bson.ObjectId
	Url string
	IsStarted bool 
} 

type XData struct{ 
	Id  bson.ObjectId "bson:'_id'"
	Uid bson.ObjectId
	Aid bson.ObjectId
	Kid bson.ObjectId
	Sid bson.ObjectId
	Tid bson.ObjectId
	DetailUrl string
	OwnerId bson.ObjectId 
	NameValues map[string]interface{}
	CreatedAt int64
	LastUpdatedAt int64
	NextUpdatedAt int64
	UpdateInterval int64
	DetailQry string
	Status int
	DataType int
}

const(
	StatusStructrueChanged=0
	StatusRemoved=1
	StatusBlocked=2
	StatusExcluded=3
	StatusOk=4
)

