package d4_database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitializeSqlx(dsn string) {
	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatal("connect to db failed", err)
		return
	}
	defer db.Close()

	SearchRows(db)

	HighestSalary(db)

	FindBooks(db)
}

type Employee struct {
	Id         uint   `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     uint   `db:"salary"`
}

func CreateTable(db *sqlx.DB) {
	schema := `CREATE TABLE employees (
				id SERIAL PRIMARY KEY,
				name VARCHAR(100) NOT NULL,
				department VARCHAR(100) NOT NULL,
				salary INT NOT NULL
				);`
	db.MustExec(schema)
}

func InsertRow(db *sqlx.DB) {
	// NamedQuery
	employee := Employee{Name: "Jane", Department: "技术部", Salary: 4000}
	query := `INSERT INTO employees (name, department, salary) VALUES(:name, :department, :salary) RETURNING id;`
	rows, _ := db.NamedQuery(query, &employee)

	defer rows.Close()

	var id int
	if rows.Next() {
		rows.Scan(&id)
		fmt.Println(id)
	}
}

func InsertRows(db *sqlx.DB) {
	// NamedExec

	employees := []Employee{
		{Name: "王耀武", Department: "技术部", Salary: 7000},
		{Name: "薛岳", Department: "战术部", Salary: 8000},
		{Name: "麦克阿瑟", Department: "美国部", Salary: 12000},
		{Name: "杜聿明", Department: "技术部", Salary: 7000},
		{Name: "孙立人", Department: "技术部", Salary: 9000},
	}

	tx := db.MustBegin()
	for _, v := range employees {
		query := `INSERT INTO employees (name, department, salary) VALUES(:name, :department, :salary)`
		tx.NamedExec(query, &v)
	}
	tx.Commit()
}

func SearchOne(db *sqlx.DB) {

	var employee Employee
	err := db.Get(&employee, `SELECT * FROM employees WHERE name = $1`, "王耀武")
	if err != nil {
		fmt.Println("No row found. ", err)
		return
	}
	fmt.Println(employee)
}

func SearchRows(db *sqlx.DB) {
	var employees []Employee
	//Select返回切片
	err := db.Select(&employees, `SELECT * FROM employees WHERE department = $1`, "技术部")
	if err != nil {
		fmt.Println("Do not fount any employee matched. ", err)
		return
	}
	for _, v := range employees {
		fmt.Println(v)
	}
}

func HighestSalary(db *sqlx.DB) {
	var employee Employee
	err := db.Get(&employee, `SELECT * FROM employees ORDER BY salary DESC LIMIT 1`)
	if err != nil {
		fmt.Println("Cannot find any")
		return
	}
	fmt.Println(employee)
}

//2

type Book struct {
	Id     uint   `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  uint   `db:"price"`
}

func CreateBook(db *sqlx.DB) {
	schema := `CREATE TABLE books (
				id SERIAL PRIMARY KEY,
				title VARCHAR(50) NOT NULL,
				author VARCHAR(20) NOT NULL,
				price INT NOT NULL
				);`
	db.MustExec(schema)
}

func InsertBook(db *sqlx.DB) {
	book := Book{Title: "Iron Man II", Author: "Gray Valderin", Price: 88}
	query := "INSERT INTO books (title, author, price) VALUES(:title, :author, :price) RETURNING id"
	rows, err := db.NamedQuery(query, &book)
	if err != nil {
		return
	}
	defer rows.Close()
	fmt.Println(rows)

	var id int
	if rows.Next() {
		rows.Scan(&id)
		fmt.Println(id)
	}
}

func InsertBooks(db *sqlx.DB) {
	books := []Book{
		{Title: "The Great GatesBy", Author: "Gates", Price: 188},
		{Title: "Finance", Author: "Alanxdra Sandrio", Price: 59},
		{Title: "Mathematics", Author: "Hua", Price: 48},
		{Title: "Human", Author: "Vacyli", Price: 288},
	}
	tx := db.MustBegin()
	var query string
	for _, v := range books {
		query = "INSERT INTO books(title, author, price) VALUES(:title, :author, :price) RETURNING id"
		tx.NamedExec(query, &v)
	}
	tx.Commit()
}

func FindBook(db *sqlx.DB) {
	var book Book
	err := db.Get(&book, `SELECT * FROM books WHERE title = $1`, "Iron Man II")
	if err != nil {
		fmt.Println("not found")
		return
	}
	fmt.Println(book)
}

func FindBooks(db *sqlx.DB) {
	var books []Book
	err := db.Select(&books, `SELECT * FROM books WHERE price > $1 ORDER BY price DESC`, 50)
	if err != nil {
		fmt.Println("not found any")
		return
	}
	for _, v := range books {
		fmt.Println(v)
	}
}
