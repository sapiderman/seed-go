package connector_test

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/sapiderman/seed-go/internal/config"
	"github.com/sapiderman/seed-go/internal/connector"

	"github.com/jmoiron/sqlx"
)

func Test_DropAllTables(t *testing.T) {

	psgqlConnectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Get("psql.host"), config.Get("psql.port"), config.Get("psql.user"), config.Get("psql.pass"), config.Get("psql.dbname"))
	dbtest, err := sqlx.Connect("postgres", psgqlConnectStr)
	if err != nil {
		t.Fatal(err)
	}
	defer dbtest.Close()

	db := connector.DbPool{Db: dbtest}

	err = db.DropAllTables()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_CreateAllTables(t *testing.T) {

	var dbtest *sqlx.DB

	psgqlConnectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Get("psql.host"), config.Get("psql.port"), config.Get("psql.user"), config.Get("psql.pass"), config.Get("psql.dbname"))
	dbtest, err := sqlx.Connect("postgres", psgqlConnectStr)
	if err != nil {
		t.Fatal(err)
	}
	defer dbtest.Close()

	db := connector.DbPool{Db: dbtest}

	err = db.DropAllTables()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_InsertDevice(t *testing.T) {

	psgqlConnectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Get("psql.host"), config.Get("psql.port"), config.Get("psql.user"), config.Get("psql.pass"), config.Get("psql.dbname"))
	dbtest, err := sqlx.Connect("postgres", psgqlConnectStr)
	if err != nil {
		t.Fatal(err)
	}
	defer dbtest.Close()

	p := connector.DbPool{Db: dbtest}

	// sqlStatement := `INSERT INTO devices (id, created_at, updated_at, deleted_at, phone_brand, phone_model, push_id, device_id)
	//  VALUES (:id, :created_at, :updated_at, :deleted_at, :phone_brand, :phone_model, :push_id, :device_id)`
	// _, err = dbtest.Exec(sqlStatement)

	testDev := connector.Device{
		ID:         "1",
		CreatedAt:  "2004-10-19 10:23:54+02",
		UpdatedAt:  "2004-10-19 10:23:54+02",
		DeletedAt:  "2004-10-19 10:23:54+02",
		PhoneBrand: "test_brand",
		PhoneModel: "test_model",
		PushID:     "test_push_id",
		DeviceID:   "test_device_id",
	}

	_, err = p.Db.NamedExec(connector.InsertDeviceSQL, testDev)
	err = p.InsertDevice(&testDev)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_ListAllDevices(t *testing.T) {
	psgqlConnectStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.Get("psql.host"), config.Get("psql.port"), config.Get("psql.user"), config.Get("psql.pass"), config.Get("psql.dbname"))
	dbtest, err := sqlx.Connect("postgres", psgqlConnectStr)
	if err != nil {
		t.Fatal(err)
	}
	defer dbtest.Close()
	p := connector.DbPool{Db: dbtest}

	// ctx := context.Background()
	dev := []connector.Device{}
	dev, err = p.ListAllDevices()
	if err != nil {
		t.Fatal(err)
	}

	if dev[0].ID == "" {
		t.Fail()
	}

}
