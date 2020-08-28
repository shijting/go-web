package database

type Load struct {
}

func NewMigration() *Load {
	return &Load{}
}

func (*Load) Migration() {

}
