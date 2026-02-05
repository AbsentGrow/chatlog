package backup

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sjzar/chatlog/internal/wechatdb"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DriverType acts as an enum for supported database drivers
type DriverType string

const (
	DriverSQLite   DriverType = "sqlite"
	DriverMySQL    DriverType = "mysql"
	DriverPostgres DriverType = "postgres"
)

// Config holds configuration for the backup service
type Config struct {
	Driver DriverType
	DSN    string // Data Source Name (connection string)
}

// Service handles the backup logic
type Service struct {
	db       *gorm.DB
	wechatDB *wechatdb.DB
}

// NewService creates a new backup service with the stored GORM connection
func NewService(cfg Config, wechatDB *wechatdb.DB) (*Service, error) {
	var dialector gorm.Dialector

	switch cfg.Driver {
	case DriverSQLite:
		dialector = sqlite.Open(cfg.DSN)
	case DriverMySQL:
		dialector = mysql.Open(cfg.DSN)
	case DriverPostgres:
		dialector = postgres.Open(cfg.DSN)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}

	// Use a logger that writes to zerolog or stdout, adjusting log level as needed
	gormLogger := logger.Default.LogMode(logger.Warn)

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// AutoMigrate the schema
	if err := db.AutoMigrate(&Contact{}, &ChatRoom{}, &Message{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database schema: %w", err)
	}

	// Optimization for SQLite
	if cfg.Driver == DriverSQLite {
		if err := optimizeSQLite(db); err != nil {
			log.Warn().Err(err).Msg("failed to apply SQLite optimizations")
		}
	}

	return &Service{
		db:       db,
		wechatDB: wechatDB,
	}, nil
}

// optimizeSQLite applies performance pragmas for SQLite
func optimizeSQLite(db *gorm.DB) error {
	pragmas := []string{
		"PRAGMA journal_mode=WAL",
		"PRAGMA synchronous=NORMAL",
		"PRAGMA cache_size=-10000000", // ~10GB cache if memory allows, adjust as needed
		"PRAGMA temp_store=MEMORY",
		"PRAGMA busy_timeout=30000",
	}

	for _, p := range pragmas {
		if err := db.Exec(p).Error; err != nil {
			return err
		}
	}
	return nil
}

// Run performs the full backup
func (s *Service) Run() error {
	if err := s.BackupContacts(); err != nil {
		log.Error().Err(err).Msg("failed to backup contacts")
		return err
	}
	if err := s.BackupChatRooms(); err != nil {
		log.Error().Err(err).Msg("failed to backup chat rooms")
		return err
	}
	if err := s.BackupMessages(); err != nil {
		log.Error().Err(err).Msg("failed to backup messages")
		return err
	}
	return nil
}

// BackupContacts fetches contacts from WeChatDB and saves them via GORM
func (s *Service) BackupContacts() error {
	log.Info().Msg("Starting contacts backup...")
	contacts, err := s.wechatDB.GetContacts("", 0, 0)
	if err != nil {
		return fmt.Errorf("failed to get contacts from source: %w", err)
	}

	// Helper to convert bool to int
	boolToInt := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}

	var contactModels []Contact
	for _, c := range contacts.Items {
		contactModels = append(contactModels, Contact{
			UserName: c.UserName,
			Alias:    c.Alias,
			Remark:   c.Remark,
			NickName: c.NickName,
			IsFriend: boolToInt(c.IsFriend), // Convert bool to int
		})
	}

	if len(contactModels) == 0 {
		log.Info().Msg("No contacts to backup.")
		return nil
	}

	// Batch insert/upsert
	result := s.db.CreateInBatches(&contactModels, 100)

	if result.Error != nil {
		return fmt.Errorf("failed to save contacts: %w", result.Error)
	}

	log.Info().Int("count", len(contactModels)).Msg("Contacts backup completed.")
	return nil
}

// BackupChatRooms fetches chat rooms from WeChatDB and saves them via GORM
func (s *Service) BackupChatRooms() error {
	log.Info().Msg("Starting chat rooms backup...")
	rooms, err := s.wechatDB.GetChatRooms("", 0, 0)
	if err != nil {
		return fmt.Errorf("failed to get chat rooms from source: %w", err)
	}

	var roomModels []ChatRoom
	for _, r := range rooms.Items {
		// Convert []model.ChatRoomUser to []string for JSONStringList
		var userNames []string
		for _, u := range r.Users {
			userNames = append(userNames, u.UserName)
		}

		roomModels = append(roomModels, ChatRoom{
			Name:     r.Name,
			Owner:    r.Owner,
			Remark:   r.Remark,
			NickName: r.NickName,
			Users:    JSONStringList(userNames), // Cast []string to JSONStringList
		})
	}

	if len(roomModels) == 0 {
		log.Info().Msg("No chat rooms to backup.")
		return nil
	}

	result := s.db.CreateInBatches(&roomModels, 100)
	if result.Error != nil {
		return fmt.Errorf("failed to save chat rooms: %w", result.Error)
	}

	log.Info().Int("count", len(roomModels)).Msg("Chat rooms backup completed.")
	return nil
}

// BackupMessages fetches and saves messages for all contacts
func (s *Service) BackupMessages() error {
	log.Info().Msg("Starting messages backup...")

	contacts, err := s.wechatDB.GetContacts("", 0, 0)
	if err != nil {
		return fmt.Errorf("failed to get contacts for message backup: %w", err)
	}

	totalMessages := 0

	for _, contact := range contacts.Items {
		// Log progress for specific contacts if needed
		// log.Debug().Str("contact", contact.UserName).Msg("Processing messages for contact")

		// Fetch all messages for this contact (time range 0 to Now)
		msgs, err := s.wechatDB.GetMessages(time.Unix(0, 0), time.Now(), contact.UserName, "", "", 0, 0)
		if err != nil {
			log.Error().Err(err).Str("contact", contact.UserName).Msg("failed to get messages")
			continue
		}

		if len(msgs) == 0 {
			continue
		}

		var msgModels []Message
		for _, m := range msgs {
			var contents JSONContent
			if m.Contents != nil {
				// Convert map[string]interface{} to JSONContent
				contents = JSONContent(m.Contents) // wechatdb.Message.Contents is likely map[string]any or struct
			}

			// Helper to convert bool to int
			boolToInt := func(b bool) int {
				if b {
					return 1
				}
				return 0
			}

			msgModels = append(msgModels, Message{
				Seq:        uint64(m.Seq), // Convert int64 to uint64
				Time:       m.Time,
				Talker:     m.Talker,
				TalkerName: m.TalkerName,
				IsChatRoom: boolToInt(m.IsChatRoom),
				Sender:     m.Sender,
				IsSelf:     boolToInt(m.IsSelf),
				Type:       int(m.Type),    // Convert int64 to int
				SubType:    int(m.SubType), // Convert int64 to int
				Content:    m.Content,
				Contents:   contents,
			})
		}

		// Batch insert
		if err := s.db.CreateInBatches(msgModels, 100).Error; err != nil {
			log.Error().Err(err).Str("contact", contact.UserName).Msg("failed to save batch messages")
			continue
		}

		totalMessages += len(msgModels)
		log.Info().Int("count", len(msgModels)).Str("contact", contact.NickName).Msg("Saved messages")

		// Optional: SQLite checkpointing if needed for large imports, but WAL handle it mostly
	}

	log.Info().Int("total_messages", totalMessages).Msg("Messages backup completed.")
	return nil
}
