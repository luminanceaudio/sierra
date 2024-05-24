-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_collection_samples" table
CREATE TABLE `new_collection_samples` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `sample_id` text NOT NULL, `collection_id` uuid NOT NULL, CONSTRAINT `collection_samples_source_samples_sample` FOREIGN KEY (`sample_id`) REFERENCES `source_samples` (`uri`) ON DELETE NO ACTION, CONSTRAINT `collection_samples_collections_collection` FOREIGN KEY (`collection_id`) REFERENCES `collections` (`id`) ON DELETE CASCADE);
-- copy rows from old table "collection_samples" to new temporary table "new_collection_samples"
INSERT INTO `new_collection_samples` (`id`, `create_time`, `sample_id`, `collection_id`) SELECT `id`, `create_time`, `sample_id`, `collection_id` FROM `collection_samples`;
-- drop "collection_samples" table after copying rows
DROP TABLE `collection_samples`;
-- rename temporary table "new_collection_samples" to "collection_samples"
ALTER TABLE `new_collection_samples` RENAME TO `collection_samples`;
-- create index "collectionsample_sample_id" to table: "collection_samples"
CREATE INDEX `collectionsample_sample_id` ON `collection_samples` (`sample_id`);
-- create index "collectionsample_collection_id" to table: "collection_samples"
CREATE INDEX `collectionsample_collection_id` ON `collection_samples` (`collection_id`);
-- create index "collectionsample_sample_id_collection_id" to table: "collection_samples"
CREATE UNIQUE INDEX `collectionsample_sample_id_collection_id` ON `collection_samples` (`sample_id`, `collection_id`);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
