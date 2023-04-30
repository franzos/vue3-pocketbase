package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pun0uiel75hr5f4")
		if err != nil {
			return err
		}

		// add
		new_level := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "zerqjxjk",
			"name": "level",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), new_level)
		collection.Schema.AddField(new_level)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("pun0uiel75hr5f4")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("zerqjxjk")

		return dao.SaveCollection(collection)
	})
}
