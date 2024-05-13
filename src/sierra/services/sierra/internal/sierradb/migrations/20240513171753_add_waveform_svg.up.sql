-- add column "waveform_storage_path" to table: "samples"
ALTER TABLE `samples` ADD COLUMN `waveform_storage_path` text NULL;
