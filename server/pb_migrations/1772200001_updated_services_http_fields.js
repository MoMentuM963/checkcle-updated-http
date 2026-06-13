/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_863811952")

  // add method field
  collection.fields.addAt(999, new TextField({
    "id": "field_method",
    "name": "method",
    "required": false,
    "presentable": false
  }))

  // add headers field
  collection.fields.addAt(999, new TextField({
    "id": "field_headers",
    "name": "headers",
    "required": false,
    "presentable": false
  }))

  // add body field
  collection.fields.addAt(999, new TextField({
    "id": "field_body",
    "name": "body",
    "required": false,
    "presentable": false
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_863811952")

  // rollback - remove the fields
  collection.fields.removeById("field_method")
  collection.fields.removeById("field_headers")
  collection.fields.removeById("field_body")

  return app.save(collection)
})
