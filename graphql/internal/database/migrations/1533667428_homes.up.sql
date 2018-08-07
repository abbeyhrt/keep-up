CREATE TABLE homes (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(64) NOT NULL,
  description VARCHAR(64),
  avatar_url VARCHAR(512) NOT NULL DEFAULT 'http://http://www.gravatar.com/avatar/?d=mm',
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);
