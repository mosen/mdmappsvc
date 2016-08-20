CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Types such as: Local, Munki, HTTP (r/o), S3, Rackspace etc.
CREATE TABLE IF NOT EXISTS source_types (
  uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name TEXT
);

INSERT INTO source_types (name) VALUES ('Munki');
INSERT INTO source_types (name) VALUES ('Local Filesystem');
INSERT INTO source_types (name) VALUES ('Generic HTTPS');
INSERT INTO source_types (name) VALUES ('Amazon S3');
