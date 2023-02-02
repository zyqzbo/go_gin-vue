package controller

import (
	"github.com/gin-gonic/gin"
	"go_gin+vue/model"
	"go_gin+vue/repository"
	"go_gin+vue/response"
	"go_gin+vue/vo"
	"log"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	// 按住 control键 鼠标点击可以显示Generate 点进去点 Implement Methods 可以自动创建下面方法对接口的实现
	//DB *gorm.DB // DB连接池

	// 改用repository 文件的封装好的
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	//db := common.GetDB()             // 取数据库连接池
	//改用repository 文件的封装好的DB连接池
	repository := repository.NewCategoryRepository() // 取数据库连接池
	repository.DB.AutoMigrate(model.Category{})      // 添加自动迁移 （自动生成数据库表和字段）

	return CategoryController{Repository: repository} // 将db存进去
}

func (c CategoryController) Create(cxt *gin.Context) {
	// 绑定body中的参数 下面两种都是

	//var requestCategory model.Category
	//cxt.Bind(&requestCategory)
	//
	//if requestCategory.Name == "" {
	//	response.Fail(cxt, nil, "数据验证错误，分类名称必填")
	//	return
	//}

	var requestCategory vo.CreateCategoryRequest
	if err := cxt.ShouldBind(&requestCategory); err != nil {
		log.Println(err.Error())
		response.Fail(cxt, nil, "数据验证错误")
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		//response.Fail(cxt, nil, "创建失败")
		panic(err)
		return
	}

	response.Success(cxt, gin.H{"category": category}, "")
}

func (c CategoryController) Update(cxt *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := cxt.ShouldBind(&requestCategory); err != nil {
		response.Fail(cxt, nil, "数据验证错误，分类名称必填")
		return
	}

	// 获取path 中的参数
	categoryId, _ := strconv.Atoi(cxt.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(cxt, nil, "分类不存在")
		return
	}

	// 更新分类 参数可以传入：map struct name的value
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(cxt, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(cxt *gin.Context) {
	// 获取path 中的参数
	categoryId, _ := strconv.Atoi(cxt.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(cxt, nil, "分类不存在")
		return
	}

	response.Success(cxt, gin.H{"category": category}, "查询成功")
}

func (c CategoryController) Delete(cxt *gin.Context) {
	// 获取path 中的参数
	categoryId, _ := strconv.Atoi(cxt.Params.ByName("id"))

	err := c.Repository.DeleteById(categoryId)
	if err != nil {
		response.Fail(cxt, nil, "删除失败，请重试")
		return
	}

	response.Success(cxt, nil, "删除成功")
}
