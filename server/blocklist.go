package main

import (
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Blocklist)(nil)

type Blocklist struct {
	models.BaseModel

	Value string `db:"value" json:"value"`
}

type BlocklistApi struct {
	dao  *daos.Dao
	data Blocklist
}

func (m *Blocklist) TableName() string {
	return "blocklist"
}

func BlocklistQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Blocklist{})
}

func FindBlocklistRecordByValue(dao *daos.Dao, value string) (*Blocklist, error) {
	blocklistRecord := &Blocklist{}

	err := BlocklistQuery(dao).
		AndWhere(dbx.NewExp("LOWER(value)={:value}", dbx.Params{"value": value})).
		Limit(1).
		One(blocklistRecord)

	if err != nil {
		return nil, err
	}

	return blocklistRecord, nil
}

func CreateBlocklistRecord(api BlocklistApi) error {
	newRecord := &Blocklist{
		Value: api.data.Value,
	}

	if err := api.dao.Save(newRecord); err != nil {
		log.Printf("Error saving blocklist record: %v\n", err)
		return err
	}

	return nil
}

func DeleteBlocklistRecord(api BlocklistApi) error {
	record, err := FindBlocklistRecordByValue(api.dao, api.data.Value)
	if err != nil {
		return err
	}

	return api.dao.Delete(record)
}
