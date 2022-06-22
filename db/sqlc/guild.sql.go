// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: guild.sql

package db

import (
	"context"
)

const createGuild = `-- name: CreateGuild :one
INSERT INTO guild (discord_id, name, icon, owner_id)
VALUES ($1, $2, $3, $4)
RETURNING id, discord_id, name, icon, owner_id
`

type CreateGuildParams struct {
	DiscordID string `json:"discord_id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	OwnerID   string `json:"owner_id"`
}

func (q *Queries) CreateGuild(ctx context.Context, arg CreateGuildParams) (Guild, error) {
	row := q.db.QueryRowContext(ctx, createGuild,
		arg.DiscordID,
		arg.Name,
		arg.Icon,
		arg.OwnerID,
	)
	var i Guild
	err := row.Scan(
		&i.ID,
		&i.DiscordID,
		&i.Name,
		&i.Icon,
		&i.OwnerID,
	)
	return i, err
}

const deleteGuild = `-- name: DeleteGuild :exec
DELETE
FROM guild
WHERE id = $1
`

func (q *Queries) DeleteGuild(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteGuild, id)
	return err
}

const getGuild = `-- name: GetGuild :one
SELECT id, discord_id, name, icon, owner_id
FROM guild
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetGuild(ctx context.Context, id int64) (Guild, error) {
	row := q.db.QueryRowContext(ctx, getGuild, id)
	var i Guild
	err := row.Scan(
		&i.ID,
		&i.DiscordID,
		&i.Name,
		&i.Icon,
		&i.OwnerID,
	)
	return i, err
}

const listGuilds = `-- name: ListGuilds :many
SELECT id, discord_id, name, icon, owner_id
FROM guild
ORDER BY id
OFFSET $1 LIMIT $2
`

type ListGuildsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (q *Queries) ListGuilds(ctx context.Context, arg ListGuildsParams) ([]Guild, error) {
	rows, err := q.db.QueryContext(ctx, listGuilds, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Guild
	for rows.Next() {
		var i Guild
		if err := rows.Scan(
			&i.ID,
			&i.DiscordID,
			&i.Name,
			&i.Icon,
			&i.OwnerID,
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