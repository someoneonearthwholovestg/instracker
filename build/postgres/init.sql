-- TABLE "users" --
CREATE TABLE IF NOT EXISTS "users" (
  "id"                INT PRIMARY KEY,
  "username"          TEXT,
  "firstname"         TEXT,
  "lastname"          TEXT,
  "state"             INT DEFAULT 0 NOT NULL,
  "registration_date" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "instagram_users" (
  "username"  TEXT PRIMARY KEY,
  "followers" INT,
  "following" INT
);

CREATE TABLE IF NOT EXISTS "subscriptions" (
  "id"             INT PRIMARY KEY
  "user_id"        INT REFERENCES "users" ("id"),
  "insta_username" TEXT NOT NULL REFERENCES "instagram_users" ("username")
);

CREATE TYPE "GROUP_TYPE" AS ENUM('following', 'followers');

CREATE TABLE IF NOT EXISTS "following_followers" (
  "username"       TEXT,
  "fullname"       TEXT,
  "URL"            TEXT,
  "refer_username" TEXT NOT NULL REFERENCES "instagram_users" ("username"),
  "group_type"     GROUP_TYPE
);