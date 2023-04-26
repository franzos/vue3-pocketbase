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

		// add
		new_spamScore := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "juyabbvj",
			"name": "spamScore",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), new_spamScore)
		collection.Schema.AddField(new_spamScore)

		// add
		new_userFlaggedAsSpam := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "bxqxo8vi",
			"name": "userFlaggedAsSpam",
			"type": "bool",
			"required": false,
			"unique": false,
			"options": {}
		}`), new_userFlaggedAsSpam)
		collection.Schema.AddField(new_userFlaggedAsSpam)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("elm5htenu0fys16")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("juyabbvj")

		// remove
		collection.Schema.RemoveField("bxqxo8vi")

		return dao.SaveCollection(collection)
	})
}
