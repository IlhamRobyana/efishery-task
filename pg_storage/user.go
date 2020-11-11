package pg_storage

import (
	"github.com/ilhamrobyana/efishery-task/entity"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct{}

func (u *User) Create(user entity.User) (entity.User, error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return entity.User{}, e
	}

	registered := checkRegisteredByPhone(client, user.Phone)
	if !registered {
		user.UUID, _ = uuid.NewV4()
		e = client.Create(&user).Error
		if e != nil {
			return entity.User{}, e
		}
	}
	return user, nil
}

func (u *User) GetByPhone(phone string) (user entity.User, e error) {
	client, e := GetPGClient()
	defer client.Close()

	if e != nil {
		return
	}
	e = client.Where("phone = ?", phone).First(&user).Error
	return user, e
}

func checkRegisteredByPhone(client *gorm.DB, phone string) bool {
	user := entity.User{}
	client.Where("phone=?", phone).Find(&user)
	if user.UUID.String() != "" {
		return true
	}
	return false
}
