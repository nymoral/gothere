CREATE TABLE guesses
(
  pk        BIGSERIAL PRIMARY KEY,

  game_pk   BIGSERIAL REFERENCES games (pk),
  user_pk   BIGSERIAL REFERENCES users (pk),

  result1   INTEGER,
  result2   INTEGER,
  points    INTEGER,
  total     INTEGER,
  given     TIMESTAMP NOT NULL DEFAULT NOW()
);
