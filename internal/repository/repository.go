package repository

import (
	"channels/internal/db"
	"channels/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{connection: conn}
}

// func (r *Repository) GetUserByChan(c chan models.All) (chan models.All, error) {
// 	var all models.All
// 	tx := db.DataB.Table("name").Find(&all)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	fmt.Println("name: ", all.FirstName)

// 	tx = db.DataB.Table("pic").Find(&all)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	fmt.Println("pic: ", all.Pic)
// 	tx = db.DataB.Table("info").Find(&all)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	fmt.Println("info: ", all.Email)
// 	c <- all
// 	return c, nil
// }


// func (r *Repository) GetUserByChan(c chan []models.All) (chan []models.All, error) {
//     batchSize := 100
//     var all []models.All
//     if err := db.DataB.Table("name").FindInBatches(&all, batchSize, func(tx *gorm.DB, batchIndex int) error {
//         fmt.Println("Batch: ", batchIndex)
//         c <- all
//         return nil
//     }).Error; err != nil {
//         return nil, err
//     }

//     all = nil
//     if err := db.DataB.Table("pic").FindInBatches(&all, batchSize, func(tx *gorm.DB, batchIndex int) error {
//         fmt.Println("Batch: ", batchIndex)
//         c <- all
//         return nil
//     }).Error; err != nil {
//         return nil, err
//     }

//     all = nil
//     if err := db.DataB.Table("info").FindInBatches(&all, batchSize, func(tx *gorm.DB, batchIndex int) error {
//         fmt.Println("Batch: ", batchIndex)
//         c <- all
//         return nil
//     }).Error; err != nil {
//         return nil, err
//     }

//     close(c)
//     return c, nil
// }

func (r *Repository) GetUserByChan(c chan []models.All) (chan []models.All, error) {
    batchSize := 100

    var all []models.All
    if err := db.DataB.Table("name").Find(&all).Error; err != nil {
        return nil, err
    }
    if err := db.DataB.Table("name").FindInBatches(&all, batchSize, func(tx *gorm.DB, batchIndex int) error {
        fmt.Println("Batch: ", batchIndex)
        c <- all
        return nil
    }).Error; err != nil {
        return nil, err
    }

    all = nil
    if err := db.DataB.Table("pic").Find(&all).Error; err != nil {
        return nil, err
    }
    if err := db.DataB.Table("pic").FindInBatches(&all, batchSize, func(tx *gorm.DB, batchIndex int) error {
        fmt.Println("Batch: ", batchIndex)
        c <- all
        return nil
    }).Error; err != nil {
        return nil, err
    }

    all = nil
    if err := db.DataB.Table("info").Find(&all).Error; err != nil {
        return nil, err
    }
    if err := db.DataB.Table("info").FindInBatches(&all, batchSize, func(tx *gorm.DB, batchIndex int) error {
        fmt.Println("Batch: ", batchIndex)
        c <- all
        return nil
    }).Error; err != nil {
        return nil, err
    }

    //close(c)
    return c, nil
}
