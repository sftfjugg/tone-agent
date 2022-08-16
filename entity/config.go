package entity

type Config struct {
	TSN   string `json:"tsn"`
	Mode  string `json:"mode"`
	Proxy string `json:"proxy"`
}

//func (c *Config) GetConf() *Config {
//	yamlFile, err := ioutil.ReadFile("toneagent.config.yaml")
//	if err != nil {
//		log.Printf("yamlFile.Get err   #%v ", err)
//	}
//
//	err = yaml.Unmarshal(yamlFile, c)
//	if err != nil {
//		log.Fatalf("Unmarshal: %v", err)
//	}
//	return c
//}