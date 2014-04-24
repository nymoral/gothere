SELECT gs.result1,
       gs.result2,
       gs.points,
       gs.total,
       G.happened
    FROM games G
    LEFT JOIN users AS U ON U.admin=false
    LEFT JOIN guesses AS gs ON gs.game_pk=G.pk AND
        gs.user_pk=U.pk AND (U.pk=$1 OR G.closed=true OR G.happened=true)
    ORDER BY
         U.points ASC,
         U.correct,
         U.pk,
         G.starts
;
