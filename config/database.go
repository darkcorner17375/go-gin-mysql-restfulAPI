package config

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type configDatabase struct {
	Host                      string
	Port                      string
	DbName                    string
	UserName                  string
	Password                  string
	Charset                   string        //要支持完整的UTF-8編碼,需設置成: utf8mb4
	AutoMigrate               bool          // 初始化時調用數據遷移
	ParseTime                 bool          //解析time.Time類型
	TimeZone                  string        // 時區,若設置 Asia/Shanghai,需寫成: Asia%2fShanghai
	DefaultStringSize         uint          // string 類型字段的默認長度
	DisableDatetimePrecision  bool          // 禁用 datetime 精度
	SkipInitializeWithVersion bool          // 根據當前 MySQL 版本自動配置
	SlowSql                   time.Duration //慢SQL
	LogLevel                  string        // 日誌記錄級別
	IgnoreRecordNotFoundError bool          // 是否忽略ErrRecordNotFound(未查到記錄錯誤)
	// Gorm                      gorm
}

// // gorm 配置信息
// type gorm struct {
// 	SkipDefaultTx   bool   //是否跳過默認事務
// 	CoverLogger     bool   //是否覆蓋默認logger
// 	PreparedStmt    bool   // 設置SQL緩存
// 	CloseForeignKey bool   // 禁用外鍵約束
// 	TablePrefix     string // 表前綴
// 	SingularTable   bool   //是否使用單數表名(默認複數)，啓用後，User結構體表將是user
// }

var Database configDatabase
var DB *gorm.DB

func initDatabaseConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		Database.Host = ""
		Database.Port = ""
		Database.DbName = ""
		Database.UserName = ""
		Database.Password = ""
	case gin.DebugMode:
		Database.Host = os.Getenv("DB_HOST")
		Database.Port = os.Getenv("DB_PORT")
		Database.DbName = os.Getenv("DB_NAME")
		Database.UserName = os.Getenv("DB_User")
		Database.Password = os.Getenv("DB_PASSWORD")
	}
}

// type User struct {
// 	ID   int
// 	Name string
// 	Age  int
// }

func InitGorm() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		Database.UserName, Database.Password, Database.Host, Database.Port, Database.DbName, "utf8mb4",
		true, "Local")

	// 設置gorm配置
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: false, //是否跳過默認事務
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "app_",
			SingularTable: true,
		},
		// 執行任何SQL時都會創建一個prepared statement並將其緩存，以提高後續的效率
		PrepareStmt: false,
		//在AutoMigrate 或 CreateTable 時，GORM 會自動創建外鍵約束，若要禁用該特性，可將其設置爲 true
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	// // 是否覆蓋默認sql配置
	// setNewLogger(gormConfig)
	// if mysqlConfig.Gorm.CoverLogger {
	// 	setNewLogger(gormConfig)
	// }

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		SkipInitializeWithVersion: false,
	}), gormConfig)
	if err != nil {
		panic(fmt.Sprintf("Gorm & Mysql資料庫連接失敗: %s", err))
	}
	// 賦值給全局變量
	DB = db

	// // 建立表格
	// db.AutoMigrate(&User{})

	// // 新增資料
	// user := User{Name: "John", Age: 18}
	// db.Create(&user)

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
