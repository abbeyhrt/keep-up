DROP TYPE IF EXISTS oauth_provider;
CREATE TYPE oauth_provider AS ENUM ('facebook', 'google');

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL,
  email VARCHAR(64) UNIQUE NOT NULL,
  avatar_url VARCHAR(512) NOT NULL DEFAULT 'http://http://www.gravatar.com/avatar/?d=mm',
  provider oauth_provider NOT NULL,
  provider_id VARCHAR NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  UNIQUE(provider, provider_id)
);
