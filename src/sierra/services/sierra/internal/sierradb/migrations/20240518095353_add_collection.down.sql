-- reverse: create index "collectionsample_sample_id_collection_id" to table: "collection_samples"
DROP INDEX `collectionsample_sample_id_collection_id`;
-- reverse: create index "collectionsample_collection_id" to table: "collection_samples"
DROP INDEX `collectionsample_collection_id`;
-- reverse: create index "collectionsample_sample_id" to table: "collection_samples"
DROP INDEX `collectionsample_sample_id`;
-- reverse: create "collection_samples" table
DROP TABLE `collection_samples`;
-- reverse: create index "collection_name" to table: "collections"
DROP INDEX `collection_name`;
-- reverse: create index "collection_id" to table: "collections"
DROP INDEX `collection_id`;
-- reverse: create "collections" table
DROP TABLE `collections`;
