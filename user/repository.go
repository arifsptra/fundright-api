package user

import "gorm.io/gorm"

// save the User struct to database
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// function to save the struct data to database
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// function to find the user by email
func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// function to find the user by id
func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// function to update data User
func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}