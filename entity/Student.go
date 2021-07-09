package entity

type Student struct {
	ID        int      `gorm:"type:int;not null;primary_key"`
	LastName  string   `gorm:"type:varchar(100);not null;column:lastname"`
	FirstName string   `gorm:"type:varchar(100);not null;column:firstname"`
	Age       int      `gorm:"type:int;not null"`
	Classes   []*Class `gorm:"many2many:student_classes"`
}
