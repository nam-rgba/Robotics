CREATE TABLE "coach" (
  "coach_id" bigserial PRIMARY KEY,
  "fullname" VARCHAR,
  "email" VARCHAR UNIQUE,
  "country" VARCHAR,
  "title" varchar,
  "company" VARCHAR,
  "numberofcandidate" int,
  "password" VARCHAR(255) NOT NULL
);

CREATE TABLE "candidate" (
  "can_id" bigserial PRIMARY KEY,
  "fullname" VARCHAR,
  "title" varchar,
  "email" VARCHAR UNIQUE,
  "country" VARCHAR,
  "ranklocal" int,
  "rankworld" int,
  "company" VARCHAR,
  "dateofbirth" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "coach_id" bigserial REFERENCES coach(coach_id) ON DELETE CASCADE,
   "password" VARCHAR(255) NOT NULL
);

CREATE TABLE "organization" (
  "org_id" bigserial PRIMARY KEY,
  "name" VARCHAR,
  "logo" VARCHAR,
  "description" text
);

CREATE TABLE "competition" (
  -- basic
  "com_id" bigserial PRIMARY KEY,
  "name" VARCHAR,
  "decription" text,
  "rules" text,
  "agelimit" int,
  "images" text[],
  -- time
  "status" varchar(10) DEFAULT 'Comming',
  "stime" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "etime" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "registerstart" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "registerend" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  -- location
  "location" VARCHAR,
  "venue" VARCHAR,
  "address" VARCHAR,
  "city" VARCHAR,

  -- details
  "max_team_size" int DEFAULT 4,
  "participant_fee" int DEFAULT 0,
  "first_prize" int DEFAULT 0,
  "second_prize" int DEFAULT 0,
  "third_prize" int DEFAULT 0,

  -- organization
  "org_id" bigserial REFERENCES organization(org_id) ON DELETE CASCADE
);

CREATE TABLE "team" (
  "team_id" bigserial PRIMARY KEY,
  "coach_id" bigserial REFERENCES coach(coach_id) ON DELETE CASCADE,
  "join_code" VARCHAR UNIQUE,
  "teamname" VARCHAR,
  "competiton_id" int REFERENCES competition(com_id) ON DELETE CASCADE,
  "maxteam" int
);

CREATE TABLE "challenges" (
"chal_id" bigserial PRIMARY KEY,
  "com_id" int REFERENCES competition(com_id) ON DELETE CASCADE,
  "name" VARCHAR,
  "stime" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "etime" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "maxteam" int
);

CREATE TABLE "team_candidate"(
  "team_id" bigserial REFERENCES team(team_id) ON DELETE CASCADE,
  "can_id" bigserial REFERENCES candidate(can_id) ON DELETE CASCADE,
  "invitation_status" varchar(10) DEFAULT 'pending',
  PRIMARY KEY (team_id, can_id)
);




