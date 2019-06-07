package config

type (
	// Config datos de conexion
	Config struct {
		Server   string
		Database string
	}
)

func (c *Config) Read() {
	c.Server = "mongodb://jmmrcp:J538MTUSbg3v3Vh@ds263876.mlab.com:63876"
	c.Database = "justicia"
}
