-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, 
    url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUserByName :many
SELECT p.*
FROM posts p
JOIN feeds f ON p.feed_id = f.id
JOIN feed_follows ff ON f.id = ff.feed_id
JOIN users u ON ff.user_id = u.id
WHERE u.name = $1
ORDER BY p.published_at DESC, p.created_at DESC
LIMIT $2;
