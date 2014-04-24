SELECT
    team1,
    team2,
    result1,
    result2,
    to_char(starts, 'MM-DD'),
    happened,
    closed
    FROM games
    ORDER BY starts
;
