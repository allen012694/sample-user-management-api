package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/allen012694/usersystem/config"
	"github.com/allen012694/usersystem/utils"
	_ "github.com/go-sql-driver/mysql"
)

const DEFAULT_PWD = "123456"

// POPULATE 10 million users
func main() {
	start := time.Now()

	// open connection
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	hashedPassword := utils.SHA256WithSalt(DEFAULT_PWD, config.SALT)
	statement1Million := fmt.Sprintf(`INSERT INTO user_tab (username, nickname, password_hash)
			WITH cteTen AS (
				SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL
				SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5 UNION ALL
				SELECT 6 UNION ALL SELECT 7 UNION ALL SELECT 8 UNION ALL
				SELECT 9
			)
			SELECT @usrnm := CONCAT('u_', UUID_SHORT()), @usrnm, '%v'
			FROM cteTen AS POW1, cteTen AS POW2, cteTen AS POW3,
					cteTen AS POW4, cteTen AS POW5, cteTen AS POW6`, hashedPassword)

	for i := 0; i < 10; i += 1 {
		_, err = db.Exec(statement1Million)
		if err == nil {
			fmt.Println("1,000,000 records inserted.")
		}
	}
	fmt.Println("DONE after", time.Since(start), "seconds")

}

// QUICKLY POPULATE 1mil records
// INSERT INTO user_tab (username, nickname, password_hash)
// WITH cteTen AS (
//    SELECT 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL
//    SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5 UNION ALL
//    SELECT 6 UNION ALL SELECT 7 UNION ALL SELECT 8 UNION ALL
//    SELECT 9
// )
// SELECT @usrnm := CONCAT('u_', UUID_SHORT()), @usrnm, SHA2(RANDOM_BYTES(10), 256)
// FROM cteTen AS POW1, cteTen AS POW2, cteTen AS POW3,
//      cteTen AS POW4, cteTen AS POW5, cteTen AS POW6;
