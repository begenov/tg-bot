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

CREATE TABLE IF NOT EXISTS vacancies (
    id INTEGER PRIMARY KEY,
    user_profile_id INT,
    company VARCHAR(255),
    bin VARCHAR(20),
    sphere VARCHAR(255),
    position VARCHAR(255),
    salary VARCHAR(255),
    requirements TEXT,
    responsibilities TEXT
);
