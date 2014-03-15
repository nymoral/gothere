SELECT
    users.pk,
    users.points,
    guesses.result1,
    guesses.result2
    FROM (SELECT points, pk FROM users WHERE admin=false) as users
    LEFT JOIN (SELECT user_pk, result1, result2
                FROM guesses WHERE game_pk=$1)
    AS guesses
    ON guesses.user_pk=users.pk
;

