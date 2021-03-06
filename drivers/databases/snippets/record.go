package snippets

import (
	"fmt"
	"time"
	"twilux/business/snippets"
	"twilux/drivers/databases/users"

	"gorm.io/gorm"
)

type Snippet struct {
	Id        string `gorm:"primaryKey;size:10"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         `gorm:"size:50;not null"`
	Descb     string
	Snippet   string     `gorm:"not null"`
	User      string     `gorm:"not null;size:20;index"`
	UserInfo  users.User `gorm:"foreignkey:User;references:Username"`
}

func (snippet Snippet) ToDomain() snippets.Domain {
	return snippets.Domain{
		Id:        snippet.Id,
		CreatedAt: snippet.CreatedAt,
		UpdatedAt: snippet.UpdatedAt,
		DeletedAt: snippet.DeletedAt,
		Title:     snippet.Title,
		Descb:     snippet.Descb,
		Snippet:   snippet.Snippet,
		Username:  snippet.User,
	}
}

func FromDomain(domain snippets.Domain) Snippet {
	return Snippet{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Title:     domain.Title,
		Descb:     domain.Descb,
		Snippet:   domain.Snippet,
		User:      domain.Username,
	}
}

func ToSingleDomain(data Snippet) (result snippets.Domain) {
	fmt.Println("ToSingleDomain db/snippets/record")
	fmt.Println(result)
	result = data.ToDomain()
	return result
}

func ToListDomain(data []Snippet) (result []snippets.Domain) {
	result = []snippets.Domain{}
	fmt.Println("ToListDOmain db/snippets/record")
	fmt.Println(result)
	for _, snippet := range data {
		result = append(result, snippet.ToDomain())
	}
	return result
}
