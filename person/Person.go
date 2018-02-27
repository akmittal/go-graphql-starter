package person

import (
	"github.com/jmoiron/sqlx"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	// Postgres driver
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// Person data type
type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	EmailID   string `db:"email"`
	Username  string `db:"username"`
	Password  string `db:"password"`
}

func (p Person) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.FirstName, validation.Required, validation.Length(2, 50)),
		validation.Field(&p.LastName, validation.Required, validation.Length(2, 50)),
		validation.Field(&p.EmailID, validation.Required, is.Email),
		validation.Field(&p.Username, validation.Required, validation.Length(5, 50), is.Alphanumeric))
}

// Save the person object to DB
func (p Person) Save(db *sqlx.DB) (Person, error) {
	tx := db.MustBegin()
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), 10)
	if err != nil {
		return Person{}, err
	}
	p.Password = string(hash)
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	_, err = tx.NamedExec("INSERT INTO users (first_name, last_name, email, username, password) VALUES (:first_name, :last_name, :email, :username, :password)", p)
	if err != nil {
		return Person{}, err
	}
	err = tx.Commit()
	return p, err
}

// ComparePassword :Compares person password
func (p *Person) ComparePassword(db *sqlx.DB) (Person, error) {

	rows, err := db.Queryx("SELECT * FROM users WHERE USERNAME=$1", p.Username)
	if err != nil {
		return Person{}, err
	}
	var person Person
	for rows.Next() {
		rows.StructScan(&person)
	}

	err = bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(p.Password))
	if err != nil {
		return Person{}, err
	}
	return person, err
}

func GetByUsername(username string, db *sqlx.DB) (Person, error) {
	var user Person

	row := db.QueryRowx("SELECT * FROM users WHERE USERNAME=$1", username)

	err := row.StructScan(&user)
	if err != nil {
		return Person{}, err
	}
	return user, nil

}
