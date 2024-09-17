CREATE TABLE coach (
  "coach_id" bigserial PRIMARY KEY,
  "fullname" VARCHAR,
  "email" VARCHAR UNIQUE,
  "country" VARCHAR,
  "title" varchar,
  "company" VARCHAR
);

CREATE TABLE candicate (
  "can_id" bigserial PRIMARY KEY,
  "fullname" VARCHAR,
  "title" varchar,
  "email" VARCHAR UNIQUE,
  "country" VARCHAR,
  "ranklocal" int,
  "rankworld" int,
  "company" VARCHAR,
  "dateofbirth" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "coach_id" int
);

CREATE TABLE team (
  "team_id" bigserial PRIMARY KEY,
  "coach_id" int,
  "member_id" int,
  "teamname" VARCHAR,
  "competiton_id" int
);

CREATE TABLE competition (
  "com_id" bigserial PRIMARY KEY,
  "name" VARCHAR,
  "stime" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "etime" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "decription" text,
  "rules" text,
  "agelimit" int,
  "prizevalue" VARCHAR
);

CREATE TABLE challenges (
"chal_id" bigserial PRIMARY KEY,
  "com_id" int,
  "name" VARCHAR,
  "stime" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "etime" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "maxteam" int
);


