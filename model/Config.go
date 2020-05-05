type Config struct{
	App struct{
		Name string `yaml:"name"`
		Port int `yaml:"port"`
		Version string `yaml:"version"`
	}`yaml:"app"`

	Kafka struct{
		Brokers string `yaml:"brokers"`
		GroupId string `yaml:"group_id"`
		Topic string `yaml:"topic"`
	}`yaml:"kafka"`
}