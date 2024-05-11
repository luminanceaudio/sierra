# Design

## Sources

Sources indicate where to look for sounds. This can be a local directory, a Google Drive directory or a URL directing to an S3 compatible file server.

## Analyzers

Analyzers are modules that process an audio sample and analyze it to output metadata.

### Transient Analyzer

Analyzes the peaks of a sample and recgonizes whether it is "soft" or "hard" based on the amount of high frequencies.

### Length Analyzer

Measures the "actual length" of a sample without trailing silence to determine its length.

### BPM Analyzer

Measure the approximate BPM of a sample.

### Note analyzer

Analyze the

## Users

Share favorite sounds with each other.

## Plugins

Free plugins repository.

## Server

A server could optionally be implemented. It would receive samples sent in the background from clients, analyze them and store their metadata. Other clients could later query samples metadata to avoid locally analyzing samples.

This could be presented as "Remote Indexing" vs "Local Indexing", the former being faster but requiring internet connectivity.
