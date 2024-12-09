// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package tdb

import (
	"context"
	"database/sql"
)

const getPageDataById = `-- name: GetPageDataById :one
SELECT page_id,data from pages
WHERE page_id=?
`

type GetPageDataByIdRow struct {
	PageID int32
	Data   sql.NullString
}

func (q *Queries) GetPageDataById(ctx context.Context, pageID int32) (GetPageDataByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getPageDataById, pageID)
	var i GetPageDataByIdRow
	err := row.Scan(&i.PageID, &i.Data)
	return i, err
}

const listAllPages = `-- name: ListAllPages :many
SELECT page_id, parent_id, route_id, label, data FROM pages
`

func (q *Queries) ListAllPages(ctx context.Context) ([]Page, error) {
	rows, err := q.db.QueryContext(ctx, listAllPages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Page
	for rows.Next() {
		var i Page
		if err := rows.Scan(
			&i.PageID,
			&i.ParentID,
			&i.RouteID,
			&i.Label,
			&i.Data,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPageFromParent = `-- name: ListPageFromParent :many
SELECT page_id, parent_id, route_id, label, data FROM pages
WHERE parent_id=?
`

func (q *Queries) ListPageFromParent(ctx context.Context, parentID sql.NullInt32) ([]Page, error) {
	rows, err := q.db.QueryContext(ctx, listPageFromParent, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Page
	for rows.Next() {
		var i Page
		if err := rows.Scan(
			&i.PageID,
			&i.ParentID,
			&i.RouteID,
			&i.Label,
			&i.Data,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPageFromRoute = `-- name: ListPageFromRoute :many
SELECT page_id, parent_id, route_id, label, data FROM pages
WHERE route_id=?
`

func (q *Queries) ListPageFromRoute(ctx context.Context, routeID sql.NullInt32) ([]Page, error) {
	rows, err := q.db.QueryContext(ctx, listPageFromRoute, routeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Page
	for rows.Next() {
		var i Page
		if err := rows.Scan(
			&i.PageID,
			&i.ParentID,
			&i.RouteID,
			&i.Label,
			&i.Data,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePageDate = `-- name: UpdatePageDate :exec
UPDATE pages SET
data = ? 
WHERE page_id = ?
`

type UpdatePageDateParams struct {
	Data   sql.NullString
	PageID int32
}

func (q *Queries) UpdatePageDate(ctx context.Context, arg UpdatePageDateParams) error {
	_, err := q.db.ExecContext(ctx, updatePageDate, arg.Data, arg.PageID)
	return err
}