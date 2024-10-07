package initialize

import (
	"fmt"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		fmt.Println("Unable to decode configuration %v", err)
	}
}