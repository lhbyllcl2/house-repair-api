package sysmodel

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
	sysLogFieldNames          = builder.RawFieldNames(&SysLog{})
	sysLogRows                = strings.Join(sysLogFieldNames, ",")
	sysLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysLogModel interface {
		Insert(data SysLog) (sql.Result, error)
		FindOne(id int64) (*SysLog, error)
		Update(data SysLog) error
		Delete(id int64) error
	}

	defaultSysLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysLog struct {
		Id           int64     `db:"id"`             // 编号
		UserName     string    `db:"user_name"`      // 用户名
		Operation    string    `db:"operation"`      // 用户操作
		Method       string    `db:"method"`         // 请求方法
		Params       string    `db:"params"`         // 请求参数
		Time         int64     `db:"time"`           // 执行时长(毫秒)
		Ip           string    `db:"ip"`             // IP地址
		CreateBy     string    `db:"create_by"`      // 创建人
		CreateTime   time.Time `db:"create_time"`    // 创建时间
		LastUpdateBy string    `db:"last_update_by"` // 更新人
		UpdateTime   time.Time `db:"update_time"`    // 更新时间
	}
)

func NewSysLogModel(conn sqlx.SqlConn) SysLogModel {
	return &defaultSysLogModel{
		conn:  conn,
		table: "`sys_log`",
	}
}

func (m *defaultSysLogModel) Insert(data SysLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysLogRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.UserName, data.Operation, data.Method, data.Params, data.Time, data.Ip, data.CreateBy, data.LastUpdateBy)
	return ret, err
}

func (m *defaultSysLogModel) FindOne(id int64) (*SysLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLogRows, m.table)
	var resp SysLog
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

func (m *defaultSysLogModel) Update(data SysLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLogRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.UserName, data.Operation, data.Method, data.Params, data.Time, data.Ip, data.CreateBy, data.LastUpdateBy, data.Id)
	return err
}

func (m *defaultSysLogModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
