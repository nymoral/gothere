SELECT password, admin
    FROM users
    WHERE email=$1
;
