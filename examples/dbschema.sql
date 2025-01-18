CREATE TABLE IF NOT EXISTS members (
  id SERIAL PRIMARY KEY,
  nickname VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(60) NOT NULL, -- 60 means the length of password hash-sum, not the plain text password
  member_uuid UUID NOT NULL,
  join_date TIMESTAMP NOT NULL,
  sex VARCHAR(30),
  about TEXT
);

CREATE TABLE IF NOT EXISTS positions (
  id SERIAL PRIMARY KEY,
  position TEXT
);

CREATE TABLE IF NOT EXISTS admins (
  id SERIAL PRIMARY KEY,
  nickname VARCHAR(50) NOT NULL,
  password VARCHAR(60) NOT NULL, -- 60 means the length of password hash-sum, not the plain text password
  admin_uuid UUID NOT NULL,
  position_id INTEGER NOT NULL REFERENCES positions(id),
  join_date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS invite_tokens (
  id SERIAL PRIMARY KEY,
  token VARCHAR(30)
);
