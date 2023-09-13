CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "todos" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "todo_name" varchar NOT NULL,
  "description" text NOT NULL,
  "is_completed" boolean NOT NULL DEFAULT false,
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "todos" ("username");

ALTER TABLE "todos" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");