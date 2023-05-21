CREATE TABLE IF NOT EXISTS user_profile (
  id INTEGER PRIMARY KEY,
  chat_id INT NOT NULL, 
  name TEXT,
  phone TEXT, 
  language TEXT, 
  role INT(1) NOT NULL,
  age INT(2),
  gender INT(1)
);


CREATE TABLE IF NOT EXISTS job_seeker (
		id INTEGER PRIMARY KEY,
    user_profile_id INT,
		sphere TEXT,
		profession TEXT,
		salary TEXT,
    FOREIGN KEY (user_profile_id) REFERENCES user_profile(id)
);