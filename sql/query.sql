-- name: ListAllPages :many
SELECT * FROM pages;

-- name: ListPageFromRoute :many
SELECT * FROM pages
WHERE route_id=?;

-- name: GetPageDataById :one
SELECT page_id,data from pages
WHERE page_id=?;

-- name: ListPageFromParent :many
SELECT * FROM pages
WHERE parent_id=?;

-- name: UpdatePageDate :exec
UPDATE pages SET
data = ? 
WHERE page_id = ?;
