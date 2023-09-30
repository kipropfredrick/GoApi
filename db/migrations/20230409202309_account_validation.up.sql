CREATE TABLE "account_validation" (
    "id" bigserial PRIMARY KEY,
    "hashed_code" text NOT NULL,
    "user_id" bigserial NOT NULL
)