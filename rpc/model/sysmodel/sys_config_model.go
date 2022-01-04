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
	sysConfigFieldNames          = builderx.RawFieldNames(&SysConfig{})
	sysConfigRows                = strings.Join(sysConfigFieldNames, ",")
	sysConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(sysConfigFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(sysConfigFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysConfigModel interface {
		Insert(data SysConfig) (sql.Result, error)
		FindOne(id int64) (*SysConfig, error)
		Update(data SysConfig) error
		Delete(id int64) error
	}

	defaultSysConfigModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysConfig struct {
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

func NewSysConfigModel(conn sqlx.SqlConn) SysConfigModel {
	return &defaultSysConfigModel{
		conn:  conn,
		table: "`sys_config`",
	}
}

func (m *defaultSysConfigModel) Insert(data SysConfig) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysConfigRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Value, data.Label, data.Type, data.Description, data.Sort, data.CreateBy, data.LastUpdateBy, data.Remarks, data.IsDelete)
	return ret, err
}

func (m *defaultSysConfigModel) FindOne(id int64) (*SysConfig, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysConfigRows, m.table)
	var resp SysConfig
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

func (m *defaultSysConfigModel) Update(data SysConfig) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysConfigRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Value, data.Label, data.Type, data.Description, data.Sort, data.CreateBy, data.LastUpdateBy, data.Remarks, data.IsDelete, data.Id)
	return err
}

func (m *defaultSysConfigModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
