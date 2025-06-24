package id

import (
	users "chatapp_demo/internal/user/core/domain/model/vo/user"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GeneratorImpl struct {
	DB *gorm.DB
}

// Next は指定されたシーケンス名の次の番号を返します
func (g *GeneratorImpl) Next(sequenceName string) (*users.Userid, error) {

	errorValue,err := users.NewUserid(1)
	if err != nil{
		fmt.Println(err)
	}


	var seq Sequence

	tx := g.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// SELECT FOR UPDATE 相当を実現するには、Lockingを使う
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).FirstOrCreate(&seq, Sequence{SequenceName: sequenceName}).Error; err != nil {
		tx.Rollback()
		return errorValue, err
	}

	seq.CurrentValue += 1

	if err := tx.Save(&seq).Error; err != nil {
		tx.Rollback()
		return errorValue, err
	}

	returnValue := seq.CurrentValue
	if err := tx.Commit().Error; err != nil {
		return errorValue, err
	}

	validValue,err := users.NewUserid(returnValue)
	if err != nil{
		fmt.Println(err)
	}

	return validValue, nil
}
