package dbdiff

type DatabaseInfo struct {
	SchemaName  string      `db:"schemaname"`
	TableName   string      `db:"tablename"`
	TableOwner  string      `db:"tableowner"`
	TableSpace  interface{} `db:"tablespace"`
	HasIndexes  bool        `db:"hasindexes"`
	HasRules    bool        `db:"hasrules"`
	HasTriggers bool        `db:"hastriggers"`
	RowSecurity bool        `db:"rowsecurity"`
}

// select * from pg_catalog.pg_tables where schemaname = 'public'

type TableInfo struct {
	TableCatalog           string      `db:"table_catalog"`
	TableSchema            string      `db:"table_schema"`
	TableName              string      `db:"table_name"`
	ColumnName             string      `db:"column_name"`
	OrdinalPosition        int         `db:"ordinal_position"`
	ColumnDefault          string      `db:"column_default"`
	IsNullable             string      `db:"is_nullable"`
	DataType               string      `db:"data_type"`
	CharacterMaximumLength interface{} `db:"character_maximum_length"`
	CharacterOctetLength   interface{} `db:"character_octet_length"`
	NumericPrecision       int         `db:"numeric_precision"`
	NumericPrecisionRadix  int         `db:"numeric_precision_radix"`
	NumericScale           int         `db:"numeric_scale"`
	DatetimePrecision      interface{} `db:"datetime_precision"`
	IntervalType           interface{} `db:"interval_type"`
	IntervalPrecision      interface{} `db:"interval_precision"`
	CharacterSetCatalog    interface{} `db:"character_set_catalog"`
	CharacterSetSchema     interface{} `db:"character_set_schema"`
	CharacterSetName       interface{} `db:"character_set_name"`
	CollationCatalog       interface{} `db:"collation_catalog"`
	CollationSchema        interface{} `db:"collation_schema"`
	CollationName          interface{} `db:"collation_name"`
	DomainCatalog          interface{} `db:"domain_catalog"`
	DomainSchema           interface{} `db:"domain_schema"`
	DomainName             interface{} `db:"domain_name"`
	UdtCatalog             string      `db:"udt_catalog"`
	UdtSchema              string      `db:"udt_schema"`
	UdtName                string      `db:"udt_name"`
	ScopeCatalog           interface{} `db:"scope_catalog"`
	ScopeSchema            interface{} `db:"scope_schema"`
	ScopeName              interface{} `db:"scope_name"`
	MaximumCardinality     interface{} `db:"maximum_cardinality"`
	DtdIdentifier          string      `db:"dtd_identifier"`
	IsSelfReferencing      string      `db:"is_self_referencing"`
	IsIdentity             string      `db:"is_identity"`
	IdentityGeneration     interface{} `db:"identity_generation"`
	IdentityStart          interface{} `db:"identity_start"`
	IdentityIncrement      interface{} `db:"identity_increment"`
	IdentityMaximum        interface{} `db:"identity_maximum"`
	IdentityMinimum        interface{} `db:"identity_minimum"`
	IdentityCycle          string      `db:"identity_cycle"`
	IsGenerated            string      `db:"is_generated"`
	GenerationExpression   interface{} `db:"generation_expression"`
	IsUpdatable            string      `db:"is_updatable"`
}

// select * from information_schema.columns where table_name = ?
