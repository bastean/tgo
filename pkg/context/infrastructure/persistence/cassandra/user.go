package cassandra

import (
	"encoding/json"
	"fmt"

	"github.com/bastean/tgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/tgo/pkg/context/domain/errors"
	"github.com/bastean/tgo/pkg/context/domain/repository"
	"github.com/gocql/gocql"
)

type User struct {
	*Cassandra
	Table string
}

func (cassandra *User) Save(user *user.User) error {
	exists, _ := cassandra.Search(&repository.UserSearchCriteria{
		Username: user.Username,
	})

	if exists != nil {
		return errors.NewAlreadyExist(&errors.Bubble{
			Where: "Save",
			What:  fmt.Sprintf("%s already registered", user.Username.Value),
			Why: errors.Meta{
				"Username": user.Username.Value,
			},
		})
	}

	new, err := json.Marshal(user.ToPrimitive())

	new = AddExtraKeyQuotes(new)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Save",
			What:  "Failure to JSON encoding",
			Why: errors.Meta{
				"Username": user.Username.Value,
			},
			Who: err,
		})
	}

	insert := fmt.Sprintf("INSERT INTO %s.%s JSON ?", cassandra.Keyspace, cassandra.Table)

	err = cassandra.Session.Query(insert, new).Exec()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Save",
			What:  "Failure to save a user",
			Why: errors.Meta{
				"Username": user.Username.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (cassandra *User) Update(user *user.User) error {
	updated, err := json.Marshal(user.ToPrimitive())

	updated = AddExtraKeyQuotes(updated)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Save",
			What:  "Failure to JSON encoding",
			Why: errors.Meta{
				"Username": user.Username.Value,
			},
			Who: err,
		})
	}

	insert := fmt.Sprintf("INSERT INTO %s.%s JSON ? DEFAULT UNSET", cassandra.Keyspace, cassandra.Table)

	err = cassandra.Session.Query(insert, updated).Exec()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "Failure to update a user",
			Why: errors.Meta{
				"Username": user.Username.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (cassandra *User) Delete(username *user.Username) error {
	delete := fmt.Sprintf("DELETE FROM %s.%s WHERE \"Username\" = ?", cassandra.Keyspace, cassandra.Table)

	err := cassandra.Session.Query(delete, username.Value).Exec()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Delete",
			What:  "Failure to delete a user",
			Why: errors.Meta{
				"Username": username.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (cassandra *User) Search(criteria *repository.UserSearchCriteria) (*user.User, error) {
	result := make(map[string]any)

	choose := fmt.Sprintf("SELECT JSON * FROM %s.%s WHERE \"Username\" = ?", cassandra.Keyspace, cassandra.Table)

	err := cassandra.Session.Query(choose, criteria.Username.Value).MapScan(result)

	switch {
	case errors.Is(err, gocql.ErrNotFound):
		return nil, errors.NewNotExist(&errors.Bubble{
			Where: "Search",
			What:  fmt.Sprintf("%s not found", criteria.Username.Value),
			Why: errors.Meta{
				"Username": criteria.Username.Value,
			},
			Who: err,
		})
	case err != nil:
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "Failure to search a user",
			Why: errors.Meta{
				"Username": criteria.Username.Value,
			},
			Who: err,
		})
	}

	primitive := new(user.Primitive)

	encoded := RemoveExtraKeyQuotes(result["[json]"].(string))

	err = json.Unmarshal(encoded, primitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "Failure to JSON decoding",
			Why: errors.Meta{
				"Username": criteria.Username.Value,
				"JSON":     encoded,
			},
			Who: err,
		})
	}

	found, err := user.FromPrimitive(primitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "Failure to create an User from a Primitive",
			Why: errors.Meta{
				"Primitive": primitive,
				"Username":  criteria.Username.Value,
			},
			Who: err,
		})
	}

	return found, nil
}

func OpenUser(session *Cassandra, table string) (repository.User, error) {
	types := &struct {
		Coins string
	}{
		Coins: "coins",
	}

	create := fmt.Sprintf("CREATE TYPE IF NOT EXISTS %s.%s (\"Currency\" text, \"List\" list<text>)", session.Keyspace, types.Coins)

	err := session.Query(create).Exec()

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "OpenUser",
			What:  "Failure to create a Type",
			Why: errors.Meta{
				"Type": types.Coins,
			},
			Who: err,
		})
	}

	create = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (\"Username\" text PRIMARY KEY, \"Coins\" frozen<%s>)", session.Keyspace, table, types.Coins)

	err = session.Query(create).Exec()

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "OpenUser",
			What:  "Failure to create a Table",
			Why: errors.Meta{
				"Table": table,
			},
			Who: err,
		})
	}

	return &User{
		Cassandra: session,
		Table:     table,
	}, nil
}
