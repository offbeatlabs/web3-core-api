create table "tokens" (
    "updatedAt" datetime not null,
	"symbol" varchar not null,
	"name" varchar not null,
	"contractAddress" varchar not null,
	"chain" varchar not null,
	"logo" text,
	"providerTokenId" varchar not null,
	"provider" varchar not null,
	"usdPrice" double,
	"usdMarketCap" double,
	"usd24hourChange" double,
	primary key ("providerTokenId", "provider")
);

create unique index "idx_unique_token_entry" on "tokens" ("contractAddress","chain","provider");