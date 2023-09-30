package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// "fmt"
	"server/internal/models"
	// "golang.org/x/tools/go/analysis/passes/tests"
)

//inner most layer interacting with database
//create interface
type DTBX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

//create struct type of repository
type repository struct {
	db DTBX
}

//create func and inject it with interface
func NewUserRepository(db DTBX) models.Repository {
	// return a pointer to the repository
	return &repository{
		db: db,
	}
}

//function to create user
func (repo *repository) CreateUser(ctx context.Context, hashed_code string, user *models.User) (*models.User, error) {
	var lastinsertid int
	// repo.db.QueryContext(ctx, "UPDATE users set id_number=$2,password=$3,plate_numer=$4,license_number=$5 WHERE phone = $1",user.Phone,user.ID_NO,user.Password,user.PlateNumber,user.LicenseNumber)

	query := "INSERT INTO users (first_name,last_name,email,id_number,phone,password) VALUES ($1,$2,$3,$4,$5,$6) returning id"
	err := repo.db.QueryRowContext(ctx, query, user.FirstName, user.LastName, user.Email, user.ID_NO, user.Phone, user.Password).Scan(&lastinsertid)
	if err != nil {
		return &models.User{}, err
	}
	// create post
	var lastinsertids int
	// for i := 1; i <= 100;i++ {
	test := "INSERT INTO tests (name,user_id) VALUES ($1,$2) returning id"
	errs := repo.db.QueryRowContext(ctx, test, user.FirstName, 3).Scan(&lastinsertids)
	if errs != nil {
		return &models.User{}, err
	}
	validate := "INSERT INTO account_validation (hashed_code,user_id) VALUES ($1,$2) returning id"
	valid := repo.db.QueryRowContext(ctx, validate, hashed_code, lastinsertid).Scan(&lastinsertids)
	if valid != nil {
		return &models.User{}, err
	}
	// }
	//typecast id to interface
	user.ID = int64(lastinsertids)
	var tesxt = models.TestIn{}
	tesxt.ID = int64(lastinsertids)
	tesxt.NAME = user.FirstName
	tesxt.Users = "lastinsertid"
	user.Test = []models.TestIn{tesxt}
	return user, nil
}

//get user by email
func (user *repository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	//empty struct
	u := models.User{}
	//    t:=models.TestIn{}
	query := "SELECT users.id, users.email, users.first_name,users.last_name,users.phone,users.id_number, users.password FROM users join tests on tests.user_id=users.id WHERE users.email = $1"
	err := user.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.Phone, &u.ID_NO, &u.Password)
	if err != nil {
		return &models.User{}, err
	}
	//    get relationship data
	//    queryx := "SELECT * FROM tests WHERE user_id = $1"
	rows, err := user.db.QueryContext(ctx, "SELECT * FROM tests WHERE user_id = $1 limit $2", u.ID, 50)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []models.TestIn
	for rows.Next() {
		var website models.TestIn
		if err := rows.Scan(&website.ID, &website.NAME, &website.Users); err != nil {
			return nil, err
		}
		all = append(all, website)
	}
	//    return all, nil
	u.Test = all

	return &u, nil
}
func (user *repository) Check_exist_email(ctx context.Context, email string) (*models.User, error) {
	//empty struct
	t := models.User{}
	queryx := "SELECT id,email FROM users WHERE email = $1"
	err := user.db.QueryRowContext(ctx, queryx, email).Scan(&t.ID, &t.Email)
	if err != nil {
		return &models.User{}, nil
	}
	return &t, nil
}
func (user *repository) UpdateUser(ctx context.Context, email string, userx *models.User) (*models.User, error) {
	//empty struct
	t := models.User{}
	_, err := user.db.ExecContext(ctx, "UPDATE users set id_number=$2,password=$3,plate_number=$4,license_number=$5,is_verified=$6 WHERE id = $1", email, userx.ID_NO, userx.Password, userx.PlateNumber, userx.LicenseNumber, false)
	if err != nil {
		log.Fatal(err)
		fmt.Println("update", err)
	}
	t.Phone = userx.Phone
	t.LicenseNumber = userx.LicenseNumber
	return userx, nil
}
func (user *repository) GetCode(ctx context.Context, uid int64) (*models.User, error) {
	//empty struct
	u := models.User{}
	//    t:=models.TestIn{}
	query := "SELECT hashed_code,users.is_verified FROM account_validation join users on users.id=account_validation.user_id WHERE account_validation.user_id = $1"
	err := user.db.QueryRowContext(ctx, query, uid).Scan(&u.Password, &u.IsVerified)
	if err != nil {
		return nil, err
	}
	//  //    get relationship data
	//  //    queryx := "SELECT * FROM tests WHERE user_id = $1"
	// 	rows, err := user.db.QueryContext(ctx, "SELECT * FROM tests WHERE user_id = $1 limit $2",u.ID,50)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	defer rows.Close()

	// 	var all []models.TestIn
	// 	for rows.Next() {
	// 		var website models.TestIn
	// 		if err := rows.Scan(&website.ID, &website.NAME, &website.Users); err != nil {
	// 			return nil, err
	// 		}
	// 		all = append(all, website)
	// 	}
	//  //    return all, nil
	// 	u.Test=all

	return &u, nil
}
func (user *repository) Verify(ctx context.Context, uid int64) (*models.CreateUserResponse, error) {
	//empty struct
	//  t := models.User{}
	_, err := user.db.ExecContext(ctx, "UPDATE users set is_verified=$2 WHERE id = $1", uid, true)
	if err != nil {
		log.Fatal(err)
		fmt.Println("update", err)
	}
	userx := &models.CreateUserResponse{
		IsVerified: true,
	}
	return userx, nil
}
