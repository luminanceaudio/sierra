-- create "samples" table
CREATE TABLE `samples` (`sha256` text NOT NULL, `format` text NULL, `duration` integer NULL, PRIMARY KEY (`sha256`));
-- create index "sample_sha256" to table: "samples"
CREATE UNIQUE INDEX `sample_sha256` ON `samples` (`sha256`);
-- create "sources" table
CREATE TABLE `sources` (`uri` text NOT NULL, `create_time` datetime NOT NULL, PRIMARY KEY (`uri`));
-- create index "source_uri" to table: "sources"
CREATE UNIQUE INDEX `source_uri` ON `sources` (`uri`);
-- create "source_samples" table
CREATE TABLE `source_samples` (`uri` text NOT NULL, `relative_path` text NOT NULL, `source` text NULL, `sample` text NULL, PRIMARY KEY (`uri`), CONSTRAINT `source_samples_sources_source` FOREIGN KEY (`source`) REFERENCES `sources` (`uri`) ON DELETE SET NULL, CONSTRAINT `source_samples_samples_sample` FOREIGN KEY (`sample`) REFERENCES `samples` (`sha256`) ON DELETE SET NULL);
-- create index "sourcesample_uri" to table: "source_samples"
CREATE UNIQUE INDEX `sourcesample_uri` ON `source_samples` (`uri`);
-- create index "sourcesample_relative_path" to table: "source_samples"
CREATE INDEX `sourcesample_relative_path` ON `source_samples` (`relative_path`);
