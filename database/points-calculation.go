package database

import (
    "log"
    "database/sql"
    _ "github.com/lib/pq"
    "gothere/models"
)

func abs(i int) (int) {
    if i < 0 {
        return -i
    }
    return i
}

func outcome(t1, t2 int) (int) {
    // Returns an outcome code for a given result:
    // 0 -- if the game was even
    // 1 -- first team won
    // 2 -- second team won
    if t1 == t2 {
        return 0
    }
    if t1 > t2 {
        return 1
    }
    return 2
}

func SinglePoints(u1, u2, t1, t2 int) (int, bool) {
    correct_out := outcome(t1, t2)
    user_out := outcome(u1, u2)

    points := 0

    if correct_out != user_out {
        // Didn't guess the outcome.
        points = 3
    }
    if t1 == u1 && t2 == u2 {
        // Correct guess!
        points = -3
    }
    points += abs(u1 - t1)
    points += abs(u2 - t2)

    return points, user_out == correct_out
}

func fillPoints(points []models.Points, t1, t2 int) {
    max := -24
    correct := 0

    for i, P := range points {
        if P.Given {
            toGive, out := SinglePoints(int(P.Result1.(int64)), int(P.Result2.(int64)), t1, t2)
            if out {
                correct += 1
            }
            if toGive >= max {
                max = toGive
            }
            points[i].Points = toGive
       }
    }
    for i, P := range points {
        if ! P.Given {
            points[i].Points = max
        } else {
            if correct == 1 && P.Points == -3 {
                points[i].Points = -7
            }
        }
    }
}

func CalcPoints(db *sql.DB, pk int, t1, t2 int) {

    rows, err := db.Query(qGetPoints, pk)
    if err != nil {
        log.Fatal(err)
    }

    points := make([]models.Points, 0)
    var P models.Points

    for rows.Next() {
        err := rows.Scan(&P.UserPk, &P.Total, &P.Result1, &P.Result2)
        if err != nil {
            log.Println(err)
        } else {
            P.Given = P.Result1 != nil && P.Result2 != nil
            points = append(points, P)
        }
    }
    fillPoints(points, t1, t2)
    sendPointsToDb(db, pk, points)
}

func sendPointsToDb(db *sql.DB, pk int, points []models.Points) {
    for _, P := range(points) {
        c := 0
        // Correct will be increased by c
        if P.Points < -2 {
            c = 1
        }
        _, err := db.Exec(qUpdatePoints, P.UserPk, P.Points, c)
        // Updating Users table with new points.
        if err != nil {
            log.Println(err)
        }
        if P.Given {
            // Guess already given, safe to update.
            _, err := db.Exec(qUpdateGuessPoints, pk, P.UserPk, P.Total+P.Points, P.Points)
            if err != nil {
                log.Println(err)
            }
        } else {
            // Insert a guess.
            _, err := db.Exec(qInsertGuessPoints, pk, P.UserPk, P.Total+P.Points, P.Points)
            if err != nil {
                log.Println(err)
            }
        }
    }
}
