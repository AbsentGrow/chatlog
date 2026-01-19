package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sjzar/chatlog/internal/wechatdb"
)

type SQLiteMigration struct {
	DBPath   string
	DB       *sql.DB
	WeChatDB *wechatdb.DB
}

func NewSQLiteMigration(sqlitePath string, wechatDB *wechatdb.DB) (*SQLiteMigration, error) {
	sqlDB, err := sql.Open("sqlite3", sqlitePath)
	if err != nil {
		log.Error().Err(err).Msg("failed to open sqlite database")
		return nil, err
	}
	log.Info().Str("path", sqlitePath).Msg("sqlite database opened")

	m := &SQLiteMigration{
		DBPath:   sqlitePath,
		DB:       sqlDB,
		WeChatDB: wechatDB,
	}
	if err := m.Init(); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *SQLiteMigration) Init() error {
	if err := m.OptimizeDB(); err != nil {
		log.Error().Err(err).Msg("failed to optimize database")
		return err
	}
	return nil
}

func (m *SQLiteMigration) OptimizeDB() error {
	if m.DB == nil {
		return fmt.Errorf("DB is not initialized")
	}

	pragmas := []struct {
		query string
		desc  string
	}{
		{"PRAGMA journal_mode=WAL", "WAL mode"},
		{"PRAGMA synchronous=NORMAL", "synchronous mode"},
		{"PRAGMA cache_size=-10000000", "cache size"},
		{"PRAGMA mmap_size=268435456", "mmap size"},
		{"PRAGMA temp_store=MEMORY", "temp store"},
		{"PRAGMA busy_timeout=30000", "busy timeout"},
	}

	for _, p := range pragmas {
		if _, err := m.DB.Exec(p.query); err != nil {
			log.Error().Err(err).Msgf("failed to set %s", p.desc)
			return err
		}
	}

	return nil
}

func (m *SQLiteMigration) Export() error {
	if err := m.ExportContact(); err != nil {
		log.Error().Err(err).Msg("failed to export contact")
		return err
	}
	if err := m.ExportChatRoom(); err != nil {
		log.Error().Err(err).Msg("failed to export chat room")
		return err
	}
	if err := m.ExportMessage(); err != nil {
		log.Error().Err(err).Msg("failed to export message")
		return err
	}
	return nil
}

func (m *SQLiteMigration) ExportContact() error {
	contacts, err := m.WeChatDB.GetContacts("", 0, 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get contacts")
		return err
	}

	ctx := context.Background()

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to begin transaction")
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
			INSERT INTO wechat_contact (user_name, alias, remark, nick_name, is_friend)
			VALUES ($1, $2, $3, $4, $5)
		`)
	if err != nil {
		log.Error().Err(err).Msg("failed to prepare statement")
		return err
	}
	defer stmt.Close()

	successCount := 0
	failCount := 0
	for _, contact := range contacts.Items {
		log.Info().Interface("contact", contact).Msg("contact")
		_, err := stmt.ExecContext(ctx, contact.UserName, contact.Alias, contact.Remark, contact.NickName, contact.IsFriend)
		if err != nil {
			log.Error().Err(err).Str("user_name", contact.UserName).Msg("failed to insert contact")
			failCount++
			continue
		}
		successCount++
	}

	if err := tx.Commit(); err != nil {
		log.Error().Err(err).Msg("failed to commit transaction")
		return err
	}

	log.Info().Ints("count", []int{successCount, failCount}).Msg("contacts saved to sqlite")
	return nil
}

func (m *SQLiteMigration) ExportChatRoom() error {
	rooms, err := m.WeChatDB.GetChatRooms("", 0, 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get rooms")
		return err
	}

	ctx := context.Background()

	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to begin transaction")
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
				INSERT INTO wechat_chat_room (name, owner, remark, nick_name, users)
				VALUES ($1, $2, $3, $4, $5)
			`)
	if err != nil {
		log.Error().Err(err).Msg("failed to prepare statement")
		return err
	}
	defer stmt.Close()

	successCount := 0
	failCount := 0
	for _, room := range rooms.Items {
		log.Info().Interface("room", room).Msg("room")
		usersJSON, err := json.Marshal(room.Users)
		if err != nil {
			log.Error().Err(err).Str("chat_room", room.Name).Msg("failed to marshal users")
			failCount++
			continue
		}
		_, err = stmt.ExecContext(ctx, room.Name, room.Owner, room.Remark, room.NickName, string(usersJSON))
		if err != nil {
			log.Error().Err(err).Str("chat_room", room.Name).Msg("failed to insert chat room")
			failCount++
			continue
		}
		successCount++
	}

	if err := tx.Commit(); err != nil {
		log.Error().Err(err).Msg("failed to commit transaction")
		return err
	}

	log.Info().Ints("count", []int{successCount, failCount}).Msg("chat rooms saved to sqlite")
	return nil
}

func (m *SQLiteMigration) ExportMessage() error {
	insertStmt, err := m.DB.Prepare(`
		INSERT INTO wechat_message (seq, time, talker, talker_name, is_chat_room, sender, is_self, type, sub_type, content, contents)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Error().Err(err).Msg("failed to prepare insert statement")
		return err
	}
	defer insertStmt.Close()

	totalMessages := 0

	contacts, err := m.WeChatDB.GetContacts("", 0, 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to get contacts")
		return err
	}

	for _, contact := range contacts.Items {
		log.Info().Interface("contact", contact).Msg("contact")
		messages, err := m.WeChatDB.GetMessages(time.Unix(0, 0), time.Now(), contact.UserName, "", "", 0, 0)
		if err != nil {
			log.Error().Err(err).Msg("failed to get messages")
			return err
		}

		if len(messages) == 0 {
			continue
		}

		log.Info().Int("batch_size", len(messages)).Msg("batch importing messages")

		ctx := context.Background()
		tx, err := m.DB.BeginTx(ctx, nil)
		if err != nil {
			log.Error().Err(err).Msg("failed to begin transaction")
			return err
		}

		for _, message := range messages {
			var contentsJSON string
			if message.Contents != nil {
				contentsBytes, err := json.Marshal(message.Contents)
				if err != nil {
					log.Error().Err(err).Msg("failed to marshal contents")
					tx.Rollback()
					return err
				}
				contentsJSON = string(contentsBytes)
			}

			_, err := insertStmt.Exec(
				message.Seq,
				message.Time,
				message.Talker,
				message.TalkerName,
				boolToInt64(message.IsChatRoom),
				message.Sender,
				boolToInt64(message.IsSelf),
				message.Type,
				message.SubType,
				message.Content,
				contentsJSON,
			)
			if err != nil {
				log.Error().Err(err).Msg("failed to insert message")
				tx.Rollback()
				return err
			}
		}

		if err := tx.Commit(); err != nil {
			log.Error().Err(err).Msg("failed to commit transaction")
			return err
		}

		totalMessages += len(messages)
		log.Info().Int("count", len(messages)).Msg("messages saved to sqlite")
	}

	log.Info().Ints("stats", []int{len(contacts.Items), totalMessages}).Msg("import completed")
	return nil
}
