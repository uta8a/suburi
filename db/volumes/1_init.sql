CREATE TABLE userinfo (
  id SERIAL NOT NULL,
  username varchar(30) NOT NULL,
  password_hash text NOT NULL,
  user_type text NOT NULL,
  UNIQUE (id),
  UNIQUE (username),
  PRIMARY KEY (id)
);


INSERT INTO userinfo (username, password_hash, user_type) VALUES (
  'mofuadmin',
  'REDACTED1',
  'admin'
);

INSERT INTO userinfo (username, password_hash, user_type) VALUES (
  'mofutester',
  'REDACTED2',
  'tester'
);

INSERT INTO userinfo (username, password_hash, user_type) VALUES (
  'mofuchan',
  'REDACTED3',
  'player'
);
