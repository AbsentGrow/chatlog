-- SQLite Database Initialization Script
-- Database: chatlog
-- Function: Storage for WeChat chat history, contacts, and group chat information

PRAGMA foreign_keys = ON;

-- =====================================================
-- Part 1: Contact Table
-- =====================================================
CREATE TABLE IF NOT EXISTS wechat_contact (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      user_name TEXT NOT NULL UNIQUE,     -- WeChat User ID (Unique Identifier)
      alias TEXT,                          -- WeChat Alias (WeChat ID)
      remark TEXT,                         -- Remark Name
      nick_name TEXT,                      -- Nickname
      is_friend INTEGER DEFAULT 0,         -- is friend: 0-No, 1-Yes
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Contact Table Indexes
CREATE INDEX IF NOT EXISTS idx_wechat_contact_nick_name ON wechat_contact(nick_name);
CREATE INDEX IF NOT EXISTS idx_wechat_contact_remark ON wechat_contact(remark);
CREATE INDEX IF NOT EXISTS idx_wechat_contact_is_friend ON wechat_contact(is_friend);

-- =====================================================
-- Part 2: Chat Room Table
-- =====================================================
CREATE TABLE IF NOT EXISTS wechat_chat_room (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,                  -- Chat Room Name
      owner TEXT NOT NULL DEFAULT '',       -- Owner user_name
      remark TEXT NOT NULL DEFAULT '',      -- Chat Room Remark
      nick_name TEXT NOT NULL DEFAULT '',   -- Chat Room Nickname
      users TEXT DEFAULT '[]',              -- Members JSON Array
      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Chat Room Table Indexes
CREATE INDEX IF NOT EXISTS idx_wechat_chat_room_name ON wechat_chat_room(name);
CREATE INDEX IF NOT EXISTS idx_wechat_chat_room_owner ON wechat_chat_room(owner);

-- =====================================================
-- Part 3: Message Table (Core Table)
-- =====================================================
CREATE TABLE IF NOT EXISTS wechat_message (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      seq INTEGER NOT NULL,                -- Message Sequence
      time TIMESTAMP NOT NULL,             -- Message Time
      talker TEXT NOT NULL,                -- Conversation ID (User ID or Chat Room ID)
      talker_name TEXT NOT NULL DEFAULT '',-- Conversation Name
      is_chat_room INTEGER NOT NULL DEFAULT 0,  -- Is Chat Room: 0-Private Chat, 1-Group Chat
      sender TEXT NOT NULL,                -- Sender user_name
      sender_name TEXT NOT NULL DEFAULT '',-- Sender Name
      is_self INTEGER NOT NULL DEFAULT 0,  -- Is Sent By Self: 0-No, 1-Yes
      type INTEGER NOT NULL DEFAULT 0,     -- Message Type
      sub_type INTEGER NOT NULL DEFAULT 0, -- Sub Message Type
      content TEXT NOT NULL DEFAULT '',    -- Message Content
      contents TEXT NOT NULL DEFAULT '{}', -- Extended Content (JSON)
      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Message Table Indexes (Optimized for query scenarios)
CREATE INDEX IF NOT EXISTS idx_wechat_message_talker ON wechat_message(talker);              -- Query by conversation
CREATE INDEX IF NOT EXISTS idx_wechat_message_time ON wechat_message(time DESC);             -- Order by time
CREATE INDEX IF NOT EXISTS idx_wechat_message_talker_time ON wechat_message(talker, time DESC); -- By conversation + time (Most used)
CREATE INDEX IF NOT EXISTS idx_wechat_message_is_chat_room ON wechat_message(is_chat_room, talker); -- Filter Group/Private Chat
CREATE INDEX IF NOT EXISTS idx_wechat_message_sender ON wechat_message(sender);              -- Query by sender

-- =====================================================
-- Part 4: Full Text Search (FTS5)
-- For fast search of message content
-- =====================================================
CREATE VIRTUAL TABLE IF NOT EXISTS wechat_message_fts USING fts5(
      content,                             -- Search Content Column
      content='wechat_message',            -- Associated Source Table
      content_rowid='id',                  -- Source Table Primary Key Mapping
      tokenize='unicode61'                 -- Use unicode61 tokenizer for mixed Chinese/English support
);

-- Trigger: Sync to FTS table on new message insertion
CREATE TRIGGER IF NOT EXISTS wechat_message_ai AFTER INSERT ON wechat_message BEGIN
      INSERT INTO wechat_message_fts(rowid, content) VALUES (new.id, new.content);
END;

-- Trigger: Remove from FTS table on message deletion
CREATE TRIGGER IF NOT EXISTS wechat_message_ad AFTER DELETE ON wechat_message BEGIN
      INSERT INTO wechat_message_fts(wechat_message_fts, rowid, content) VALUES('delete', old.id, old.content);
END;

-- Trigger: Sync FTS table on message update
CREATE TRIGGER IF NOT EXISTS wechat_message_au AFTER UPDATE ON wechat_message BEGIN
      INSERT INTO wechat_message_fts(wechat_message_fts, rowid, content) VALUES('delete', old.id, old.content);
      INSERT INTO wechat_message_fts(rowid, content) VALUES (new.id, new.content);
END;

-- Migration: Import existing messages into FTS table
-- Note: Execute on first initialization, no need to repeat
INSERT INTO wechat_message_fts(rowid, content)
SELECT id, content FROM wechat_message WHERE content != '';

-- =====================================================
-- Part 5: Performance Optimization Configuration (Optional)
-- Recommended to be set at application layer, below are reference configs
-- =====================================================
-- PRAGMA journal_mode = WAL;            -- Enable WAL mode to improve concurrent write
-- PRAGMA synchronous = NORMAL;          -- Synchronous mode: Balance performance and safety
-- PRAGMA cache_size = -64000;           -- Cache size: 64MB
-- PRAGMA temp_store = MEMORY;           -- Store temporary tables in memory
-- PRAGMA mmap_size = 268435456;         -- Memory map size: 256MB
