INSERT INTO guesses
    (user_pk, game_pk, points, total, result1, result2)
    VALUES ($2, $1, $4, $3, -1, -1);
