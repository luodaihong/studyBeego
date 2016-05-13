package controllers

type DatabaseCheck struct{}

func (this *DatabaseCheck) Check() error {
	//dosomething like connect db
	return nil //return errors.New("can't connect database")
}
