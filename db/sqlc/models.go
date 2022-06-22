// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"time"
)

type Admin struct {
	ID            int64     `json:"id"`
	DiscordID     string    `json:"discord_id"`
	Username      string    `json:"username"`
	Discriminator string    `json:"discriminator"`
	Verified      bool      `json:"verified"`
	Email         string    `json:"email"`
	Avatar        string    `json:"avatar"`
	Flags         int64     `json:"flags"`
	Banner        string    `json:"banner"`
	AccentColor   int64     `json:"accent_color"`
	PublicFlags   int64     `json:"public_flags"`
	CreatedAt     time.Time `json:"created_at"`
}

type Config struct {
	ID        int64     `json:"id"`
	Version   string    `json:"version"`
	Filename  string    `json:"filename"`
	CreatedAt time.Time `json:"created_at"`
	GuildID   int64     `json:"guild_id"`
}

type Guild struct {
	ID        int64  `json:"id"`
	DiscordID string `json:"discord_id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	OwnerID   string `json:"owner_id"`
}

type GuildAdmin struct {
	GuildID int64 `json:"guild_id"`
	AdminID int64 `json:"admin_id"`
}
