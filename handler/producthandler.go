package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/leyle/ginbase/middleware"
	"github.com/leyle/ginbase/returnfun"
	"github.com/leyle/ginbase/util"
	"github.com/leyle/testproduct/context"
	"github.com/leyle/testproduct/model"
	. "github.com/leyle/ginbase/consolelog"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

// 新建
type AddProductForm struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}
func CreateProductHandler(ctx *context.Context, c *gin.Context) {
	var form AddProductForm
	err := c.BindJSON(&form)
	middleware.StopExec(err)

	Logger.Debugf(middleware.GetReqId(c), "新建product，name是[%s]", form.Name)

	// 检查 name 是否重复

	ds := ctx.Ds.CopyDs()
	defer ds.Close()

	product := &model.Product{
		Id: util.GenerateDataId(),
		Name: form.Name,
		Info: form.Info,
		CreateT: util.GetCurTime(),
	}
	product.UpdateT = product.CreateT

	err = product.Save(ds)
	middleware.StopExec(err)

	returnfun.ReturnOKJson(c, product)
	return
}

// 修改

// 上线

// 下线

// 明细
func GetProductInfoHandler(ctx *context.Context, c *gin.Context) {
	id := c.Param("id")

	ds := ctx.Ds.CopyDs()
	defer ds.Close()

	p, err := model.GetProductById(ds, id)
	middleware.StopExec(err)

	returnfun.ReturnOKJson(c, p)
}

// 搜索
func QueryProductHandler(ctx *context.Context, c *gin.Context) {
	var andCondition []bson.M

	name := c.Query("name")
	if name != "" {
		andCondition = append(andCondition, bson.M{"name": bson.M{"$regex": name}})
	}

	location := c.Query("location")
	if location != "" {
		andCondition = append(andCondition, bson.M{"location": location})
	}

	st := c.Query("st")
	if st != "" {
		ist, err := strconv.ParseInt(st, 10, 64)
		if err == nil {
			andCondition = append(andCondition, bson.M{"createT.seconds": bson.M{"$gte": ist}})
		}
	}

	query := bson.M{}
	if len(andCondition) > 0 {
		query = bson.M{"$and": andCondition}
	}

	db := ctx.Ds.CopyDs()
	defer db.Close()

	Q := db.C(model.CollectionNameProduct).Find(query)
	total, err := Q.Count()
	middleware.StopExec(err)

	page, size, skip := util.GetPageAndSize(c)
	var ps []*model.Product

	err = Q.Sort("-_id").Skip(skip).Limit(size).All(&ps)
	middleware.StopExec(err)

	retData := gin.H{
		"total": total,
		"page": page,
		"size": size,
		"data": ps,
	}

	returnfun.ReturnOKJson(c, retData)
	return
}
