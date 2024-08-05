package datas

import (
	"gofiber/app/models/model/adminsModel"
	"gofiber/app/models/model/productsModel"
	"gofiber/app/models/model/types"
)

// InitProductCategory 初始化产品分类
func InitProductCategory() []*productsModel.Category {
	return []*productsModel.Category{
		{GormModel: types.GormModel{ID: 1}, Icon: "/category/womenFashion.png", AdminId: adminsModel.SuperAdminId, Name: "womenWearAndFashion"}, // 女装与时尚
		{GormModel: types.GormModel{ID: 2}, Icon: "", ParentId: 1, AdminId: adminsModel.SuperAdminId, Name: "womenWear"},
		{GormModel: types.GormModel{ID: 3}, Icon: "/category/dress.jpg", ParentId: 2, AdminId: adminsModel.SuperAdminId, Name: "dress"},
		{GormModel: types.GormModel{ID: 4}, Icon: "/category/womenTShirt.jpg", ParentId: 2, AdminId: adminsModel.SuperAdminId, Name: "womenTShirt"},
		{GormModel: types.GormModel{ID: 5}, Icon: "/category/womenShirt.jpg", ParentId: 2, AdminId: adminsModel.SuperAdminId, Name: "womenShirt"},
		{GormModel: types.GormModel{ID: 6}, Icon: "/category/coat.jpg", ParentId: 2, AdminId: adminsModel.SuperAdminId, Name: "coat"},
		{GormModel: types.GormModel{ID: 7}, Icon: "/category/womenJeans.jpg", ParentId: 2, AdminId: adminsModel.SuperAdminId, Name: "womenJeans"},
		{GormModel: types.GormModel{ID: 8}, Icon: "", ParentId: 1, AdminId: adminsModel.SuperAdminId, Name: "womenAccessories"},
		{GormModel: types.GormModel{ID: 9}, Icon: "/category/womenHat.jpg", ParentId: 8, AdminId: adminsModel.SuperAdminId, Name: "womenHat"},
		{GormModel: types.GormModel{ID: 10}, Icon: "/category/womenScarf.jpg", ParentId: 8, AdminId: adminsModel.SuperAdminId, Name: "womenScarf"},
		{GormModel: types.GormModel{ID: 11}, Icon: "/category/womenSocks.jpg", ParentId: 8, AdminId: adminsModel.SuperAdminId, Name: "womenSocks"},

		{GormModel: types.GormModel{ID: 20}, Icon: "/category/menFashion.png", AdminId: adminsModel.SuperAdminId, Name: "menWearAndFashion"}, // 男装与时尚
		{GormModel: types.GormModel{ID: 21}, Icon: "", ParentId: 20, AdminId: adminsModel.SuperAdminId, Name: "menWear"},
		{GormModel: types.GormModel{ID: 22}, Icon: "/category/menTShirt.jpg", ParentId: 21, AdminId: adminsModel.SuperAdminId, Name: "menTShirt"},
		{GormModel: types.GormModel{ID: 23}, Icon: "/category/jacket.jpg", ParentId: 21, AdminId: adminsModel.SuperAdminId, Name: "jacket"},
		{GormModel: types.GormModel{ID: 24}, Icon: "/category/menJeans.jpg", ParentId: 21, AdminId: adminsModel.SuperAdminId, Name: "menJeans"},
		{GormModel: types.GormModel{ID: 25}, Icon: "/category/menShirt.jpg", ParentId: 21, AdminId: adminsModel.SuperAdminId, Name: "menShirt"},
		{GormModel: types.GormModel{ID: 26}, Icon: "/category/suit.jpg", ParentId: 21, AdminId: adminsModel.SuperAdminId, Name: "suit"},
		{GormModel: types.GormModel{ID: 27}, Icon: "", ParentId: 20, AdminId: adminsModel.SuperAdminId, Name: "menAccessories"},
		{GormModel: types.GormModel{ID: 28}, Icon: "/category/menHat.jpg", ParentId: 27, AdminId: adminsModel.SuperAdminId, Name: "menHat"},
		{GormModel: types.GormModel{ID: 29}, Icon: "/category/menScarf.jpg", ParentId: 27, AdminId: adminsModel.SuperAdminId, Name: "menScarf"},
		{GormModel: types.GormModel{ID: 30}, Icon: "/category/menSocks.jpg", ParentId: 27, AdminId: adminsModel.SuperAdminId, Name: "menSocks"},

		{GormModel: types.GormModel{ID: 40}, Icon: "/category/computersAccessories.png", AdminId: adminsModel.SuperAdminId, Name: "computersAndAccessories"}, // 电脑及配件
		{GormModel: types.GormModel{ID: 41}, Icon: "", ParentId: 40, AdminId: adminsModel.SuperAdminId, Name: "desktopAndAllInOne"},
		{GormModel: types.GormModel{ID: 42}, Icon: "/category/allInOne.jpg", ParentId: 41, AdminId: adminsModel.SuperAdminId, Name: "allInOne"},
		{GormModel: types.GormModel{ID: 43}, Icon: "/category/host.jpg", ParentId: 41, AdminId: adminsModel.SuperAdminId, Name: "host"},
		{GormModel: types.GormModel{ID: 44}, Icon: "/category/desktopMachine.jpg", ParentId: 41, AdminId: adminsModel.SuperAdminId, Name: "desktopMachine"},
		{GormModel: types.GormModel{ID: 45}, Icon: "/category/notebook.jpg", ParentId: 41, AdminId: adminsModel.SuperAdminId, Name: "notebook"},
		{GormModel: types.GormModel{ID: 46}, Icon: "/category/mouse.jpg", ParentId: 41, AdminId: adminsModel.SuperAdminId, Name: "mouse"},
		{GormModel: types.GormModel{ID: 47}, Icon: "/category/keyboard.jpg", ParentId: 41, AdminId: adminsModel.SuperAdminId, Name: "keyboard"},
		{GormModel: types.GormModel{ID: 48}, Icon: "", ParentId: 40, AdminId: adminsModel.SuperAdminId, Name: "networkEquipment"},
		{GormModel: types.GormModel{ID: 49}, Icon: "/category/router.jpg", ParentId: 48, AdminId: adminsModel.SuperAdminId, Name: "router"},
		{GormModel: types.GormModel{ID: 50}, Icon: "/category/printServer.jpg", ParentId: 48, AdminId: adminsModel.SuperAdminId, Name: "printServer"},
		{GormModel: types.GormModel{ID: 51}, Icon: "/category/concentrator.jpg", ParentId: 48, AdminId: adminsModel.SuperAdminId, Name: "concentrator"},

		{GormModel: types.GormModel{ID: 60}, Icon: "/category/foodAndBeverage.png", AdminId: adminsModel.SuperAdminId, Name: "foodAndBeverage"}, // 食品和饮料
		{GormModel: types.GormModel{ID: 61}, Icon: "", ParentId: 60, AdminId: adminsModel.SuperAdminId, Name: "food"},
		{GormModel: types.GormModel{ID: 62}, Icon: "/category/candies.jpg", ParentId: 61, AdminId: adminsModel.SuperAdminId, Name: "candies"},
		{GormModel: types.GormModel{ID: 63}, Icon: "/category/jellyPudding.jpg", ParentId: 61, AdminId: adminsModel.SuperAdminId, Name: "jellyPudding"},
		{GormModel: types.GormModel{ID: 64}, Icon: "/category/roastedNuts.jpg", ParentId: 61, AdminId: adminsModel.SuperAdminId, Name: "roastedNuts"},
		{GormModel: types.GormModel{ID: 65}, Icon: "/category/meatSnack.jpg", ParentId: 61, AdminId: adminsModel.SuperAdminId, Name: "meatSnack"},
		{GormModel: types.GormModel{ID: 66}, Icon: "/category/puffedFood.jpg", ParentId: 61, AdminId: adminsModel.SuperAdminId, Name: "puffedFood"},
		{GormModel: types.GormModel{ID: 67}, Icon: "/category/driedFruitPreserves.jpg", ParentId: 61, AdminId: adminsModel.SuperAdminId, Name: "driedFruitPreserves"},
		{GormModel: types.GormModel{ID: 68}, Icon: "", ParentId: 60, AdminId: adminsModel.SuperAdminId, Name: "wineAndTea"},
		{GormModel: types.GormModel{ID: 69}, Icon: "/category/nectar.jpg", ParentId: 68, AdminId: adminsModel.SuperAdminId, Name: "nectar"},
		{GormModel: types.GormModel{ID: 70}, Icon: "/category/wine.jpg", ParentId: 68, AdminId: adminsModel.SuperAdminId, Name: "wine"},
		{GormModel: types.GormModel{ID: 71}, Icon: "/category/cola.jpg", ParentId: 68, AdminId: adminsModel.SuperAdminId, Name: "cola"},

		{GormModel: types.GormModel{ID: 80}, Icon: "/category/babyToy.png", AdminId: adminsModel.SuperAdminId, Name: "babyToy"}, // 母婴玩具
		{GormModel: types.GormModel{ID: 81}, Icon: "", ParentId: 80, AdminId: adminsModel.SuperAdminId, Name: "toddlersToy"},
		{GormModel: types.GormModel{ID: 82}, Icon: "/category/playMat.jpg", ParentId: 81, AdminId: adminsModel.SuperAdminId, Name: "playMat"},
		{GormModel: types.GormModel{ID: 83}, Icon: "/category/kiddieRide.jpg", ParentId: 81, AdminId: adminsModel.SuperAdminId, Name: "kiddieRide"},
		{GormModel: types.GormModel{ID: 84}, Icon: "/category/bathToy.jpg", ParentId: 81, AdminId: adminsModel.SuperAdminId, Name: "bathToy"},
		{GormModel: types.GormModel{ID: 85}, Icon: "/category/percussionToy.jpg", ParentId: 81, AdminId: adminsModel.SuperAdminId, Name: "percussionToy"},
		{GormModel: types.GormModel{ID: 86}, Icon: "/category/rockingHorse.jpg", ParentId: 81, AdminId: adminsModel.SuperAdminId, Name: "rockingHorse"},
		{GormModel: types.GormModel{ID: 87}, Icon: "/category/playTent.jpg", ParentId: 81, AdminId: adminsModel.SuperAdminId, Name: "playTent"},
		{GormModel: types.GormModel{ID: 88}, Icon: "", ParentId: 80, AdminId: adminsModel.SuperAdminId, Name: "plushFabric"},
		{GormModel: types.GormModel{ID: 89}, Icon: "/category/plushPuppet.jpg", ParentId: 88, AdminId: adminsModel.SuperAdminId, Name: "plushPuppet"},
		{GormModel: types.GormModel{ID: 90}, Icon: "/category/plushStorage.jpg", ParentId: 88, AdminId: adminsModel.SuperAdminId, Name: "plushStorage"},
		{GormModel: types.GormModel{ID: 91}, Icon: "/category/plushBackpack.jpg", ParentId: 88, AdminId: adminsModel.SuperAdminId, Name: "plushBackpack"},

		{GormModel: types.GormModel{ID: 100}, Icon: "/category/sportsAndOutdoor.png", AdminId: adminsModel.SuperAdminId, Name: "sportsAndOutdoor"}, // 运动及户外
		{GormModel: types.GormModel{ID: 101}, Icon: "", ParentId: 100, AdminId: adminsModel.SuperAdminId, Name: "sportsOutdoors"},
		{GormModel: types.GormModel{ID: 102}, Icon: "/category/surfingWetsuit.jpg", ParentId: 101, AdminId: adminsModel.SuperAdminId, Name: "surfingWetsuit"},
		{GormModel: types.GormModel{ID: 103}, Icon: "/category/basketballClothes.jpg", ParentId: 101, AdminId: adminsModel.SuperAdminId, Name: "basketballClothes"},
		{GormModel: types.GormModel{ID: 104}, Icon: "/category/yogaClothes.jpg", ParentId: 101, AdminId: adminsModel.SuperAdminId, Name: "yogaClothes"},
		{GormModel: types.GormModel{ID: 105}, Icon: "/category/climbingBoot.jpg", ParentId: 101, AdminId: adminsModel.SuperAdminId, Name: "climbingBoot"},
		{GormModel: types.GormModel{ID: 106}, Icon: "/category/alpenstock.jpg", ParentId: 101, AdminId: adminsModel.SuperAdminId, Name: "alpenstock"},
		{GormModel: types.GormModel{ID: 107}, Icon: "", ParentId: 100, AdminId: adminsModel.SuperAdminId, Name: "outdoorSportsBag"},
		{GormModel: types.GormModel{ID: 108}, Icon: "/category/campingBag.jpg", ParentId: 107, AdminId: adminsModel.SuperAdminId, Name: "campingBag"},
		{GormModel: types.GormModel{ID: 109}, Icon: "/category/gymBag.jpg", ParentId: 107, AdminId: adminsModel.SuperAdminId, Name: "gymBag"},
		{GormModel: types.GormModel{ID: 110}, Icon: "/category/equipmentPack.jpg", ParentId: 107, AdminId: adminsModel.SuperAdminId, Name: "equipmentPack"},
		{GormModel: types.GormModel{ID: 112}, Icon: "/category/protectiveClothing.jpg", ParentId: 107, AdminId: adminsModel.SuperAdminId, Name: "protectiveClothing"},
		{GormModel: types.GormModel{ID: 113}, Icon: "/category/electronicEquipment.jpg", ParentId: 107, AdminId: adminsModel.SuperAdminId, Name: "electronicEquipment"},

		{GormModel: types.GormModel{ID: 120}, Icon: "/category/homeLife.png", AdminId: adminsModel.SuperAdminId, Name: "homeLife"}, // 家居生活
		{GormModel: types.GormModel{ID: 121}, Icon: "", ParentId: 120, AdminId: adminsModel.SuperAdminId, Name: "beddings"},
		{GormModel: types.GormModel{ID: 122}, Icon: "/category/pillowCore.jpg", ParentId: 121, AdminId: adminsModel.SuperAdminId, Name: "pillowCore"},
		{GormModel: types.GormModel{ID: 123}, Icon: "/category/quilt.jpg", ParentId: 121, AdminId: adminsModel.SuperAdminId, Name: "quilt"},
		{GormModel: types.GormModel{ID: 124}, Icon: "/category/mattress.jpg", ParentId: 121, AdminId: adminsModel.SuperAdminId, Name: "mattress"},
		{GormModel: types.GormModel{ID: 125}, Icon: "/category/fourPieceSet.jpg", ParentId: 121, AdminId: adminsModel.SuperAdminId, Name: "fourPieceSet"},
		{GormModel: types.GormModel{ID: 126}, Icon: "", ParentId: 120, AdminId: adminsModel.SuperAdminId, Name: "articlesOfDailyUse"},
		{GormModel: types.GormModel{ID: 127}, Icon: "/category/towelBath.jpg", ParentId: 126, AdminId: adminsModel.SuperAdminId, Name: "towelBath"},
		{GormModel: types.GormModel{ID: 128}, Icon: "/category/slipper.jpg", ParentId: 126, AdminId: adminsModel.SuperAdminId, Name: "slipper"},
		{GormModel: types.GormModel{ID: 129}, Icon: "/category/toiletBrush.jpg", ParentId: 126, AdminId: adminsModel.SuperAdminId, Name: "toiletBrush"},

		{GormModel: types.GormModel{ID: 130}, Icon: "/category/jeweleryAndWatches.png", AdminId: adminsModel.SuperAdminId, Name: "jeweleryAndWatches"}, // 珠宝及手表
		{GormModel: types.GormModel{ID: 131}, Icon: "", ParentId: 130, AdminId: adminsModel.SuperAdminId, Name: "womenJeweleryAccessories"},
		{GormModel: types.GormModel{ID: 132}, Icon: "/category/ring.jpg", ParentId: 131, AdminId: adminsModel.SuperAdminId, Name: "ring"},
		{GormModel: types.GormModel{ID: 133}, Icon: "/category/bracelet.jpg", ParentId: 131, AdminId: adminsModel.SuperAdminId, Name: "bracelet"},
		{GormModel: types.GormModel{ID: 134}, Icon: "/category/necklace.jpg", ParentId: 131, AdminId: adminsModel.SuperAdminId, Name: "necklace"},
		{GormModel: types.GormModel{ID: 135}, Icon: "/category/hairAccessory.jpg", ParentId: 131, AdminId: adminsModel.SuperAdminId, Name: "hairAccessory"},
		{GormModel: types.GormModel{ID: 136}, Icon: "/category/pearlAccessories.jpg", ParentId: 131, AdminId: adminsModel.SuperAdminId, Name: "pearlAccessories"},
		{GormModel: types.GormModel{ID: 137}, Icon: "/category/earring.jpg", ParentId: 131, AdminId: adminsModel.SuperAdminId, Name: "earring"},
		{GormModel: types.GormModel{ID: 138}, Icon: "", ParentId: 130, AdminId: adminsModel.SuperAdminId, Name: "jewelryAccessory"},
		{GormModel: types.GormModel{ID: 139}, Icon: "/category/jewelCase.jpg", ParentId: 138, AdminId: adminsModel.SuperAdminId, Name: "jewelCase"},
		{GormModel: types.GormModel{ID: 140}, Icon: "/category/jewelryCleaning.jpg", ParentId: 138, AdminId: adminsModel.SuperAdminId, Name: "jewelryCleaning"},
		{GormModel: types.GormModel{ID: 141}, Icon: "/category/looseDiamond.jpg", ParentId: 138, AdminId: adminsModel.SuperAdminId, Name: "looseDiamond"},

		{GormModel: types.GormModel{ID: 150}, Icon: "/category/mobileDigital.png", AdminId: adminsModel.SuperAdminId, Name: "mobilePhoneNumber"}, // 手机数码
		{GormModel: types.GormModel{ID: 151}, Icon: "", ParentId: 150, AdminId: adminsModel.SuperAdminId, Name: "mobilePhoneCommunication"},
		{GormModel: types.GormModel{ID: 152}, Icon: "/category/cellPhone.jpg", ParentId: 151, AdminId: adminsModel.SuperAdminId, Name: "cellPhone"},
		{GormModel: types.GormModel{ID: 153}, Icon: "/category/simToolkit.jpg", ParentId: 151, AdminId: adminsModel.SuperAdminId, Name: "simToolkit"},
		{GormModel: types.GormModel{ID: 154}, Icon: "", ParentId: 150, AdminId: adminsModel.SuperAdminId, Name: "mobilePhoneSpareParts"},
		{GormModel: types.GormModel{ID: 155}, Icon: "/category/earphone.jpg", ParentId: 154, AdminId: adminsModel.SuperAdminId, Name: "earphone"},
		{GormModel: types.GormModel{ID: 156}, Icon: "/category/phoneScreenFilm.jpg", ParentId: 154, AdminId: adminsModel.SuperAdminId, Name: "phoneScreenFilm"},
		{GormModel: types.GormModel{ID: 157}, Icon: "/category/selfieStick.jpg", ParentId: 154, AdminId: adminsModel.SuperAdminId, Name: "selfieStick"},
		{GormModel: types.GormModel{ID: 158}, Icon: "/category/protectiveContainment.jpg", ParentId: 154, AdminId: adminsModel.SuperAdminId, Name: "protectiveContainment"},
		{GormModel: types.GormModel{ID: 159}, Icon: "/category/mobilePhoneLens.jpg", ParentId: 154, AdminId: adminsModel.SuperAdminId, Name: "mobilePhoneLens"},
		{GormModel: types.GormModel{ID: 160}, Icon: "/category/charger.jpg", ParentId: 154, AdminId: adminsModel.SuperAdminId, Name: "charger"},

		{GormModel: types.GormModel{ID: 170}, Icon: "/category/cosmetics.png", AdminId: adminsModel.SuperAdminId, Name: "cosmetics"}, // 美容化妆
		{GormModel: types.GormModel{ID: 171}, Icon: "", ParentId: 170, AdminId: adminsModel.SuperAdminId, Name: "facialSkinCare"},
		{GormModel: types.GormModel{ID: 172}, Icon: "/category/facialMask.jpg", ParentId: 171, AdminId: adminsModel.SuperAdminId, Name: "facialMask"},
		{GormModel: types.GormModel{ID: 173}, Icon: "/category/creamLotion.jpg", ParentId: 171, AdminId: adminsModel.SuperAdminId, Name: "creamLotion"},
		{GormModel: types.GormModel{ID: 174}, Icon: "/category/essence.jpg", ParentId: 171, AdminId: adminsModel.SuperAdminId, Name: "essence"},
		{GormModel: types.GormModel{ID: 175}, Icon: "/category/facialCleansing.jpg", ParentId: 171, AdminId: adminsModel.SuperAdminId, Name: "facialCleansing"},
		{GormModel: types.GormModel{ID: 176}, Icon: "/category/sunCream.jpg", ParentId: 171, AdminId: adminsModel.SuperAdminId, Name: "sunCream"},
		{GormModel: types.GormModel{ID: 177}, Icon: "/category/lipPomade.jpg", ParentId: 171, AdminId: adminsModel.SuperAdminId, Name: "lipPomade"},
		{GormModel: types.GormModel{ID: 178}, Icon: "", ParentId: 170, AdminId: adminsModel.SuperAdminId, Name: "hairAndSkinCare"},
		{GormModel: types.GormModel{ID: 179}, Icon: "/category/liquidShampoo.jpg", ParentId: 178, AdminId: adminsModel.SuperAdminId, Name: "liquidShampoo"},
		{GormModel: types.GormModel{ID: 180}, Icon: "/category/hairDyeLiquid.jpg", ParentId: 178, AdminId: adminsModel.SuperAdminId, Name: "hairDyeLiquid"},
		{GormModel: types.GormModel{ID: 181}, Icon: "/category/hairConditioner.jpg", ParentId: 178, AdminId: adminsModel.SuperAdminId, Name: "hairConditioner"},
	}
}
