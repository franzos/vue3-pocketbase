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

		collection, err := dao.FindCollectionByNameOrId("elm5htenu0fys16")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("7mnwq96h")

		// add
		new_sourceIp := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qzcbthh5",
			"name": "sourceIp",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_sourceIp)
		collection.Schema.AddField(new_sourceIp)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("elm5htenu0fys16")
		if err != nil {
			return err
		}

		// add
		del_sourceIp := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "7mnwq96h",
			"name": "sourceIp",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_sourceIp)
		collection.Schema.AddField(del_sourceIp)

		// remove
		collection.Schema.RemoveField("qzcbthh5")

		return dao.SaveCollection(collection)
	})
}
