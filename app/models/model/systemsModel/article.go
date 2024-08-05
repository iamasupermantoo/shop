package systemsModel

import (
	"gofiber/app/models/model/types"
	"time"
)

const (
	// ArticleStatusDisable 禁用
	ArticleStatusDisable = -1
	// ArticleStatusActive 激活
	ArticleStatusActive = 10
	// ArticleTypeDefault 基础文章
	ArticleTypeDefault = 1

	// ArticleTypeHelpers 帮助中心
	ArticleTypeHelpers = 10
	// ArticleTypeAbout 关于我们
	ArticleTypeAbout = 11
	// ArticleTypeProduct 产品中心
	ArticleTypeProduct = 12
	// ArticleTypeService 服务
	ArticleTypeService = 13
	// ArticleTypeSocial 社交
	ArticleTypeSocial = 15
)

// Article 系统文章管理
type Article struct {
	types.GormModel
	AdminId uint   `gorm:"type:int unsigned not null;comment:管理ID" json:"adminId"`
	Image   string `gorm:"type:varchar(255) not null;comment:图片" json:"image"`
	Name    string `gorm:"type:varchar(60) not null;comment:标题" json:"name"`
	Content string `gorm:"type:varchar(60);comment:内容" json:"content"`
	Link    string `gorm:"type:varchar(255);comment:链接" json:"link"`
	Type    int    `gorm:"type:smallint not null;default:1;comment:1基础文章10帮助中心" json:"type"`
	Status  int    `gorm:"type:smallint not null;default:10;comment:状态 -1禁用 10开启" json:"status"`
	Data    string `gorm:"type:text;comment:数据" json:"data"`
}

// SystemArticleInfo 管理系统文章信息
type SystemArticleInfo struct {
	ID        uint      `json:"id"`        // ID
	Image     string    `json:"image"`     // 封面
	Name      string    `json:"name"`      // 文章标题
	Content   string    `json:"content"`   // 内容
	Link      string    `json:"link"`      // 链接
	CreatedAt time.Time `json:"createdAt"` // 时间
}
