-- +goose Up
-- +goose StatementBegin
CREATE TABLE "articles" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "title" varchar(140) NOT NULL,
  "description" varchar(1000) NOT NULL,
  "photos" varchar[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "articles";
-- +goose StatementEnd
