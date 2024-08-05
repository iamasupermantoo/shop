package systemsService

import (
	"github.com/gomodule/redigo/redis"
	"gofiber/app/models/model/systemsModel"
	"gofiber/app/module/database"
)

type SystemArticle struct {
	rdsConn redis.Conn
	adminId uint
}

func NewSystemArticle(rdsConn redis.Conn, adminId uint) *SystemArticle {
	return &SystemArticle{rdsConn: rdsConn, adminId: adminId}
}

// GetArticleList 获取文章列表
func (_SystemArticle *SystemArticle) GetArticleList(articleType int) []*systemsModel.SystemArticleInfo {
	articleList := make([]*systemsModel.Article, 0)
	database.Db.Model(&systemsModel.Article{}).Where("admin_id = ?", _SystemArticle.adminId).Where("type = ?", articleType).
		Where("status = ?", systemsModel.ArticleStatusActive).Find(&articleList)

	data := make([]*systemsModel.SystemArticleInfo, 0)
	for _, article := range articleList {
		data = append(data, &systemsModel.SystemArticleInfo{
			ID:        article.ID,
			Image:     article.Image,
			Name:      article.Name,
			Link:      article.Link,
			CreatedAt: article.CreatedAt,
		})
	}
	return data
}
