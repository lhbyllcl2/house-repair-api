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
	sysUserFieldNames          = builderx.RawFieldNames(&SysUser{})
	sysUserRows                = strings.Join(sysUserFieldNames, ",")
	sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`", "`is_delete`"), ",")
	sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysUserModel interface {
		Insert(data SysUser) (sql.Result, error)
		UserInfoWithFields(name, email, mobile string) (*SysUser, error)
		FindOne(id int64) (*SysUser, error)
		FindOneByName(name string) (*SysUser, error)
		Update(data SysUser) error
		Delete(id int64) error
	}

	defaultSysUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysUser struct {
		NickName     string    `db:"nick_name"`      // 昵称
		Avatar       string    `db:"avatar"`         // 头像
		Password     string    `db:"password"`       // 密码
		Salt         string    `db:"salt"`           // 加密盐
		Email        string    `db:"email"`          // 邮箱
		Mobile       string    `db:"mobile"`         // 手机号
		Status       int64     `db:"status"`         // 状态  0：禁用   1：正常
		CreateBy     string    `db:"create_by"`      // 创建人
		CreateTime   time.Time `db:"create_time"`    // 创建时间
		LastUpdateBy string    `db:"last_update_by"` // 更新人
		UpdateTime   time.Time `db:"update_time"`    // 更新时间
		IsDelete     int64     `db:"is_delete"`      // 是否删除  1：已删除  0：正常
		JobId        int64     `db:"job_id"`         // 岗位Id
		Id           int64     `db:"id"`             // 编号
		Name         string    `db:"name"`           // 用户名
	}
)

func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &defaultSysUserModel{
		conn:  conn,
		table: "`sys_user`",
	}
}

func (m *defaultSysUserModel) Insert(data SysUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)", m.table, sysUserRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.NickName, data.Avatar, data.Password, data.Salt, data.Email, data.Mobile, data.Status, data.CreateBy, data.LastUpdateBy, data.JobId, data.Name)
	return ret, err
}

func (m *defaultSysUserModel) UserInfoWithFields(name, email, mobile string) (*SysUser, error) {
	var resp SysUser
	query := fmt.Sprintf("select %s from %s where `name` = ? or `email`=? or `mobile`=? limit 1", sysUserRows, m.table)
	err := m.conn.QueryRow(&resp, query, name, email, mobile)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultSysUserModel) FindOne(id int64) (*SysUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRows, m.table)
	var resp SysUser
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

func (m *defaultSysUserModel) FindOneByName(name string) (*SysUser, error) {
	var resp SysUser
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", sysUserRows, m.table)
	err := m.conn.QueryRow(&resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysUserModel) Update(data SysUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.NickName, data.Avatar, data.Password, data.Salt, data.Email, data.Mobile, data.Status, data.CreateBy, data.LastUpdateBy, data.IsDelete, data.JobId, data.Name, data.Id)
	return err
}

func (m *defaultSysUserModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
