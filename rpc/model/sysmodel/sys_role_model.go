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
	sysRoleFieldNames          = builderx.RawFieldNames(&SysRole{})
	sysRoleRows                = strings.Join(sysRoleFieldNames, ",")
	sysRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysRoleModel interface {
		Insert(data SysRole) (sql.Result, error)
		FindOne(id int64) (*SysRole, error)
		NameIsExist(name string) bool
		Update(data SysRole) error
		Delete(id int64) error
	}

	defaultSysRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysRole struct {
		Id           int64     `db:"id"`             // 编号
		Name         string    `db:"name"`           // 角色名称
		Remark       string    `db:"remark"`         // 备注
		CreateBy     string    `db:"create_by"`      // 创建人
		CreateTime   time.Time `db:"create_time"`    // 创建时间
		LastUpdateBy string    `db:"last_update_by"` // 更新人
		UpdateTime   time.Time `db:"update_time"`    // 更新时间
		Status       int64     `db:"status"`         // 状态  1:启用,0:禁用
		IsDelete     int64     `db:"is_delete"`      // 是否删除  1：已删除  0：正常
	}
)

func NewSysRoleModel(conn sqlx.SqlConn) SysRoleModel {
	return &defaultSysRoleModel{
		conn:  conn,
		table: "`sys_role`",
	}
}

func (m *defaultSysRoleModel) Insert(data SysRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, "name,remark,status")
	ret, err := m.conn.Exec(query, data.Name, data.Remark, data.Status)
	return ret, err
}
func (m *defaultSysRoleModel) NameIsExist(name string) bool {
	query := fmt.Sprintf("select %s from %s where name = ? and is_delete=0 limit 1", sysRoleRows, m.table)
	var resp SysRole
	_ = m.conn.QueryRow(&resp, query, name)
	if resp.Id > 0 {
		return true
	}
	return false
}
func (m *defaultSysRoleModel) FindOne(id int64) (*SysRole, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleRows, m.table)
	var resp SysRole
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

func (m *defaultSysRoleModel) Update(data SysRole) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Remark, data.CreateBy, data.LastUpdateBy, data.UpdateTime, data.Status, data.IsDelete, data.Id)
	return err
}

func (m *defaultSysRoleModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
