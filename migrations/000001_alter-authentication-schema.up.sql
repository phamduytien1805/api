CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" VARCHAR(32) NOT NULL UNIQUE,
  "email" VARCHAR(254) NOT NULL UNIQUE,
  "email_verified" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "state" SMALLINT NOT NULL,
  "stated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_credentials" (
  "user_id" uuid NOT NULL,
  "credential" varchar NOT NULL,
  "salt" varchar NOT NULL,
);

CREATE TABLE "user_social_tokens" (
  "user_id" uuid,
  "provider_type" varchar(50),
  "access_token" varchar,
  "refresh_token" varchar,
  "expires" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY (user_id, provider_type)
);

ALTER TABLE "user_social_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_credentials" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

COMMENT ON COLUMN "users"."state" IS 'enum defined in code';

COMMENT ON COLUMN "user_social_tokens"."provider_type" IS 'enum defined in code';
