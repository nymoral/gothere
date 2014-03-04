SELECT
    pk,
    team1,
    team2,
    to_char(starts, 'MM-DD')
    FROM games
    WHERE happened=false AND
        closed=true
    ORDER BY starts
;
