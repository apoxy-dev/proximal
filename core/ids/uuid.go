package ids

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/martinlindhe/base36"
)

func uuidToExternalID(id uuid.UUID) string {
	a := [16]byte(id)
	return strings.ToLower(base36.EncodeBytes(a[:]))
}

func externalIDToUUID(id string) (uuid.UUID, error) {
	dec := base36.DecodeToBytes(strings.ToUpper(id))
	return uuid.FromBytes(dec)
}

type IDPrefix string

const (
	Key IDPrefix = "k"
)

type PrefixedID struct {
	Prefix IDPrefix
	ID     uuid.UUID
}

func ParsePrefixedID(id string) (PrefixedID, error) {
	prefix := IDPrefix(id[0])
	uid, err := externalIDToUUID(id[1:])
	return PrefixedID{
		Prefix: prefix,
		ID:     uid,
	}, err
}

func (pid *PrefixedID) String() string {
	return fmt.Sprintf("%s%s", pid.Prefix, uuidToExternalID(pid.ID))
}

// NewKey returns a new PrefixedID with the Key prefix and a new uuid.
func NewKey() PrefixedID {
	return PrefixedID{
		Prefix: Key,
		ID:     uuid.New(),
	}
}
