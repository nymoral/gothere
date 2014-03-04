UPDATE games
    SET closed = TRUE
    WHERE pk=$1
;
