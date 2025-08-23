-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
),
joined AS (
    SELECT i.*,
    u.name AS user_name,
    f.name AS feed_name
    FROM inserted_feed_follow i
    INNER JOIN users u ON i.user_id = u.id
    INNER JOIN feeds f ON i.feed_id = f.id
)
SELECT * FROM joined;

-- name: GetFeedFollowsForUser :many
SELECT 
    ff.*, 
    u.name AS user_name,
    f.name AS feed_name
FROM feed_follows ff
LEFT JOIN users u on ff.user_id = u.id
LEFT JOIN feeds f on ff.feed_id = f.id

WHERE u.name = $1
ORDER BY ff.created_at DESC;


-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows ff
WHERE ff.user_id = $1
    AND ff.feed_id = (
        SELECT id
        FROM feeds
        WHERE url = $2
    );