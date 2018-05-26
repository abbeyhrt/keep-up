CREATE TABLE sessions_table (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  userID UUID REFERENCES users(id),
  created_at TIMESTAMP DEFAULT NOW()
);
