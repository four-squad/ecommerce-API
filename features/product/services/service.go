package services

import (
	"ecommerce/features/product"
	"ecommerce/helper"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator"
)

type productUseCase struct {
	qry product.ProductData
	vld *validator.Validate
}

func New(pd product.ProductData) product.ProductService {
	return &productUseCase{
		qry: pd,
		vld: validator.New(),
	}
}
func (puu *productUseCase) GetAll() ([]product.CoreProduct, error) {
	res, err := puu.qry.GetAll()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "data tidak bisa diolah"
		}
		return []product.CoreProduct{}, errors.New(msg)
	}
	return res, nil
}

func (puu *productUseCase) Add(newProduct product.CoreProduct, token interface{}, file *multipart.FileHeader) (product.CoreProduct, error) {
	id := helper.ExtractToken(token)
	fmt.Println("======service=====")
	if file != nil {
		if file.Size > 5000000 {
			return product.CoreProduct{}, errors.New("file size is too big")
		}

		formFile, err := file.Open()
		if err != nil {
			return product.CoreProduct{}, errors.New("open file error")
		}

		if !helper.TypeFile(formFile) {
			return product.CoreProduct{}, errors.New("use jpg or png type file")
		}
		defer formFile.Close()
		formFile, _ = file.Open()
		uploadUrl, err := helper.NewMediaUpload().AvatarUpload(helper.Avatar{Avatar: formFile})

		if err != nil {
			return product.CoreProduct{}, errors.New("server error")
		}

		newProduct.Image = uploadUrl
	}

	// err := puu.vld.Struct(newProduct)
	// if err != nil {
	// 	log.Println(err)
	// 	if _, ok := err.(*validator.InvalidValidationError); ok {
	// 		log.Println(err)
	// 	}
	// 	return product.CoreProduct{}, errors.New("field required wajib diisi")
	// }
	fmt.Println("======service2=====")
	res, err := puu.qry.Add(newProduct, uint(id))
	fmt.Println("======service3=====")

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "kesalahan input"
		} else {
			msg = "data tidak bisa diolah"
		}
		return product.CoreProduct{}, errors.New(msg)
	}

	return res, nil

}

func (puu *productUseCase) GetById(idProduct uint) ([]product.CoreProduct, error) {
	res, err := puu.qry.GetById(idProduct)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "data tidak bisa diolah"
		}
		return []product.CoreProduct{}, errors.New(msg)
	}

	return res, nil
}

func (puu *productUseCase) Update(token interface{}, id uint, tmp product.CoreProduct, file *multipart.FileHeader) (product.CoreProduct, error) {
	id2 := helper.ExtractToken(token)
	if file != nil {
		if file.Size > 5000000 {
			return product.CoreProduct{}, errors.New("file size is too big")
		}

		formFile, err := file.Open()
		if err != nil {
			return product.CoreProduct{}, errors.New("open file error")
		}

		if !helper.TypeFile(formFile) {
			return product.CoreProduct{}, errors.New("use jpg or png type file")
		}
		defer formFile.Close()
		formFile, _ = file.Open()
		uploadUrl, err := helper.NewMediaUpload().AvatarUpload(helper.Avatar{Avatar: formFile})

		if err != nil {
			return product.CoreProduct{}, errors.New("server error")
		}

		tmp.Image = uploadUrl
	}
	res, err := puu.qry.Update(uint(id2), id, tmp)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "kesalahan input"
		} else {
			msg = "data tidak bisa diolah"
		}
		return product.CoreProduct{}, errors.New(msg)
	}

	return res, nil
}

func (puu *productUseCase) Delete(token interface{}, productId uint) error {
	userId := helper.ExtractToken(token)
	err := puu.qry.Delete(uint(userId), productId)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "kesalahan input"
		} else {
			msg = "data tidak bisa diolah"
		}
		return errors.New(msg)
	}

	return nil
}
