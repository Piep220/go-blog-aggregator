-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name, f.url, u.name AS user_name
FROM feeds f
LEFT JOIN users u on f.user_id = u.id;

-- name: GetFeedFromUrl :one
SELECT * FROM feeds
WHERE url = $1 LIMIT 1;

-- name: MarkFeedFetched :exec
UPDATE feeds 
SET updated_at = CURRENT_TIMESTAMP, 
    last_fetched_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST, created_at ASC
LIMIT 1;