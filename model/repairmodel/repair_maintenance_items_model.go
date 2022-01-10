package repairmodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	repairMaintenanceItemsFieldNames          = builder.RawFieldNames(&RepairMaintenanceItems{})
	repairMaintenanceItemsRows                = strings.Join(repairMaintenanceItemsFieldNames, ",")
	repairMaintenanceItemsRowsExpectAutoSet   = strings.Join(stringx.Remove(repairMaintenanceItemsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	repairMaintenanceItemsRowsWithPlaceHolder = strings.Join(stringx.Remove(repairMaintenanceItemsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	RepairMaintenanceItemsModel interface {
		Insert(data RepairMaintenanceItems) (sql.Result, error)
		FindOne(id int64) (*RepairMaintenanceItems, error)
		Update(data RepairMaintenanceItems) error
		Delete(id int64) error
	}

	defaultRepairMaintenanceItemsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RepairMaintenanceItems struct {
		Id                 int64     `db:"id"`
		ParentId           int64     `db:"parent_id"`           // 父级id
		ItemName           string    `db:"item_name"`           // 分类名称
		ProblemDescription string    `db:"problem_description"` // 问题描述
		RepairMode         string    `db:"repair_mode"`         // 维修方式
		Norms              string    `db:"norms"`               // 规格
		Condition          string    `db:"condition"`           // 起修条件
		Unit               string    `db:"unit"`                // 单位
		ExclusivePrice     float64   `db:"exclusive_price"`     // 含税单价
		IsDelete           int64     `db:"is_delete"`           // 是否被删除
		CreateTime         time.Time `db:"create_time"`         // 创建时间
		UpdateAt           time.Time `db:"update_at"`           // 更新时间
	}
)

func NewRepairMaintenanceItemsModel(conn sqlx.SqlConn) RepairMaintenanceItemsModel {
	return &defaultRepairMaintenanceItemsModel{
		conn:  conn,
		table: "`repair_maintenance_items`",
	}
}

func (m *defaultRepairMaintenanceItemsModel) Insert(data RepairMaintenanceItems) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, repairMaintenanceItemsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ParentId, data.ItemName, data.ProblemDescription, data.RepairMode, data.Norms, data.Condition, data.Unit, data.ExclusivePrice, data.IsDelete, data.UpdateAt)
	return ret, err
}

func (m *defaultRepairMaintenanceItemsModel) FindOne(id int64) (*RepairMaintenanceItems, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", repairMaintenanceItemsRows, m.table)
	var resp RepairMaintenanceItems
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

func (m *defaultRepairMaintenanceItemsModel) Update(data RepairMaintenanceItems) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, repairMaintenanceItemsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ParentId, data.ItemName, data.ProblemDescription, data.RepairMode, data.Norms, data.Condition, data.Unit, data.ExclusivePrice, data.IsDelete, data.UpdateAt, data.Id)
	return err
}

func (m *defaultRepairMaintenanceItemsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
