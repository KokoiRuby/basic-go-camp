package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("duplicate email")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.Ctime = now
	user.Utime = now
	err := dao.db.WithContext(ctx).Create(&user).Error
	// type assertion
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		const uniqueIndexErrNo uint16 = 1062
		if mysqlErr.Number == uniqueIndexErrNo {
			// conflict
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	//err := dao.db.WithContext(ctx).Where("email = ?, email").First(&user).Error
	err := dao.db.WithContext(ctx).First(&user, "email = ?", email).Error
	return user, err
}

// User maps to DB entity/model/PO Persistent Object
type User struct {
	// best practice uint64
	Id       int64  `gorm:"primaryKey, autoincrement"`
	Email    string `gorm:"type:varchar(100);unique"`
	Password string `gorm:"type:varchar(100)"`

	// extra

	// utc
	Ctime int64 `gorm:"index"`
	Utime int64 `gorm:"index"`
}

type UserDetail struct {
}
