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
	sysRoleMenuFieldNames          = builder.RawFieldNames(&SysRoleMenu{})
	sysRoleMenuRows                = strings.Join(sysRoleMenuFieldNames, ",")
	sysRoleMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleMenuFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysRoleMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleMenuFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysRoleMenuModel interface {
		Insert(data SysRoleMenu) (sql.Result, error)
		FindOne(id int64) (*SysRoleMenu, error)
		Update(data SysRoleMenu) error
		Delete(id int64) error
	}

	defaultSysRoleMenuModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysRoleMenu struct {
		Id           int64     `db:"id"`             // 编号
		RoleId       int64     `db:"role_id"`        // 角色ID
		MenuId       int64     `db:"menu_id"`        // 菜单ID
		CreateBy     string    `db:"create_by"`      // 创建人
		CreateTime   time.Time `db:"create_time"`    // 创建时间
		LastUpdateBy string    `db:"last_update_by"` // 更新人
		UpdateTime   time.Time `db:"update_time"`    // 更新时间
	}
)

func NewSysRoleMenuModel(conn sqlx.SqlConn) SysRoleMenuModel {
	return &defaultSysRoleMenuModel{
		conn:  conn,
		table: "`sys_role_menu`",
	}
}

func (m *defaultSysRoleMenuModel) Insert(data SysRoleMenu) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, sysRoleMenuRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.RoleId, data.MenuId, data.CreateBy, data.LastUpdateBy)
	return ret, err
}

func (m *defaultSysRoleMenuModel) FindOne(id int64) (*SysRoleMenu, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleMenuRows, m.table)
	var resp SysRoleMenu
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

func (m *defaultSysRoleMenuModel) Update(data SysRoleMenu) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleMenuRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.RoleId, data.MenuId, data.CreateBy, data.LastUpdateBy, data.Id)
	return err
}

func (m *defaultSysRoleMenuModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
