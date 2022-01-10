package repairmodel

import (
	"database/sql"

	"gorm.io/gorm"
)

type (
	RepairPartModel interface {
		Insert(data RepairPart) error
		FindOne(id int64) (*RepairPart, error)
		Update(data RepairPart) error
		Delete(id int64) error
	}

	defaultRepairPartModel struct {
		conn  *gorm.DB
		table string
	}

	RepairPart struct {
		Id       int64         `gorm:"id"`
		Name     string        `gorm:"name"`      // 名称
		ParentId sql.NullInt64 `gorm:"parent_id"` // 父级id
		Sort     int64         `gorm:"sort"`      // 排序
		IsDelete int64         `gorm:"is_delete"` // 是否被删除
	}
)

func NewRepairPartModel(conn *gorm.DB) RepairPartModel {
	return &defaultRepairPartModel{
		conn:  conn,
		table: "`repair_part`",
	}
}

func (m *defaultRepairPartModel) Insert(data RepairPart) error {
	err := m.conn.Create(data).Error
	return err
}

func (m *defaultRepairPartModel) FindOne(id int64) (*RepairPart, error) {
	var resp RepairPart
	err := m.conn.Where("id=?", id).Find(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRepairPartModel) Update(data RepairPart) error {
	err := m.conn.Model(&RepairPart{}).Updates(data).Error
	return err
}

func (m *defaultRepairPartModel) Delete(id int64) error {
	return m.conn.Delete(&RepairPart{}, id).Error
}
