package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	dbName   = "postgres"
	user     = "postgres"
	password = "password"
)

//Get a postgres client
func getPostgresClient() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	cli, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return cli, err
}

//Convert a string into float
func convertToTime(str string) float64 {
	if s, err := strconv.ParseFloat(str, 32); err == nil {
		return s
	}
	return 0
}

//Total time = planning time + execution time
// Sample response of explain analyse for insert command
// EXPLAIN ANALYSE insert into sales.employee (id,name,role,salary) values(6,'name1','role1',1);
// Insert on employee  (cost=0.00..0.01 rows=1 width=72) (actual time=0.158..0.158 rows=0 loops=1)
// ->  Result  (cost=0.00..0.01 rows=1 width=72) (actual time=0.003..0.003 rows=1 loops=1)
// Planning time: 0.072 ms
// Execution time: 0.288 ms

func extractTime(str string) float64 {
	parts := strings.Split(str, "Execution time: ")
	if len(parts) > 1 {
		timeInMs := strings.Split(parts[1], " ms")
		if len(timeInMs) > 0 {
			//		fmt.Println("Execution time is:", timeInMs[0])
			return convertToTime(timeInMs[0])
		}
	}
	parts = strings.Split(str, "Planning time: ")
	if len(parts) > 1 {
		timeInMs := strings.Split(parts[1], " ms")
		if len(timeInMs) > 0 {
			//		fmt.Println("Planning time is:", timeInMs[0])
			return convertToTime(timeInMs[0])
		}
	}
	return 0
}

//Execute query without transaction
func execute(cli *sql.DB, commandList []string) (float64, error) {
	var time float64
	for _, command := range commandList {
		rows, err := cli.Query(command)
		if err != nil {
			fmt.Printf("Error1 is:%v\n", err)
			return time, err
		}

		for rows.Next() {
			var str string
			if err := rows.Scan(&str); err != nil {
				fmt.Printf("Error2 is:%v\n", err)
			}
			time += extractTime(str)
		}
		if err := rows.Err(); err != nil {
			fmt.Printf("Error3 is:%v\n", err)
		}
	}
	//	fmt.Println("Total time taken : ", fmt.Sprintf("%.3f", time))
	return time, nil
}

//Execute query with transaction
func executeWithTransaction(cli *sql.DB, commandList []string) (float64, error) {
	var time float64
	ctx := context.Background()
	tx, err := cli.BeginTx(ctx, nil)
	if err != nil {
		fmt.Printf("Error1 is:%v\n", err)
	}

	for _, command := range commandList {

		rows, err := cli.Query(command)
		if err != nil {
			fmt.Printf("Error1 is:%v\n", err)
			tx.Rollback()
			return time, err
		}

		for rows.Next() {
			var str string
			if err := rows.Scan(&str); err != nil {
				tx.Rollback()
				fmt.Printf("Error2 is:%v\n", err)
			}
			time += extractTime(str)
		}
		if err := rows.Err(); err != nil {
			tx.Rollback()
			fmt.Printf("Error3 is:%v\n", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("Error3 is:%v\n", err)
	}
	//	fmt.Println("Total time taken : ", fmt.Sprintf("%.3f", time))
	return time, nil
}

func main() {
	cli, err := getPostgresClient()
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer cli.Close()

	for c := 0; c < 10; c++ {
		var t float64
		_, err = cli.Query(`TRUNCATE TABLE sales.employee`)
		if err != nil {
			fmt.Println("Error in truncate")
			return
		}
		if c%2 == 0 {
			fmt.Println("Executing with transaction")
		} else {
			fmt.Println("Executing without transaction")
		}
		for i := 0; i < 1; i++ {
			str := `EXPLAIN ANALYSE insert into sales.employee (id,name,role,salary) values( + ` + strconv.Itoa(i) + `,'name1','role1',1);`
			commandList := []string{str}
			var time float64
			if c%2 == 0 {
				time, err = executeWithTransaction(cli, commandList)
			} else {
				time, err = execute(cli, commandList)
			}
			if err == nil {
				t = t + time
			}
		}
		fmt.Println("Total time taken1 : ", fmt.Sprintf("%.3f", t))
	}
}
