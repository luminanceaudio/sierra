-- reverse: create index "collectionsample_sample_id_collection_id" to table: "collection_samples"
DROP INDEX `collectionsample_sample_id_collection_id`;
-- reverse: create index "collectionsample_collection_id" to table: "collection_samples"
DROP INDEX `collectionsample_collection_id`;
-- reverse: create index "collectionsample_sample_id" to table: "collection_samples"
DROP INDEX `collectionsample_sample_id`;
-- reverse: create "new_collection_samples" table
DROP TABLE `new_collection_samples`;
