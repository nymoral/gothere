SELECT firstname, substr(lastname, 1, 1), pk
    FROM users
        WHERE admin=false
        ORDER BY points ASC,
            correct, pk
;
