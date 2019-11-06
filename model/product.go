package model

import (
	. "github.com/leyle/ginbase/consolelog"
	"github.com/leyle/ginbase/dbandmq"
	"github.com/leyle/ginbase/util"
	"gopkg.in/mgo.v2"
)

// 产品的新增、修改、上线、下线、明细查看、搜素
const CollectionNameProduct = "product"
var IKProduct = &dbandmq.IndexKey{
	Collection:    CollectionNameProduct,
	SingleKey:     []string{"name"},
	UniqueKey:     []string{"name"},
}
type Product struct {
	Id string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Info string `json:"info" bson:"info"`
	CreateT *util.CurTime `json:"createT" bson:"createT"`
	UpdateT *util.CurTime `json:"-" bson:"updateT"`
}

func (p *Product) Save(db *dbandmq.Ds) error {
	err := db.C(CollectionNameProduct).Insert(p)
	if err != nil {
		Logger.Errorf("", "存储product[%s][%s]表发生严重错误, %s", p.Id, p.Name, err.Error())
		return err
	}
	return nil
}

func Save(db *dbandmq.Ds, p *Product) error {
	return nil
}

func (p *Product) GetProductById(db *dbandmq.Ds, id string) error {
	return nil
}

func GetProductById(db *dbandmq.Ds, id string) (*Product, error) {
	var p *Product
	err := db.C(CollectionNameProduct).FindId(id).One(&p)
	if err != nil && err != mgo.ErrNotFound {
		Logger.Errorf("", "根据id[%s]读取product数据，发生错误, %s", id, err.Error())
		return nil, err
	}

	return p, nil
}