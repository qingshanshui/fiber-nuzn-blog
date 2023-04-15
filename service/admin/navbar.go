package admin

import (
	"errors"
	"fiber-nuzn-blog/models"
	admin2 "fiber-nuzn-blog/validator/form/admin"
	"github.com/jaevor/go-nanoid"
)

type Navbar struct{}

func NewNavbarService() *Navbar {
	return &Navbar{}
}

// Home 首页
func (t *Navbar) Home() []models.NavBar {
	mn := models.NewNavbar()
	ml := mn.GetWebNavBarListAll()
	return ml
}

// Add 添加分类的post
func (t *Navbar) Add(r admin2.NavbarRequest) error {
	// 组装数据
	canonical, _ := nanoid.Standard(36)
	uid := canonical()
	mn := models.NewNavbar()
	mn.Uid = uid
	mn.Title = r.Title
	mn.Url = r.Url
	mn.Show = r.Show
	mn.Sort = r.Sort
	// 业务处理
	err := mn.Create()
	if err != nil {
		return err
	}
	return nil
}
func (t *Navbar) EditView(id string) *models.NavBar {
	mn := models.NewNavbar()
	r := mn.GetLinkByUid(id)
	return r
}

// Edit 编辑分类的post
func (t *Navbar) Edit(r admin2.NavbarEditRequest) error {
	// 组装参数
	mn := models.NewNavbar()
	mn.Title = r.Title
	mn.Url = r.Url
	mn.Sort = r.Sort
	mn.Show = r.Show
	// 业务处理
	err := mn.Update(r.Id)
	if err != nil {
		return err
	}
	return nil
}

func (t *Navbar) Del(id string) error {
	mn := models.NewNavbar()
	dmn := mn.GetLinkByUid(id)
	if dmn == nil {
		return errors.New("删除的导航栏不存在")
	}
	err := dmn.Delete()
	if err != nil {
		return err
	}
	return nil
}
