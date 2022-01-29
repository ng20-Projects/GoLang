package models

// import "github.com/jinzhu/gorm"

type Book struct {
	// gorm.Model
	Author    string
	Name      string
	Pagecount int
}

// /usr/local/go/src/github.com/go-gorm/gorm (from $GOROOT)
// /home/lynx/go/src/github.com/go-gorm/gorm (from $GOPATH)
