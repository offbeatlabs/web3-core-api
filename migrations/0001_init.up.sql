create table "tokens" (
	"id" integer primary key,
	"updated_at" datetime not null,
	"symbol" varchar not null,
	"name" varchar not null,
	"logo" text,
	"source_token_id" varchar not null,
	"source" varchar not null,
	"usd_price" double not null default '0.0',
	"usd_market_cap" double not null default '0.0',
	"usd_24h_change" double not null default '0.0',
	"usd_24h_volume" double not null default '0.0'
);

create unique index "idx_uniq_token" on "tokens" ("source","source_token_id");
create index "idx_symbol" on "tokens" ("symbol");

create table "token_platforms" (
	"token_id" integer not null,
	"platform_name" varchar not null,
	"address" varchar not null,
	"decimal" integer
);

create unique index "idx_uniq" on "token_platforms" ("address","platform_name");