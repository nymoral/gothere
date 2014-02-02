CREATE TABLE guesses
(
  pk        BIGSERIAL PRIMARY KEY,

  game_pk   BIGSERIAL REFERENCES games (pk),
  user_pk   BIGSERIAL REFERENCES users (pk),

  result1   INTEGER,
  result12  INTEGER,
  points    INTEGER,
  total     INTEGER,
  given     TIMESTAMP NOT NULL DEFAULT NOW()
);
