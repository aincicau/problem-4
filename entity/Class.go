package entity

type Class struct {
	ID          int        `gorm:"type:int;not null;primary_key"`
	Title       string     `gorm:"type:varchar(100);not null"`
	Description string     `gorm:"type:varchar(100);not null"`
	Students    []*Student `gorm:"many2many:student_classes"`
}
