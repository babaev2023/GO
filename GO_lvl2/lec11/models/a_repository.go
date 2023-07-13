package models

import (
	"github.com/babaevs2023/GO/GO_lvl2/lec11/storage"
	_ "github.com/lib/pq"
)

func GetAllArticles(a *[]Article) error {
	if err := storage.DB.Find(a).Error; err != nil {
		return err
	}
	return nil

	//return storage.DB.Find(a).Error - можно так
}

func AddNewArticle(a *Article) error {
	return storage.DB.Create(a).Error
}

func GetArticleById(a *Article, id string) error {
	return storage.DB.Where("id = ?", id).First(a).Error
}

func UpdateArticleById(a *Article, id string) error {
	return storage.DB.Update(a).Error
}
func DeleteArticleById(a *Article, id string) error {
	return storage.DB.Where("id = ?", id).Delete(a).Error
}
