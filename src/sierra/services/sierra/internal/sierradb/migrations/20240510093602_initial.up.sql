-- create "sample_files" table
CREATE TABLE `sample_files` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `sha256` text NOT NULL, `format` text NULL, `length` integer NULL);
-- create index "samplefile_sha256" to table: "sample_files"
CREATE UNIQUE INDEX `samplefile_sha256` ON `sample_files` (`sha256`);
