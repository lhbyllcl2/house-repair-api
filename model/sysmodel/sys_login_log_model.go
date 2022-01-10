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
	sysLoginLogFieldNames          = builder.RawFieldNames(&SysLoginLog{})
	sysLoginLogRows                = strings.Join(sysLoginLogFieldNames, ",")
	sysLoginLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysLoginLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLoginLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysLoginLogModel interface {
		Insert(data SysLoginLog) (sql.Result, error)
		FindOne(id int64) (*SysLoginLog, error)
		Update(data SysLoginLog) error
		Delete(id int64) error
	}

	defaultSysLoginLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysLoginLog struct {
		Id         int64     `db:"id"`          // 编号
		UserName   string    `db:"user_name"`   // 用户名
		Status     string    `db:"status"`      // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
		Ip         string    `db:"ip"`          // IP地址
		CreateTime time.Time `db:"create_time"` // 创建时间
	}
)

func NewSysLoginLogModel(conn sqlx.SqlConn) SysLoginLogModel {
	return &defaultSysLoginLogModel{
		conn:  conn,
		table: "`sys_login_log`",
	}
}

func (m *defaultSysLoginLogModel) Insert(data SysLoginLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, sysLoginLogRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.UserName, data.Status, data.Ip)
	return ret, err
}

func (m *defaultSysLoginLogModel) FindOne(id int64) (*SysLoginLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLoginLogRows, m.table)
	var resp SysLoginLog
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

func (m *defaultSysLoginLogModel) Update(data SysLoginLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLoginLogRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.UserName, data.Status, data.Ip, data.Id)
	return err
}

func (m *defaultSysLoginLogModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
