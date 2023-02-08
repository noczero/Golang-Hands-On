package Database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"testing"
	"time"
)

func TestDBConnection(t *testing.T) {
	// make db connection
	dbConnection := GetConnection()
	defer dbConnection.Close() // close connection

	fmt.Println(dbConnection)
}

func TestExecSql(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background() // use context

	// define query
	query := "INSERT INTO customer(id, name) VALUES ('satrya', 'Satrya')"

	// use exec context, to execute query without returning value like INSERT, DROP, UPDATE.
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestQueryContextSql(t *testing.T) {
	db := GetConnection()
	defer db.Close() // close db

	ctx := context.Background()

	// query
	query := "SELECT id, name FROM customer"

	// use query context to execute SELECT
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close() // close rows

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name) // convert rows value to true value

		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
	}
}

func TestQuerySqlComplexSelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string        // varchar not null
		var email sql.NullString   // if db column is null then use it, its a struct
		var balance float64        // double type not null
		var rating int32           // int type not null
		var birthDate sql.NullTime // TIMESTAMP null
		var createdAt time.Time    // TIMESTAMP or DATE NOT NULL
		var married bool

		// convert data sequence with query column
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("====================")
		fmt.Println("id : ", id)
		fmt.Println("name : ", name)

		if email.Valid {
			fmt.Println("email : ", email.String)
		}

		fmt.Println("balance : ", balance)
		fmt.Println("rating : ", rating)

		if birthDate.Valid {
			fmt.Println("birthDate : ", birthDate.Time)
		}

		fmt.Println("married : ", married)
		fmt.Println("createdAt : ", createdAt)

	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #" // query injection '; #
	password := "123"

	// passing parameter value in query is vulnerable from query injection
	query := "SELECT username FROM user WHERE username='" + username + "' AND password='" + password + "' LIMIT 1"
	fmt.Println(query) // SELECT username FROM user WHERE username='admin'; #' AND password='123' LIMIT 1

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// if just one data use if
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Login Success!", username) // will success
	} else {
		fmt.Println("Login Failed!")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "1234"

	//username := "admin"
	//password := "admin"

	query := "SELECT username FROM user WHERE username=? AND password=? LIMIT 1" // use ? to pass parameter in query context
	rows, err := db.QueryContext(ctx, query, username, password)                 // pass paarameter as veredict params

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("Login success!", username)
	} else {
		fmt.Println("Login failed!") // it will be login failed
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; DROP TABLE user; #" // this query will drop user table if sql injection happened
	password := "admin"

	query := "INSERT INTO user(username, password) VALUES (?,?)"
	_, err := db.ExecContext(ctx, query, username, password) // use exec context with passing parameter

	if err != nil {
		panic(err)
	}
	fmt.Println("Insert to user success")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "satrya@test.com"
	comment := "test comment"

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId() // get last insert id
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id", insertId)
}

func TestName(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments(email, comment) VALUES(?,?)" // this query will be set as prepare statement
	statement, err := db.PrepareContext(ctx, query)             // set query as prepare context to insert many values efficienly

	if err != nil {
		panic(err)
	}
	defer statement.Close()

	// iterate simulation for bulking insert
	for i := 0; i < 10; i++ {
		email := "dummy" + strconv.Itoa(i) + "@test.com"
		comment := "Lorem ipsum " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully insert comment with id:", id)
	}
}
