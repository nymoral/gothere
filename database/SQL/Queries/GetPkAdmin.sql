SELECT pk, admin
    FROM users
    WHERE email=$1
;
