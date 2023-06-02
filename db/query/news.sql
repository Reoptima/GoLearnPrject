-- name: CreateNews :one
INSERT INTO news (
                  text
) VALUES (
          $1
         ) RETURNING *;
-- name: GetNews :one
SELECT * FROM news
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListNews :many
SELECT * FROM news
limit $1
    offset $2;

-- name: DeleteNews :exec
DELETE FROM news
where id = $1;