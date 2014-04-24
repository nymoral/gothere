UPDATE games
    SET closed = TRUE,
        happened = TRUE,
        result1=$1,
        result2=$2
    WHERE pk=$3
;
