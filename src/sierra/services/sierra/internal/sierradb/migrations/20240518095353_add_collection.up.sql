-- create "collections" table
CREATE TABLE `collections` (`id` uuid NOT NULL, `create_time` datetime NOT NULL, `name` text NOT NULL, PRIMARY KEY (`id`));
-- create index "collection_id" to table: "collections"
CREATE UNIQUE INDEX `collection_id` ON `collections` (`id`);
-- create index "collection_name" to table: "collections"
CREATE INDEX `collection_name` ON `collections` (`name`);
-- create "collection_samples" table
CREATE TABLE `collection_samples` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `sample_id` text NOT NULL, `collection_id` uuid NOT NULL, CONSTRAINT `collection_samples_source_samples_sample` FOREIGN KEY (`sample_id`) REFERENCES `source_samples` (`uri`) ON DELETE NO ACTION, CONSTRAINT `collection_samples_collections_collection` FOREIGN KEY (`collection_id`) REFERENCES `collections` (`id`) ON DELETE NO ACTION);
-- create index "collectionsample_sample_id" to table: "collection_samples"
CREATE INDEX `collectionsample_sample_id` ON `collection_samples` (`sample_id`);
-- create index "collectionsample_collection_id" to table: "collection_samples"
CREATE INDEX `collectionsample_collection_id` ON `collection_samples` (`collection_id`);
-- create index "collectionsample_sample_id_collection_id" to table: "collection_samples"
CREATE UNIQUE INDEX `collectionsample_sample_id_collection_id` ON `collection_samples` (`sample_id`, `collection_id`);
