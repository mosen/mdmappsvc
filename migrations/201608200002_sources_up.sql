CREATE TABLE IF NOT EXISTS sources (
  uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  public_uri TEXT NOT NULL,
  type_uuid UUID REFERENCES source_types(uuid) NOT NULL
);

