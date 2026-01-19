package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"github.com/sjzar/chatlog/internal/wechatdb"
)

type MySQLMigration struct {
	DBPath   string
	DB       *sql.DB
	WeChatDB *wechatdb.DB
}

func NewMySQLMigration(dsn string, wechatDB *wechatdb.DB) (*MySQLMigration, error) {
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Error().Err(err).Msg("failed to open mysql database")
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		log.Error().Err(err).Msg("failed to ping mysql database")
		return nil, err
	}
	log.Info().Str("dsn", maskDSN(dsn)).Msg("mysql database connected")

	m := &MySQLMigration{
		DBPath:   dsn,
		DB:       sqlDB,
		WeChatDB: wechatDB,
	}
	if err := m.Init(); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MySQLMigration) Init() error {
	if err := m.OptimizeDB(); err != nil {
		log.Error().Err(err).Msg("failed to optimize database")
		return err
	}
	return nil
}

func (m *MySQLMigration) OptimizeDB() error {
	if m.DB == nil {
		return fmt.Errorf("DB is not initialized")
	}

	settings := []struct {
		query string
		desc  string
	}{
		{"SET SESSION bulk_insert_buffer_size = 256 * 1024 * 1024", "bulk insert buffer size"},
		{"SET SESSION innodb_autoinc_lock_mode = 2", "innodb autoinc lock mode"},
		{"SET SESSION innodb_flush_log_at_trx_commit = 2", "innodb flush log at trx commit"},
	}

	for _, s := range settings {
		if _, err := m.DB.Exec(s.query); err != nil {
			log.Error().Err(err).Msgf("failed to set %s", s.desc)
			return err
		}
	}

	return nil
}

func (m *MySQLMigration) Export() error {
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

func (m *MySQLMigration) ExportContact() error {
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
			VALUES (?, ?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE user_name = user_name
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

	log.Info().Ints("count", []int{successCount, failCount}).Msg("contacts saved to mysql")
	return nil
}

func (m *MySQLMigration) ExportChatRoom() error {
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
				INSERT INTO chat_room (name, owner, remark, nick_name, users)
				VALUES (?, ?, ?, ?, ?)
				ON DUPLICATE KEY UPDATE name = name
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
		_, err := stmt.ExecContext(ctx, room.Name, room.Owner, room.Remark, room.NickName, room.Users)
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

	log.Info().Ints("count", []int{successCount, failCount}).Msg("chat rooms saved to mysql")
	return nil
}

func (m *MySQLMigration) ExportMessage() error {
	insertStmt, err := m.DB.Prepare(`
		INSERT INTO wechat_message (seq, time, talker, talker_name, is_chat_room, sender, is_self, type, sub_type, content, contents)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE seq = seq
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
			var contentsJSON []byte
			if message.Contents != nil {
				contentsJSON, err = json.Marshal(message.Contents)
				if err != nil {
					log.Error().Err(err).Msg("failed to marshal contents")
					tx.Rollback()
					return err
				}
			}

			_, err := insertStmt.Exec(
				message.Seq,
				message.Time,
				message.Talker,
				message.TalkerName,
				message.IsChatRoom,
				message.Sender,
				message.IsSelf,
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
		log.Info().Int("count", len(messages)).Msg("messages saved to mysql")
	}

	log.Info().Ints("stats", []int{len(contacts.Items), totalMessages}).Msg("import completed")
	return nil
}

func (m *MySQLMigration) Close() error {
	return m.DB.Close()
}
