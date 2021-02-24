package secretservice

import (
	"log"

	"github.com/godbus/dbus/v5"
	ss "github.com/zalando/go-keyring/secret_service"
)

func SearchLogin(searchAttributes map[string]string) ([]ss.Secret, error) {

	bus, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	defer bus.Close()

	s, err := ss.NewSecretService()
	if err != nil {
		return nil, err
	}

	sess, err := s.OpenSession()
	if err != nil {
		return nil, err
	}
	log.Println("[TRACE] Session opened")

	coll := s.GetLoginCollection()

	err = s.Unlock(coll.Path())
	if err != nil {
		return nil, err
	}
	log.Printf("[TRACE] Unlocked %s\n", coll.Path())

	ips, err := s.SearchItems(coll, searchAttributes)
	if err != nil {
		return nil, err
	}
	log.Printf("[TRACE] Found %d Items\n", len(ips))

	ret := make([]ss.Secret, len(ips))
	for i := range ips {
		sec, err := s.GetSecret(ips[i], sess.Path())
		if err != nil {
			return nil, err
		}
		log.Printf("[TRACE] Read Secret %s\n", ips[i])

		ret[i] = *sec
	}
	return ret, nil
}
