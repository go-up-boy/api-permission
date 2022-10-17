// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"api-permission/internal/types"
	"math"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`", "`delete_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`", "`delete_at`"), "=?,") + "=?"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
		FindOneByUsername(ctx context.Context, username string) (*Users, error)
		ExistByUsername(ctx context.Context, user *Users) bool
		PaginationList(ctx context.Context, p *types.ListPageItem) ([]Users, error)
	}

	defaultUsersModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Users struct {
		Id       int64        `db:"id"`
		Headimg  string       `db:"headimg"`
		Username string       `db:"username"` // 用户名
		Password string       `db:"password"` // 密码
		UpdateAt sql.NullTime `db:"update_at"`
		CreateAt time.Time    `db:"create_at"`
		DeleteAt sql.NullTime `db:"delete_at"`
	}

	Paging struct {
		Limit int64
		Page int64
		TotalPage int64
		TotalRow int64
		Data []Users
	}
)

func newUsersModel(conn sqlx.SqlConn) *defaultUsersModel {
	return &defaultUsersModel{
		conn:  conn,
		table: "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	var resp Users
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByUsername(ctx context.Context, username string) (*Users, error) {
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", usersRows, m.table)
	var resp Users
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) ExistByUsername(ctx context.Context, user *Users) bool {
	query := fmt.Sprintf("select %s from %s where `username` = ? and id <> ? limit 1", "id", m.table)
	var respId int64
	err := m.conn.QueryRowCtx(ctx, &respId, query, user.Username, user.Id)
	if err == sqlc.ErrNotFound {
		return false
	} else if respId == 0 {
		return false
	}
	return true
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, usersRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Headimg, data.Username, data.Password)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, data *Users) error {
	query := fmt.Sprintf("update %s set %s where `id` = ? limit 1", m.table, usersRowsWithPlaceHolder)
	result, err := m.conn.ExecCtx(ctx, query, data.Headimg, data.Username, data.Password, data.Id)
	row, err := result.RowsAffected()
	if row <= 0 {
		return errors.New("update fail")
	}
	return err
}

func (m *defaultUsersModel) PaginationList(ctx context.Context, p *types.ListPageItem) ([]Users, error) {
	var users []Users
	query := fmt.Sprintf("select %s from %s limit ?,?", "*", m.table)
	queryCount := fmt.Sprintf("select count(*) from %s",  m.table)
	m.conn.QueryRowCtx(ctx, &p.TotalRows, queryCount)
	err := m.conn.QueryRowsCtx(ctx, &users, query, (p.Page - 1) * p.Limit, p.Limit)
	p.TotalPage = int64(math.Ceil(float64(p.TotalRows) / float64(p.Limit)))
	switch err {
	case nil:
		return users, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}