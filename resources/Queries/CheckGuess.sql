SELECT pk FROM guesses
    WHERE
        game_pk=$1 AND
        user_pk=$2
;
