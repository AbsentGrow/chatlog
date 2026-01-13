package main

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sjzar/chatlog/internal/chatlog"
	"github.com/sjzar/chatlog/internal/wechatdb"

	_ "github.com/joho/godotenv/autoload"
)

func initLog(debug bool, miscDir string) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logPath := filepath.Join(miscDir, "logs")
	logFD, err := os.OpenFile(filepath.Join(logPath, "wechat-mem0-core.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	log.Logger = log.Output(zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true, TimeFormat: time.RFC3339}, zerolog.ConsoleWriter{Out: logFD, NoColor: true, TimeFormat: time.RFC3339}))
}

func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func main() {
	//miscDir := os.Getenv("MISC_DIR")
	//initLog(false, miscDir)

	// debug 场景没法突破内存限制，用单独的命令行可以解析出来
	m := chatlog.New(chatlog.ManagerTypeGRPC)
	m = m
	/*
		// fetch key for decrypt WeChat
		keyData, err := m.GetKey("", 0, false, false)
		if err != nil {
			if errors.Is(err, errors.ErrSIPEnabled) {
				log.Info().Msg("SIP is enabled")
				return
			}
			log.Err(err).Msgf("failed to get key")
			return
		}
		fmt.Println(keyData)
	*/
	/*
		// WeChat management
		instances := m.GetWeChatInstances()
		if len(instances) > 0 {
			instance := instances[0]
			err := instance.RefreshStatus()
			if err != nil {
				log.Err(err).Msgf("failed to refresh status")
				return
			}
			fmt.Println(instance)
		}
	*/
	/*
		// decrypt WeChat data manually
		cmdConf := make(map[string]any)
		cmdConf["platform"] = os.Getenv("PLATFORM")
		cmdConf["version"] = os.Getenv("VERSION")
		cmdConf["img_key"] = os.Getenv("IMG_KEY")
		cmdConf["data_key"] = os.Getenv("DATA_KEY")
		cmdConf["data_dir"] = os.Getenv("DATA_DIR")
		cmdConf["work_dir"] = os.Getenv("WORK_DIR")
		err := m.Decrypt("", cmdConf)
		if err != nil {
			log.Err(err).Msg("failed to decrypt")
			return
		}
		log.Info().Msg("decrypt success")
	*/

	/*
			// start http server
		cmdConf := make(map[string]any)
		cmdConf["platform"] = os.Getenv("PLATFORM")
		cmdConf["version"] = os.Getenv("VERSION")
		cmdConf["img_key"] = os.Getenv("IMG_KEY")
		cmdConf["data_key"] = os.Getenv("DATA_KEY")
		cmdConf["data_dir"] = os.Getenv("DATA_DIR")
		cmdConf["work_dir"] = os.Getenv("WORK_DIR")
		cmdConf["http_addr"] = os.Getenv("HTTP_ADDR")

		log.Info().Msgf("server cmd config: %+v", cmdConf)

		if err := m.CommandHTTPServer("", cmdConf); err != nil {
			log.Err(err).Msg("failed to start server")
			return
		}
	*/
	/*
		// db usage example
		workDir := os.Getenv("WORK_DIR")
		platform := os.Getenv("PLATFORM")
		version, _ := strconv.Atoi(os.Getenv("VERSION"))
		db, err := wechatdb.New(workDir, platform, version)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create db")
			return
		}
		db = db

		// 1. list contacts method
		contacts, err := db.GetContacts("", 0, 0)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get contacts")
			return
		}
		for _, contact := range contacts.Items {
			log.Info().Interface("contact", contact).Msg("contact")
		}

		// 2. chat rooms method
		rooms, err := db.GetChatRooms("", 0, 0)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get rooms")
			return
		}
		for _, room := range rooms.Items {
			log.Info().Interface("room", room).Msg("room")
		}

		// 3. chat session
		sessions, err := db.GetSessions("", 0, 0)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get sessions")
			return
		}
		for _, session := range sessions.Items {
			log.Info().Interface("session", session).Msg("session")
		}

		// 4. chat messages
		messages, err := db.GetMessages(time.Unix(0, 0), time.Now(), "何欢", "", "", 0, 0)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get messages")
			return
		}
		for _, message := range messages {
			log.Info().Interface("message", message.PlainTextContent()).Msg("message")
		}
	*/

	/*
		// WeChat Media file
		mediaPath := os.Getenv("IMAGE_MEDIA_FILE_PATH")
		b, err := os.ReadFile(mediaPath)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to read file")
			return
		}
		out, ext, err := dat2img.Dat2Image(b)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to read file")
			return
		}
		log.Info().Str("ext", ext).Msg("ext")

		outputPath := os.Getenv("IMAGE_MEDIA_FILE_OUTPUT_PATH") + "wechat_image." + ext
		if err := os.WriteFile(outputPath, out, 0644); err != nil {
			log.Fatal().Err(err).Msg("failed to write file")
			return
		}
		log.Info().Str("path", outputPath).Msg("image saved")
	*/

	workDir := os.Getenv("WORK_DIR")
	platform := os.Getenv("PLATFORM")
	version, _ := strconv.Atoi(os.Getenv("VERSION"))
	db, err := wechatdb.New(workDir, platform, version)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create db")
		return
	}

	ds := os.Getenv("DATA_SOURCE")
	pgDB, err := sql.Open("pgx", ds)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to postgres")
		return
	}
	defer pgDB.Close()

	/*
		contacts, err := db.GetContacts("", 0, 0)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get contacts")
			return
		}

		ctx := context.Background()

		stmt, err := pgDB.PrepareContext(ctx, `
			INSERT INTO contact (user_name, alias, remark, nick_name, is_friend)
			VALUES ($1, $2, $3, $4, $5)
		`)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to prepare statement")
			return
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

		log.Info().Ints("count", []int{successCount, failCount}).Msg("contacts saved to postgres")
	*/
	/*
		rooms, err := db.GetChatRooms("", 0, 0)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get rooms")
			return
		}
		ctx := context.Background()

		stmt, err := pgDB.PrepareContext(ctx, `
				INSERT INTO chat_room (name, owner, remark, nick_name, users)
				VALUES ($1, $2, $3, $4, $5)
			`)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to prepare statement")
			return
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
		log.Info().Ints("count", []int{successCount, failCount}).Msg("chat rooms saved to postgres")
	*/
	ctx := context.Background()

	stmt, err := pgDB.PrepareContext(ctx, `
			INSERT INTO message (seq, time, talker, talker_name, is_chat_room, sender, is_self, type, sub_type, content, contents)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		`)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to prepare statement")
		return
	}
	defer stmt.Close()

	contacts, err := db.GetContacts("", 0, 0)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get contacts")
		return
	}

	for _, contact := range contacts.Items {
		log.Info().Interface("contact", contact).Msg("contact")
		successCount := 0
		failCount := 0
		messages, err := db.GetMessages(time.Unix(0, 0), time.Now(), contact.UserName, "", "", 0, 0)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get messages")
			return
		}
		for _, message := range messages {
			log.Info().Interface("message", message.PlainTextContent()).Msg("message")
			_, err := stmt.ExecContext(ctx, message.Seq, message.Time, message.Talker, message.TalkerName, boolToInt64(message.IsChatRoom), message.Sender, boolToInt64(message.IsSelf), message.Type, message.SubType, message.Content, message.Contents)
			if err != nil {
				log.Error().Err(err).Str("message", message.Content).Msg("failed to insert message")
				failCount++
				continue
			}
			successCount++
		}
		log.Info().Ints("count", []int{successCount, failCount}).Msg("chat rooms saved to postgres")
	}

	/*
		// db migration
		dbFiles, err := pgmigrate.ListDBFiles(os.Getenv("DB_PATH"))
		if err != nil {
			log.Err(err).Msg("failed to list db files")
			return
		}

		for _, dbFile := range dbFiles {
			log.Info().Str("file", dbFile).Msg("loading db file")
		}
	*/

	/*
		// db exporter
		ddl, err := pgmigrate.GenerateDDL(os.Getenv("DB_PATH"))
		if err != nil {
			log.Err(err).Msg("failed to generate ddl")
			return
		}
		log.Info().Str("ddl", ddl).Msg("generated ddl")
	*/
}
