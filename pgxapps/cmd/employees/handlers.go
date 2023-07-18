package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// Database instance
var db *sql.DB

// Employee struct
type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

// Employees struct
type Employees struct {
	Employees []Employee `json:"employees"`
}

// Connect function
func Connect() error {
	var err error
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func GetEmployee(c *fiber.Ctx) error {
	// Select all Employee(s) from database
	rows, err := db.Query("SELECT id, name, salary, age FROM employees order by id")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age); err != nil {
			return err // Exit if we get an error
		}

		// Append Employee to Employees
		result.Employees = append(result.Employees, employee)
	}
	// Return Employees in JSON format
	return c.JSON(result)
}

func PostEmployee(c *fiber.Ctx) error {
	// New Employee struct
	u := new(Employee)

	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Insert Employee into database
	res, err := db.Query("INSERT INTO employees (name, salary, age)VALUES ($1, $2, $3)", u.Name, u.Salary, u.Age)
	if err != nil {
		return err
	}

	// Print result
	log.Println(res)

	// Return Employee in JSON format
	return c.JSON(u)
}

func UpdateEmployee(c *fiber.Ctx) error {
	// New Employee struct
	u := new(Employee)

	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Update Employee into database
	res, err := db.Query("UPDATE employees SET name=$1,salary=$2,age=$3 WHERE id=$5", u.Name, u.Salary, u.Age, u.ID)
	if err != nil {
		return err
	}

	// Print result
	log.Println(res)

	// Return Employee in JSON format
	return c.Status(201).JSON(u)
}

func DeleteEmployee(c *fiber.Ctx) error {
	// New Employee struct
	u := new(Employee)

	// Parse body into struct
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Delete Employee from database
	res, err := db.Query("DELETE FROM employees WHERE id = $1", u.ID)
	if err != nil {
		return err
	}

	// Print result
	log.Println(res)

	// Return Employee in JSON format
	return c.JSON("Deleted")
}
