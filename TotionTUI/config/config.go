package config

type HTTPServer struct{
	Address string ``
}

type Config struct {
	StoragePath string `` 
	HTTPServer
}

func MustExec() (*Config, error)  {
	


	return nil, nil
}