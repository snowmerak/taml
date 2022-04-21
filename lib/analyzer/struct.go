package analyzer

import (
	"errors"
	"reflect"
)

type Member struct {
	Name      string
	Type      string
	SubMmbers []Member
}

func MembersOf(t reflect.Type) ([]Member, error) {
	typ := t
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("not a struct")
	}
	members := make([]Member, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Type.Kind() == reflect.Struct {
			subMembers, err := MembersOf(f.Type)
			if err != nil {
				return nil, err
			}
			members = append(members, Member{
				Name:      f.Name,
				Type:      f.Type.String(),
				SubMmbers: subMembers,
			})
		} else {
			members = append(members, Member{
				Name: f.Name,
				Type: f.Type.String(),
			})
		}
	}
	return members, nil
}
