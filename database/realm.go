package database

const (
	realmGetAll = "SELECT id, name, host FROM Realm"
	realmInsert = "INSERT INTO Realm(name, host) VALUES(?, ?)"
	realmUpdate = "UPDATE Realm(name, host) VALUES(?, ?) WHERE id = ?"
)

// Realm represents a worldserver the client can connect to.
type Realm struct {
	ID   int64
	Name string
	Host string
}

// NewRealm makes a new realm.
func NewRealm(name, host string) *Realm {
	realm := new(Realm)

	realm.ID = -1
	realm.Name = name
	realm.Host = host

	return realm
}

// GetAllRealms retreives all realms from the database.
func GetAllRealms() ([]*Realm, error) {
	stmt, err := db.Prepare(realmGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	realms := make([]*Realm, 0)
	for rows.Next() {
		realm := new(Realm)
		err := rows.Scan(&realm.ID, &realm.Name, &realm.Host)
		if err != nil {
			return nil, err
		}

		realms = append(realms, realm)
	}

	return realms, nil
}

// Save will write the realm to the database. This will also insert it if it
// doesn't already exist.
func (r *Realm) Save() error {
	if r.ID < 0 {
		stmt, err := db.Prepare(realmInsert)
		if err != nil {
			return err
		}

		result, err := stmt.Exec(r.Name, r.Host)
		if err != nil {
			return err
		}

		r.ID, err = result.LastInsertId()
		if err != nil {
			return err
		}
	} else {
		stmt, err := db.Prepare(accountUpdate)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(r.Name, r.Host, r.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
