package configs

import (
	"flag"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v3"
)

// Config declares config file structure
type Config struct {
	Server struct {
		DebugMode bool   `yaml:"debug_mode"`
		Host      string `yaml:"host"`
		Port      uint   `yaml:"port"`
	} `yaml:"server"`

	StatsD struct {
		Mocked bool   `yaml:"mocked"`
		Host   string `yaml:"host"`
		Port   uint   `yaml:"port"`
	} `yaml:"statsd"`
}

// NewConfig returns a new instance of the structure
func NewConfig() *Config {
	return new(Config)
}

// NewConfigMock creates a new instance of tests
func NewConfigMock() *Config {
	cnf := new(Config)
	cnf.StatsD.Mocked = true

	return cnf
}

// ParseFromFile parsing config file to structure
func (c *Config) ParseFromFile() {
	path := c.getConfigPath()

	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configBytes, c)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Config) getConfigPath() string {
	var (
		cfgDir     string
		cfgName    string
		configPath string
	)

	flag.StringVar(&cfgDir, "cfg-dir", "./", "Path to config folder.")
	flag.StringVar(&cfgName, "cfg-name", "config.yaml", "Config yaml file name.")
	flag.Parse()

	configPath = cfgDir + cfgName
	if configPath == "" {
		log.Fatal(`
			Config path is empty!
			Use --cfg-dir and --cfg-name flags.
		`)
	}

	return configPath
}
