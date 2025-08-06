-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = now(),
    updated_at = now()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;
