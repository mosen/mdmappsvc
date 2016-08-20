CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- This table stores individual flat pkgs for installation and their relative paths
-- It does not store information about the bundles that are part of the distribution.
CREATE TABLE IF NOT EXISTS applications (
  uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  source_uuid UUID REFERENCES sources(uuid) ON DELETE SET NULL,
  source_path TEXT,
  icon_source_uuid UUID REFERENCES sources(uuid) ON DELETE SET NULL,
  icon_source_path TEXT,
  icon_needs_shine BOOLEAN NOT NULL DEFAULT FALSE,
  size BIGINT NOT NULL DEFAULT 0,

  -- not incl sub packages if a distribution
  bundle_identifier TEXT,
  bundle_version TEXT,

  title TEXT,
  kind TEXT NOT NULL DEFAULT 'software', -- may be books later
  subtitle TEXT -- Can be organization
);
