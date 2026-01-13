CREATE TABLE IF NOT EXISTS contact
(
    id         SERIAL PRIMARY KEY,
    user_name  VARCHAR(255) NOT NULL UNIQUE,
    alias      VARCHAR(255),
    remark     VARCHAR(255),
    nick_name  VARCHAR(255),
    is_friend  BOOLEAN   DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );


DROP TABLE "chat_room";
CREATE TABLE IF NOT EXISTS "chat_room"
(
    id         SERIAL PRIMARY KEY,
    name       TEXT            NOT NULL,
    owner      TEXT            NOT NULL DEFAULT '',
    remark     TEXT            NOT NULL DEFAULT '',
    nick_name  TEXT            NOT NULL DEFAULT '',
    users      JSON DEFAULT '[]',
    created_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE "message";
CREATE TABLE IF NOT EXISTS "message"
(
    id             SERIAL PRIMARY KEY,
    "seq"          BIGINT    NOT NULL,
    "time"         TIMESTAMP NOT NULL,
    "talker"       TEXT      NOT NULL,
    "talker_name"  TEXT      NOT NULL DEFAULT '',
    "is_chat_room" BIGINT    NOT NULL DEFAULT 0,
    "sender"       TEXT      NOT NULL,
    "sender_name"  TEXT      NOT NULL DEFAULT '',
    "is_self"      BIGINT    NOT NULL DEFAULT 0,
    "type"         BIGINT    NOT NULL DEFAULT 0,
    "sub_type"     BIGINT    NOT NULL DEFAULT 0,
    "content"      TEXT      NOT NULL DEFAULT '',
    "contents"     JSON      NOT NULL DEFAULT '{}',
    "created_at"   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
