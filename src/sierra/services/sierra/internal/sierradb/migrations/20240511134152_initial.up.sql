-- create "samples" table
CREATE TABLE `samples` (`sha256` text NOT NULL, `format` text NULL, `length` integer NULL, PRIMARY KEY (`sha256`));
-- create index "sample_sha256" to table: "samples"
CREATE UNIQUE INDEX `sample_sha256` ON `samples` (`sha256`);
-- create "sources" table
CREATE TABLE `sources` (`uri` text NOT NULL, PRIMARY KEY (`uri`));
-- create "source_samples" table
CREATE TABLE `source_samples` (`relative_path` text NOT NULL, `source` text NULL, `sample` text NULL, PRIMARY KEY (`relative_path`), CONSTRAINT `source_samples_sources_source` FOREIGN KEY (`source`) REFERENCES `sources` (`uri`) ON DELETE SET NULL, CONSTRAINT `source_samples_samples_sample` FOREIGN KEY (`sample`) REFERENCES `samples` (`sha256`) ON DELETE SET NULL);
