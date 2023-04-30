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
			"id": "pun0uiel75hr5f4",
			"created": "2023-04-28 20:52:18.476Z",
			"updated": "2023-04-28 20:52:18.476Z",
			"name": "logs",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "2dxtdrhe",
					"name": "message",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "r4z2u2vm",
					"name": "topic",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "og5vtymb",
					"name": "data",
					"type": "json",
					"required": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "kq7wvyrm",
					"name": "forwarder",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"collectionId": "nwj43zuhzgfkxeg",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "jz7zye8u",
					"name": "submission",
					"type": "relation",
					"required": false,
					"unique": false,
					"options": {
						"collectionId": "elm5htenu0fys16",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
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

		collection, err := dao.FindCollectionByNameOrId("pun0uiel75hr5f4")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
