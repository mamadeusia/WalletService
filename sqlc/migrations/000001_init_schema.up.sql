CREATE TABLE wallet (
  "id" UUID PRIMARY KEY,
  "user_id" varchar(32) NOT NULL,
  "address" varchar(42) NOT NULL,
  "network" varchar(32) NOT NULL,
  "asset" varchar(32) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE INDEX ON "wallet" ("user_id");
CREATE TABLE internal_wallet (
  "wallet_id" UUID NOT NULL,
  "derivation_version" varchar(64) NOT NULL,
  "index" bigint NOT NULL
);
ALTER TABLE internal_wallet
ADD FOREIGN KEY (wallet_id) REFERENCES wallet (id);