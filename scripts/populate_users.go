package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"shopee.com/entry/config"
)

const BATCH_SIZE = 10000

// Script to populate users into test DB
func main() {
	start := time.Now()
	inputAmount := os.Args[1]
	if len(inputAmount) > 0 {
		populateAmount, err := strconv.ParseInt(inputAmount, 10, 64)
		if err != nil {
			panic(err.Error())
		}

		// open connection
		db, err := sql.Open("mysql", config.DSN)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		// count current data size
		res, err := db.Query("SELECT COUNT(id) FROM user_tab")
		if err != nil {
			panic(err.Error())
		}
		defer res.Close()
		var totalExistedAmount int64
		res.Next()
		res.Scan(&totalExistedAmount)

		for i := 0; i < int(populateAmount); i += 1 {
			// next password
			data := make([]byte, 10)
			rand.Read(data)
			randomPwd := fmt.Sprintf("%x", sha256.Sum256(data))

			// next username
			username := fmt.Sprintf("user_%v", totalExistedAmount+int64(i))

			statement := fmt.Sprintf("INSERT INTO user_tab (username, nickname, password_hash) VALUES ('%v', '%v', '%v')", username, username, randomPwd)
			_, err = db.Exec(statement)
			if err != nil {
				fmt.Println("Execute", statement, "with error:", err.Error())
			}

			if i%BATCH_SIZE == 0 {
				fmt.Println("10000 records inserted...")
			}
		}
		fmt.Println(populateAmount%BATCH_SIZE, "records inserted...")

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
// SELECT CONCAT('u_', UUID_SHORT()), 'just a name', SHA2(RANDOM_BYTES(10), 256)
// FROM cteTen AS POW1, cteTen AS POW2, cteTen AS POW3,
//      cteTen AS POW4, cteTen AS POW5, cteTen AS POW6, (SELECT @row:=0) t;
