-- reverse: create index "sourcesample_uri" to table: "source_samples"
DROP INDEX `sourcesample_uri`;
-- reverse: create "source_samples" table
DROP TABLE `source_samples`;
-- reverse: create index "source_uri" to table: "sources"
DROP INDEX `source_uri`;
-- reverse: create "sources" table
DROP TABLE `sources`;
-- reverse: create index "sample_sha256" to table: "samples"
DROP INDEX `sample_sha256`;
-- reverse: create "samples" table
DROP TABLE `samples`;
