SELECT
    pk,
    team1,
    team2,
    to_char(starts, 'MM-DD')
    FROM games
    WHERE closed=false
    ORDER BY starts
;
