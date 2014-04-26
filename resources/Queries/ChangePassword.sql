UPDATE users
    SET password=$2
    WHERE email=$1
;
