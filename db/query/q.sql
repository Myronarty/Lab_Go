-- name: CreateKogut :one
INSERT INTO kogut (name, age, sex)
VALUES ($1, $2, $3)
RETURNING id, name, age, sex;

-- name: GetKogut :one
SELECT id, name, age, sex
FROM kogut
WHERE id = $1;

-- name: GetAllKoguts :many
SELECT id, name, age, sex
FROM kogut
ORDER BY id;

-- name: UpdateKogut :one
UPDATE kogut
SET name = $2, age = $3, sex = $4
WHERE id = $1
RETURNING id, name, age, sex;

-- name: DeleteKogut :exec
DELETE FROM kogut
WHERE id = $1;