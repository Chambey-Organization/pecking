package models

type Lesson struct {
	ID         uint            `gorm:"primary_key;autoIncrement:true" json:"id"`
	PracticeID uint            `json:"-"`
	Title      string          `json:"lessonTitle"`
	Active     bool            `json:"active"`
	TimerCount *uint           `json:"timerCount"`
	Content    []LessonContent `gorm:"foreignKey:LessonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"content"`
}
