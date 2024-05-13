-- add column "relative_path" to table: "source_samples"
ALTER TABLE `source_samples` ADD COLUMN `relative_path` text NOT NULL;
