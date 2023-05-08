package db

type Business struct {
	ID      int    `db: "id"`
	Name    string `db: "name`
	Address string `db: "address"`
	Phone   string `db: "phone"`
}
