package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "umwoio6covrgs3g",
			"created": "2023-04-30 14:30:07.195Z",
			"updated": "2023-04-30 14:30:07.195Z",
			"name": "blocklist",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "6zjvbi2q",
					"name": "value",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_EBzBXAc` + "`" + ` ON ` + "`" + `blocklist` + "`" + ` (` + "`" + `value` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("umwoio6covrgs3g")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
