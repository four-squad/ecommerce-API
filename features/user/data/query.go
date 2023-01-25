package data

import (
	"errors"
	"log"
	"strings"

	"ecommerce/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) isDuplicate(email string) bool {
	res := Users{}
	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("Check data error : ", err.Error())
		if strings.Contains(err.Error(), "not found") {
			return false
		}

	}
	return true
}

func (uq *userQuery) Register(newUser user.Core) error {
	cnv := CoreToData(newUser)
	isDupl := uq.isDuplicate(cnv.Email)

	if isDupl != false {
		log.Println("Duplicated data, email already exist")
		return errors.New("Duplicated data, email already exist")
	}

	err := uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("Create query error : ", err.Error())
		return err
	}

	newUser.ID = cnv.ID

	return nil
}

func (uq *userQuery) Login(email string) (user.Core, error) {
	res := Users{}

	if err := uq.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error : ", err.Error())
		return user.Core{}, errors.New("Data not found")
	}

	return ToCore(res), nil
}

func (uq *userQuery) Profile(id uint) (interface{}, error) {
	resProfile := map[string]interface{}{}
	if err := uq.db.Raw("SELECT users.id, users.avatar, users.name, users.email, users.address FROM users WHERE users.id = ?", id).Find(&resProfile).Error; err != nil {
		log.Println("Get User Data by id query error : ", err.Error())
		return nil, err
	}

	resProducts := []map[string]interface{}{}
	if err := uq.db.Raw("SELECT products.id, products.title, products.price, products.description, products.image FROM products JOIN users ON products.user_id = users.id WHERE users.id = ? AND products.deleted_at IS NULL", id).Find(&resProducts).Error; err != nil {
		log.Println("Get product by user ID query error : ", err.Error())
		return nil, err
	}

	resProfile["products"] = resProducts

	return resProfile, nil
}

func (uq *userQuery) Update(userID uint, updatedProfile user.Core) (user.Core, error) {
	getEmail := Users{}
	if err := uq.db.Where("id = ?", userID).First(&getEmail).Error; err != nil {
		log.Println("Get profile by ID query error : ", err.Error())
		return user.Core{}, err
	}

	cnvUpdated := CoreToData(updatedProfile)

	if getEmail.Email != updatedProfile.Email {
		isDupl := uq.isDuplicate(cnvUpdated.Email)

		if isDupl != false {
			log.Println("Duplicated data, email already exist")
			return user.Core{}, errors.New("Duplicated data, email already exist")
		}
	}

	qry := uq.db.Model(Users{}).Where("id = ?", cnvUpdated.ID).Updates(cnvUpdated)
	err := qry.Error

	affRow := qry.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return user.Core{}, errors.New("Failed to update, no new record or data not found")
	}

	if err != nil {
		log.Println("update query error : ", err.Error())
		return user.Core{}, errors.New("Unable to update profile")
	}

	return ToCore(cnvUpdated), nil
}

func (uq *userQuery) Deactivate(id uint) error {
	qryDelete := uq.db.Delete(&Users{}, id)

	affRow := qryDelete.RowsAffected

	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("Failed to delete, data not found")
	}

	return nil
}
