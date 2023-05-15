CREATE TABLE IF NOT EXISTS user_profile (
  id INTEGER PRIMARY KEY, 
  name TEXT,
  phone TEXT, 
  language TEXT, 
  role INT(1) NOT NULL,
  age INT(2),
  gender INT(1)
);