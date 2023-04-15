package admin

import (
	"errors"
	"fiber-nuzn-blog/models"
	admin2 "fiber-nuzn-blog/validator/form/admin"
	"github.com/jaevor/go-nanoid"
)

type Link struct{}

func NewLinkService() *Link {
	return &Link{}
}

// Home 首页
func (t *Link) Home() []models.Link {
	ml := models.NewLink()
	l := ml.GetAdminLinkList()
	return l
}

// Add 添加分类的post
func (t *Link) Add(r admin2.LinkCreateRequest) error {
	ml := models.NewLink()
	// 组装数据
	canonical, _ := nanoid.Standard(36)
	uid := canonical()
	ml.Uid = uid
	ml.Title = r.Title
	ml.Url = r.Url
	ml.Show = r.Show
	ml.Sort = r.Sort
	ml.Describe = r.Describe
	// 业务处理
	err := ml.Create()
	if err != nil {
		return err
	}
	return nil
}

func (t *Link) EditView(id string) *models.Link {
	ml := models.NewLink()
	r := ml.GetLinkByUid(id)
	return r
}

// Edit 编辑分类的post
func (t *Link) Edit(r admin2.LinkEditRequest) error {
	// 组装参数
	ml := models.NewLink()
	// 组装数据
	ml.Title = r.Title
	ml.Url = r.Url
	ml.Sort = r.Sort
	ml.Show = r.Show
	ml.Describe = r.Describe
	// 业务处理
	err := ml.Update(r.Id)
	if err != nil {
		return err
	}
	return nil
}

func (t *Link) Del(id string) error {
	// 业务处理
	ml := models.NewLink()
	dml := ml.GetLinkByUid(id)
	if dml == nil {
		return errors.New("删除的友链不存在")
	}
	err := dml.Delete()
	if err != nil {
		return err
	}
	return nil
}
