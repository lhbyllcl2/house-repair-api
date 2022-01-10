package repairmodel

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	repairPartQuestionFieldNames          = builder.RawFieldNames(&RepairPartQuestion{})
	repairPartQuestionRows                = strings.Join(repairPartQuestionFieldNames, ",")
	repairPartQuestionRowsExpectAutoSet   = strings.Join(stringx.Remove(repairPartQuestionFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	repairPartQuestionRowsWithPlaceHolder = strings.Join(stringx.Remove(repairPartQuestionFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	RepairPartQuestionModel interface {
		Insert(data RepairPartQuestion) (sql.Result, error)
		FindOne(id int64) (*RepairPartQuestion, error)
		Update(data RepairPartQuestion) error
		Delete(id int64) error
	}

	defaultRepairPartQuestionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RepairPartQuestion struct {
		Id           int64         `db:"id"`
		PartId       sql.NullInt64 `db:"part_id"` // 部位id
		Question     string        `db:"question"`
		Duration     int64         `db:"duration"`       // 工期，单位天
		TotalAmount  float64       `db:"total_amount"`   // 维修费用
		Warranty     int64         `db:"warranty"`       // 质保期
		WarrantyUnit string        `db:"warranty_unit"`  // 质保期单位，默认为月
		Sort         int64         `db:"sort"`           // 排序
		IsFixedPrice int64         `db:"is_fixed_price"` // 是否一口价
		IsDelete     int64         `db:"is_delete"`      // 是否被删除
	}
)

func NewRepairPartQuestionModel(conn sqlx.SqlConn) RepairPartQuestionModel {
	return &defaultRepairPartQuestionModel{
		conn:  conn,
		table: "`repair_part_question`",
	}
}

func (m *defaultRepairPartQuestionModel) Insert(data RepairPartQuestion) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, repairPartQuestionRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.PartId, data.Question, data.Duration, data.TotalAmount, data.Warranty, data.WarrantyUnit, data.Sort, data.IsFixedPrice, data.IsDelete)
	return ret, err
}

func (m *defaultRepairPartQuestionModel) FindOne(id int64) (*RepairPartQuestion, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", repairPartQuestionRows, m.table)
	var resp RepairPartQuestion
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRepairPartQuestionModel) Update(data RepairPartQuestion) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, repairPartQuestionRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.PartId, data.Question, data.Duration, data.TotalAmount, data.Warranty, data.WarrantyUnit, data.Sort, data.IsFixedPrice, data.IsDelete, data.Id)
	return err
}

func (m *defaultRepairPartQuestionModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
