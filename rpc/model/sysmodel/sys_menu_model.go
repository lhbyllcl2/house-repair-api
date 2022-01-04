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
	sysMenuFieldNames          = builderx.RawFieldNames(&SysMenu{})
	sysMenuRows                = strings.Join(sysMenuFieldNames, ",")
	sysMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysMenuModel interface {
		Insert(data SysMenu) (sql.Result, error)
		FindOne(id int64) (*SysMenu, error)
		Update(data SysMenu) error
		Delete(id int64) error
	}

	defaultSysMenuModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysMenu struct {
		Id            int64     `db:"id"`        // 编号
		Name          string    `db:"name"`      // 菜单名称
		ParentId      int64     `db:"parent_id"` // 父菜单ID，一级菜单为0
		Url           string    `db:"url"`
		Perms         string    `db:"perms"`          // 授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)
		Type          int64     `db:"type"`           // 类型   0：目录   1：菜单   2：按钮
		Icon          string    `db:"icon"`           // 菜单图标
		OrderNum      int64     `db:"order_num"`      // 排序
		CreateBy      string    `db:"create_by"`      // 创建人
		CreateTime    time.Time `db:"create_time"`    // 创建时间
		LastUpdateBy  string    `db:"last_update_by"` // 更新人
		UpdateTime    time.Time `db:"update_time"`    // 更新时间
		IsDelete      int64     `db:"is_delete"`      // 是否删除  1：已删除  0：正常
		VuePath       string    `db:"vue_path"`       // vue系统的path
		VueComponent  string    `db:"vue_component"`  // vue的页面
		VueIcon       string    `db:"vue_icon"`       // vue的图标
		VueRedirect   string    `db:"vue_redirect"`   // vue的路由重定向
		BackgroundUrl string    `db:"background_url"` // 后台地址
	}
)

func NewSysMenuModel(conn sqlx.SqlConn) SysMenuModel {
	return &defaultSysMenuModel{
		conn:  conn,
		table: "`sys_menu`",
	}
}

func (m *defaultSysMenuModel) Insert(data SysMenu) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysMenuRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.ParentId, data.Url, data.Perms, data.Type, data.Icon, data.OrderNum, data.CreateBy, data.LastUpdateBy, data.IsDelete, data.VuePath, data.VueComponent, data.VueIcon, data.VueRedirect, data.BackgroundUrl)
	return ret, err
}

func (m *defaultSysMenuModel) FindOne(id int64) (*SysMenu, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysMenuRows, m.table)
	var resp SysMenu
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

func (m *defaultSysMenuModel) Update(data SysMenu) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysMenuRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.ParentId, data.Url, data.Perms, data.Type, data.Icon, data.OrderNum, data.CreateBy, data.LastUpdateBy, data.IsDelete, data.VuePath, data.VueComponent, data.VueIcon, data.VueRedirect, data.BackgroundUrl, data.Id)
	return err
}

func (m *defaultSysMenuModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
