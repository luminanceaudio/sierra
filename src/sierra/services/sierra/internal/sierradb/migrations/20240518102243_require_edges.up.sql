-- disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- create "new_source_samples" table
CREATE TABLE `new_source_samples` (`uri` text NOT NULL, `relative_path` text NOT NULL, `filename` text NOT NULL, `source` text NOT NULL, `sample` text NOT NULL, PRIMARY KEY (`uri`), CONSTRAINT `source_samples_sources_source` FOREIGN KEY (`source`) REFERENCES `sources` (`uri`) ON DELETE NO ACTION, CONSTRAINT `source_samples_samples_sample` FOREIGN KEY (`sample`) REFERENCES `samples` (`sha256`) ON DELETE NO ACTION);
-- copy rows from old table "source_samples" to new temporary table "new_source_samples"
INSERT INTO `new_source_samples` (`uri`, `relative_path`, `filename`, `source`, `sample`) SELECT `uri`, `relative_path`, `filename`, `source`, `sample` FROM `source_samples`;
-- drop "source_samples" table after copying rows
DROP TABLE `source_samples`;
-- rename temporary table "new_source_samples" to "source_samples"
ALTER TABLE `new_source_samples` RENAME TO `source_samples`;
-- create index "sourcesample_uri" to table: "source_samples"
CREATE UNIQUE INDEX `sourcesample_uri` ON `source_samples` (`uri`);
-- create index "sourcesample_relative_path" to table: "source_samples"
CREATE INDEX `sourcesample_relative_path` ON `source_samples` (`relative_path`);
-- create index "sourcesample_filename" to table: "source_samples"
CREATE INDEX `sourcesample_filename` ON `source_samples` (`filename`);
-- enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
