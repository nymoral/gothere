CREATE TABLE recover
(
  pk        BIGSERIAL PRIMARY KEY,
  key       VARCHAR(128),
  user_pk   BIGSERIAL REFERENCES users (pk),
  created   TIMESTAMP NOT NULL DEFAULT NOW()
);
