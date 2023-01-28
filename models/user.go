package models

// User 用户表
type User struct {
	BasesModel
	Uid       string `gorm:"not null;comment:用户uid，唯一id" json:"uid"`       // 用户uid，唯一id
	Username  string `gorm:"comment:用户名" json:"username"`                  // 用户名
	Nickname  string `gorm:"comment:昵称" json:"nickname"`                   // 昵称
	Password  string `gorm:"comment:密码" json:"password"`                   // 密码
	Telephone uint8  `gorm:"comment:手机号" json:"telephone"`                 // 手机号
	Email     string `gorm:"comment:邮箱" json:"email"`                      // 邮箱
	HeadImg   string `gorm:"comment:头像" json:"headImg"`                    // 头像
	Sex       uint8  `gorm:"default:3;comment:性别，0:男 1:女 3:未知" json:"sex"` // 性别，0:男 1:女 3:未知
	Age       string `gorm:"comment:年龄" json:"age"`                        // 年龄
	Birthday  string `gorm:"comment:生日" json:"birthday"`                   // 生日
	Status    uint8  `gorm:"default:1;comment:状态 0:禁用 1：正常" json:"status"` // 状态 0:禁用 1：正常
}

// TableName 重命名表
func (u *User) TableName() string {
	return "nuzn_user"
}

// NewUser new一个空的结构体
func NewUser() *User {
	return &User{}
}

// Login 登录
func (u *User) Login(username, password string) error {
	return u.DB().Where("username=? AND password =?", username, password).Find(u).Error
}

// Create 创建
func (u *User) Create() error {
	return u.DB().Create(u).Error
}

// Delete 删除
func (u *User) Delete() error {
	return u.DB().Model(u).Delete(u).Error
}

// Update 修改
func (u *User) Update(uid string) error {
	return u.DB().Model(u).Where("uid = ? ", uid).Updates(u).Error
}

// GetUserByUid 通过uid查询 用户 详情
func (u *User) GetUserByUid(uid string) *User {
	if err := u.DB().Where("uid = ?", uid).First(u).Error; err != nil {
		return nil
	}
	return u
}

// GetUserListCount 获取 当前用户 总数
func (u *User) GetUserListCount() int64 {
	var count int64
	if err := u.DB().Model(u).Count(&count).Error; err != nil {
		return 0
	}
	return count
}
