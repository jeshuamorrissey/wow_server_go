CREATE TABLE Account
(
  id INTEGER PRIMARY KEY,

  -- Basic account information.
  name VARCHAR(255) NOT NULL UNIQUE,

  -- Authentication information.
  verifier CHAR(64) NOT NULL,
  salt CHAR(64) NOT NULL
);

CREATE TABLE Realm
(
  id INTEGER PRIMARY KEY,

  -- Basic realm information.
  name VARCHAR(255) NOT NULL UNIQUE,
  host VARCHAR(255) NOT NULL
);

CREATE TABLE Character
(
  id INTEGER PRIMARY KEY,

  -- Basic character information.
  name VARCHAR(255) NOT NULL,

  -- Links.
  accountID INT,
  realmID INT,

  FOREIGN KEY(accountID) REFERENCES Account(id),
  FOREIGN KEY(realmID) REFERENCES Realm(id),

  -- Constraints.
  CONSTRAINT UC_Character_Realm UNIQUE (name, realmID)
);
