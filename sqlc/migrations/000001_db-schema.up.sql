CREATE TABLE movies (
"id" bigserial PRIMARY KEY,
"title" varchar NOT NULL,
"description" varchar NOT NULL,
"director" varchar NOT NULL,
"created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE ratings (
"id" bigserial PRIMARY KEY,
"movie_id" bigint NOT NULL,
"rating" numeric(3,1) CHECK(rating > 0 AND rating < 10),
"created_at" timestamptz NOT NULL DEFAULT (now())
);
