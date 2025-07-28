package mysql

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func NewMySQLConnection(user, password, host, dbname string, port int) (*gorm.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
        user, password, host, port, dbname,
    )
    
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
