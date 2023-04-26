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
			"id": "elm5htenu0fys16",
			"created": "2023-04-26 14:59:38.157Z",
			"updated": "2023-04-26 14:59:38.157Z",
			"name": "submissions",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "nuiokun1",
					"name": "data",
					"type": "json",
					"required": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "gs8zikvw",
					"name": "forwarders",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"collectionId": "nwj43zuhzgfkxeg",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": []
					}
				}
			],
			"indexes": [],
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

		collection, err := dao.FindCollectionByNameOrId("elm5htenu0fys16")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
