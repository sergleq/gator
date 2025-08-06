-- name: CreateFeedFollow :one
WITH inserted AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT 
  inserted.id,
  inserted.created_at,
  inserted.updated_at,
  users.name AS user_name,
  feeds.name AS feed_name
FROM inserted
JOIN users ON inserted.user_id = users.id
JOIN feeds ON inserted.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT
  ff.id,
  ff.created_at,
  ff.updated_at,
  users.name AS user_name,
  feeds.name AS feed_name,
  feeds.url AS feed_url
FROM feed_follows ff
JOIN users ON ff.user_id = users.id
JOIN feeds ON ff.feed_id = feeds.id
WHERE ff.user_id = $1
ORDER BY ff.created_at DESC;

-- name: DeleteFeedFollowByUserAndURL :exec
DELETE FROM feed_follows
WHERE feed_follows.user_id = $1
  AND feed_follows.feed_id = (
    SELECT id FROM feeds WHERE url = $2
  );



