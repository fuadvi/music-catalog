package memberships

import (
	"github.com/fuadvi/music-catalog/internal/models/memberships"
)

func (r *Repository) CreateUser(model memberships.User) error {
	return r.db.Create(&model).Error
}

func (r *Repository) GetUser(email, username string, id uint) (*memberships.User, error) {
	user := memberships.User{}
	res := r.db.Where("email = ?", email).Or("username = ?", username).Or("id = ?", id).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil

}
