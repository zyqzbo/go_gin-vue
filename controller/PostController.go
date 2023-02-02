package controller

import (
	"github.com/gin-gonic/gin"
	"go_gin+vue/common"
	"go_gin+vue/model"
	"go_gin+vue/response"
	"go_gin+vue/vo"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type IPostController interface {
	RestController
	PageList(cxt *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func (p PostController) Create(cxt *gin.Context) {
	var requestPost vo.CreatePostRequst
	// 数据验证
	if err := cxt.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(cxt, nil, "数据验证错误")
		return
	}

	// 获取用户
	user, _ := cxt.Get("user")

	// 创建post
	post := model.Post{
		UserID:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	// 插入数据
	err := p.DB.Create(&post).Error
	if err != nil {
		panic(err)
		return
	}
	response.Success(cxt, nil, "创建成功")
}

func (p PostController) Update(cxt *gin.Context) {
	var requestPost vo.CreatePostRequst
	// 数据验证
	if err := cxt.ShouldBind(&requestPost); err != nil {
		//log.Println(err.Error())
		response.Fail(cxt, nil, "数据验证错误")
		return
	}

	// 获取path 中的postId
	postId := cxt.Params.ByName("id") // 因为文章的id是string类型
	var post model.Post
	if p.DB.Where("id=?", postId).First(&post).Error != nil {
		response.Fail(cxt, nil, "文章不存在！")
		return
	}

	// 当前用户是否为文章的作者
	// 获取用户
	user, _ := cxt.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserID {
		response.Fail(cxt, nil, "文章不属于您，请勿非法操作")
		return
	}

	// 更新文章
	//todo 这里的update参数有问题  先试下updates
	err := p.DB.Model(&post).Updates(requestPost).Error
	if err != nil {
		response.Fail(cxt, nil, "更新失败")
	}

	response.Success(cxt, gin.H{"post": post}, "更新成功")
}

func (p PostController) Show(cxt *gin.Context) {
	// 获取path 中的postId
	postId := cxt.Params.ByName("id") // 因为文章的id是string类型
	var post model.Post
	// Preload把当前作者下的分类找出并且返回
	if p.DB.Preload("Category").Where("id=?", postId).First(&post).Error != nil {
		response.Fail(cxt, nil, "文章不存在")
		return
	}

	response.Success(cxt, gin.H{"post": post}, "成功")
}

func (p PostController) Delete(cxt *gin.Context) {
	// 获取path 中的postId
	postId := cxt.Params.ByName("id") // 因为文章的id是string类型
	var post model.Post
	if p.DB.Where("id=?", postId).First(&post).Error != nil {
		response.Fail(cxt, nil, "文章不存在")
		return
	}

	// 当前用户是否为文章的作者
	// 获取用户
	user, _ := cxt.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserID {
		response.Fail(cxt, nil, "文章不属于您，请勿非法操作")
		return
	}

	// 删除
	p.DB.Delete(&post)
	response.Success(cxt, gin.H{"post": post}, "删除成功")
}

func (p PostController) PageList(cxt *gin.Context) {
	// strconv.Atoi() ：转化为字符串
	pageNum, _ := strconv.Atoi(cxt.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(cxt.DefaultQuery("pageSize", "20"))

	// 分页显示
	var posts []model.Post
	// 分页的文章按照创建时间来排序 Offset： 查询偏移亮 Find： 查询方法
	p.DB.Preload("Category").Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 前端渲染分页需要知道总条数
	var total int64
	p.DB.Model(model.Post{}).Count(&total)

	response.Success(cxt, gin.H{"data": posts, "total": total}, "成功")
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}
