UPDATE users SET
    points = points + $2,
    correct = correct + $3
WHERE pk=$1;

