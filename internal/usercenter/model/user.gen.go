package model

import "time"

const TableNameUcUser = "uc_user"

// UserM 用户表
type UserM struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键 ID" json:"id"`                       // 主键 ID
	UserID    string    `gorm:"column:user_id;not null;comment:用户 ID" json:"user_id"`                                  // 用户 ID
	Username  string    `gorm:"column:username;not null;comment:用户名称" json:"username"`                                 // 用户名称
	Status    int32     `gorm:"column:status;not null;default:1;comment:用户状态，0-禁用；1-启用" json:"status"`                 // 用户状态，0-禁用；1-启用
	Nickname  string    `gorm:"column:nickname;not null;comment:用户昵称" json:"nickname"`                                 // 用户昵称
	Password  string    `gorm:"column:password;not null;comment:用户加密后的密码" json:"password"`                             // 用户加密后的密码
	Email     string    `gorm:"column:email;not null;comment:用户电子邮箱" json:"email"`                                     // 用户电子邮箱
	Phone     string    `gorm:"column:phone;not null;comment:用户手机号" json:"phone"`                                      // 用户手机号
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`   // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:最后修改时间" json:"updated_at"` // 最后修改时间
}

// TableName UserM's table name
func (*UserM) TableName() string {
	return TableNameUcUser
}
