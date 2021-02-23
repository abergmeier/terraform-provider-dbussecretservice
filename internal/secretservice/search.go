package secretservice

import (
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

	coll := s.GetLoginCollection()

	err = s.Unlock(coll.Path())
	if err != nil {
		return nil, err
	}

	ips, err := s.SearchItems(coll, searchAttributes)
	if err != nil {
		return nil, err
	}

	ret := make([]ss.Secret, len(ips))
	for i := range ips {
		sec, err := s.GetSecret(ips[i], sess.Path())
		if err != nil {
			return nil, err
		}

		ret[i] = *sec
	}
	return ret, nil
}
