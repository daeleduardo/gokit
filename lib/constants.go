package lib

//Application constants, only default value will used.
type Constants struct {
	LogDirectory string `default:"/var/log/gokit/"`
}

//return application contants
func  GetConstants() Constants  {
	return Constants{}
}