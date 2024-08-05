package team

import (
	"gofiber/app/models/model/usersModel"
	"gofiber/app/models/model/walletsModel"
	"gofiber/app/module/context"
	"gofiber/app/module/database"
	"gorm.io/gorm"
	"time"
)

// Details 团队收益
func Details(ctx *context.CustomCtx, params *context.NoRequestBody) error {
	nowTime := time.Now()
	todayTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
	userId := ctx.UserId
	userTeamInfo := &userTeamDetails{}
	database.Db.Model(userTeamInfo).Where("id = ?", userId).
		Preload("Record", func(db *gorm.DB) *gorm.DB {
			return db.Select("wallet_user_bill.user_id", "u.username", "wallet_user_bill.money", "wallet_user_bill.name", "wallet_user_bill.created_at").
				Where("wallet_user_bill.type IN ?", []int{walletsModel.WalletUserBillTypeTeamEarnings, walletsModel.WalletUserBillTypeShareAward}).
				Joins("LEFT JOIN user as u on u.id=wallet_user_bill.source_id")
		}).
		Preload("TotalEarning", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id", "IFNULL(sum(money), 0) as Money").
				Where("type IN ?", []int{walletsModel.WalletUserBillTypeTeamEarnings, walletsModel.WalletUserBillTypeShareAward}).Group("user_id")
		}).
		Preload("TodayEarning", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id", "IFNULL(sum(money), 0) as Money").
				Where("created_at BETWEEN ? AND ?", todayTime.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")).
				Where("type IN ?", []int{walletsModel.WalletUserBillTypeTeamEarnings, walletsModel.WalletUserBillTypeShareAward}).Group("user_id")
		}).
		Preload("TotalPeople", func(db *gorm.DB) *gorm.DB {
			return db.Select("admin_id", "count(*) as Nums").Where("parent_id = ?", userId).Group("admin_id")
		}).
		Preload("TodayPeople", func(db *gorm.DB) *gorm.DB {
			return db.Select("admin_id", "count(*) as Nums").
				Where("parent_id = ?", userId).
				Where("created_at BETWEEN ? AND ?", todayTime.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")).
				Group("admin_id")
		}).
		Preload("TotalDeposit", func(db *gorm.DB) *gorm.DB {
			userIds := make([]uint, 0)
			database.Db.Model(&usersModel.User{}).Select("id").Where("parent_id = ?", userId).Find(&userIds)
			if userIds == nil {
				userIds = append(userIds, 0)
			}

			return db.Select("admin_id", "IFNULL(sum(money), 0) as Money").
				Where("user_id IN ?", userIds).
				Where("type = ?", walletsModel.WalletUserOrderTypeDeposit).
				Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
				Group("admin_id")
		}).
		Preload("TodayDeposit", func(db *gorm.DB) *gorm.DB {
			userIds := make([]uint, 0)
			database.Db.Model(&usersModel.User{}).Select("id").Where("parent_id = ?", userId).Find(&userIds)
			if userIds == nil {
				userIds = append(userIds, 0)
			}

			return db.Select("admin_id", "IFNULL(sum(money), 0) as Money").
				Where("user_id IN ?", userIds).
				Where("created_at BETWEEN ? AND ?", todayTime.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")).
				Where("type = ?", walletsModel.WalletUserOrderTypeDeposit).
				Where("status = ?", walletsModel.WalletUserOrderStatusComplete).
				Group("admin_id")
		}).
		Find(userTeamInfo)
	return ctx.SuccessJson(userTeamInfo)
}

type userTeamDetails struct {
	usersModel.User
	TotalPeople  userTeamPeople  `json:"totalPeople" gorm:"foreignKey:AdminId;references:AdminId"`
	TodayPeople  userTeamPeople  `json:"todayPeople" gorm:"foreignKey:AdminId;references:AdminId"`
	TotalDeposit userWalletOrder `json:"totalDeposit" gorm:"foreignKey:AdminId;references:AdminId"`
	TodayDeposit userWalletOrder `json:"todayDeposit" gorm:"foreignKey:AdminId;references:AdminId"`
	TotalEarning userTeamEarning `json:"totalEarning" gorm:"foreignKey:UserId"`
	TodayEarning userTeamEarning `json:"todayEarning" gorm:"foreignKey:UserId"`
	Record       []userRecord    `json:"record" gorm:"foreignKey:UserId"`
}

type userRecord struct {
	UserId    uint      `json:"userId"`    // 用户ID
	UserName  string    `json:"userName"`  // 用户账户
	Money     float64   `json:"money"`     // 收益金额
	Name      string    `json:"name"`      // 记录标题
	CreatedAt time.Time `json:"createdAt"` // 收益时间
}

func (userRecord) TableName() string {
	return "wallet_user_bill"
}

func (userTeamDetails) TableName() string {
	return "user"
}

type userTeamPeople struct {
	AdminId uint `json:"adminId"` // 上级ID
	Nums    int  `json:"nums"`    // 数量
}

func (userTeamPeople) TableName() string {
	return "user"
}

type userWalletOrder struct {
	AdminId uint    `json:"adminId"` // 用户ID
	Money   float64 `json:"money"`   // 总收益
}

func (userWalletOrder) TableName() string {
	return "wallet_user_order"
}

type userTeamEarning struct {
	UserId uint    `json:"userId"` // 用户ID
	Money  float64 `json:"money"`  // 总收益
}

func (userTeamEarning) TableName() string {
	return "wallet_user_bill"
}
