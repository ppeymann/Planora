module github.com/ppeymann/Planora/todo

go 1.24.2

replace github.com/ppeymann/Planora.git => ../

require (
	github.com/ppeymann/Planora.git v0.0.0-00010101000000-000000000000
	gorm.io/gorm v1.30.3
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.28.0 // indirect
)
