package domains

type Driver struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(300)" validate:"required,max=100,min=1" json:"name"`
	Email string `gorm:"type:varchar(300)" validate:"required,email,max=100,min=1" json:"email"`
}
