CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS md5_checksums (
  uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  application_uuid UUID REFERENCES applications(uuid) ON DELETE CASCADE,
  is_chunk BOOLEAN DEFAULT TRUE,
  chunk_size BIGINT DEFAULT 10485760,
  chunk_idx INT NOT NULL,
  checksum TEXT NOT NULL
);

CREATE UNIQUE INDEX uq_md5_checksum_application_chunk ON md5_checksums(application_uuid, chunk_idx);
