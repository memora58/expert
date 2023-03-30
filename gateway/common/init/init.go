package init

func init() {
	InitConfig()

	//InitMysql()

	// 初始化翻译
	if err := InitTrans("zh"); err != nil {
		panic(err)
	}
}
