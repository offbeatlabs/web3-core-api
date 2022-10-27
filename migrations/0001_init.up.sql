create table "tokens" (
    "updated_at" datetime not null,
	"symbol" varchar not null,
	"name" varchar not null,
	"contract_address" varchar not null,
	"chain" varchar not null,
	"logo" text,
	"provider_token_id" varchar not null,
	"provider" varchar not null,
	"usd_price" double,
	"usd_market_cap" double,
	"usd_24hour_change" double,
	primary key ("provider_token_id", "provider")
);

create unique index "idx_unique_token_entry" on "tokens" ("contract_address","chain","provider");