UPDATE guesses
    SET
       total=$3,
       points=$4
    WHERE game_pk=$1 AND user_pk=$2;
