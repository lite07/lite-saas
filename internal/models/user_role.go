package models

type UserRole struct {
	ID     uint `json:"id" gorm:"primary_key"`
	UserID uint `json:"user_id" gorm:"uniqueIndex:fk_unq_user_role_user_id_role_id"`
	RoleID uint `json:"role_id" gorm:"uniqueIndex:fk_unq_user_role_user_id_role_id"`
	User   User
	Role   Role
}
