// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SamplesColumns holds the columns for the "samples" table.
	SamplesColumns = []*schema.Column{
		{Name: "sha256", Type: field.TypeString, Unique: true},
		{Name: "format", Type: field.TypeString, Nullable: true},
		{Name: "length", Type: field.TypeInt64, Nullable: true},
		{Name: "waveform_svg", Type: field.TypeBytes, Nullable: true},
	}
	// SamplesTable holds the schema information for the "samples" table.
	SamplesTable = &schema.Table{
		Name:       "samples",
		Columns:    SamplesColumns,
		PrimaryKey: []*schema.Column{SamplesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "sample_sha256",
				Unique:  true,
				Columns: []*schema.Column{SamplesColumns[0]},
			},
		},
	}
	// SourcesColumns holds the columns for the "sources" table.
	SourcesColumns = []*schema.Column{
		{Name: "uri", Type: field.TypeString, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
	}
	// SourcesTable holds the schema information for the "sources" table.
	SourcesTable = &schema.Table{
		Name:       "sources",
		Columns:    SourcesColumns,
		PrimaryKey: []*schema.Column{SourcesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "source_uri",
				Unique:  true,
				Columns: []*schema.Column{SourcesColumns[0]},
			},
		},
	}
	// SourceSamplesColumns holds the columns for the "source_samples" table.
	SourceSamplesColumns = []*schema.Column{
		{Name: "uri", Type: field.TypeString, Unique: true},
		{Name: "source", Type: field.TypeString, Nullable: true},
		{Name: "sample", Type: field.TypeString, Nullable: true},
	}
	// SourceSamplesTable holds the schema information for the "source_samples" table.
	SourceSamplesTable = &schema.Table{
		Name:       "source_samples",
		Columns:    SourceSamplesColumns,
		PrimaryKey: []*schema.Column{SourceSamplesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "source_samples_sources_source",
				Columns:    []*schema.Column{SourceSamplesColumns[1]},
				RefColumns: []*schema.Column{SourcesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "source_samples_samples_sample",
				Columns:    []*schema.Column{SourceSamplesColumns[2]},
				RefColumns: []*schema.Column{SamplesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "sourcesample_uri",
				Unique:  true,
				Columns: []*schema.Column{SourceSamplesColumns[0]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SamplesTable,
		SourcesTable,
		SourceSamplesTable,
	}
)

func init() {
	SourceSamplesTable.ForeignKeys[0].RefTable = SourcesTable
	SourceSamplesTable.ForeignKeys[1].RefTable = SamplesTable
}
