package sysmodel

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	sysDictFieldNames          = builderx.RawFieldNames(&SysDict{})
	sysDictRows                = strings.Join(sysDictFieldNames, ",")
	sysDictRowsExpectAutoSet   = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysDictRowsWithPlaceHolder = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysDictModel interface {
		Insert(data SysDict) (sql.Result, error)
		FindOne(id int64) (*SysDict, error)
		Update(data SysDict) error
		Delete(id int64) error
	}

	defaultSysDictModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysDict struct {
		Id           int64     `db:"id"`             // 编号
		Value        string    `db:"value"`          // 数据值
		Label        string    `db:"label"`          // 标签名
		Type         string    `db:"type"`           // 类型
		Description  string    `db:"description"`    // 描述
		Sort         int64     `db:"sort"`           // 排序（升序）
		CreateBy     string    `db:"create_by"`      // 创建人
		CreateTime   time.Time `db:"create_time"`    // 创建时间
		LastUpdateBy string    `db:"last_update_by"` // 更新人
		UpdateTime   time.Time `db:"update_time"`    // 更新时间
		Remarks      string    `db:"remarks"`        // 备注信息
		IsDelete     int64     `db:"is_delete"`      // 是否删除  1：已删除  0：正常
	}
)

func NewSysDictModel(conn sqlx.SqlConn) SysDictModel {
	return &defaultSysDictModel{
		conn:  conn,
		table: "`sys_dict`",
	}
}

func (m *defaultSysDictModel) Insert(data SysDict) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysDictRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Value, data.Label, data.Type, data.Description, data.Sort, data.CreateBy, data.LastUpdateBy, data.Remarks, data.IsDelete)
	return ret, err
}

func (m *defaultSysDictModel) FindOne(id int64) (*SysDict, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysDictRows, m.table)
	var resp SysDict
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

func (m *defaultSysDictModel) Update(data SysDict) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysDictRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Value, data.Label, data.Type, data.Description, data.Sort, data.CreateBy, data.LastUpdateBy, data.Remarks, data.IsDelete, data.Id)
	return err
}

func (m *defaultSysDictModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
