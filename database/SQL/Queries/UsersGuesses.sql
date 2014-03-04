SELECT
    games.team1,
    games.team2,
    guesses.result1,
    guesses.result2
    FROM games
    LEFT JOIN (SELECT
        game_pk,
        result1,
        result2
        FROM guesses
        WHERE user_pk=$1)
    AS guesses
    ON games.pk=guesses.game_pk
    ORDER BY games.starts
;
