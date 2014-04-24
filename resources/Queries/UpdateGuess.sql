UPDATE guesses
    SET result1=$1,
        result2=$2,
        given=now()
    WHERE game_pk=$3 AND
    user_pk=$4
;

