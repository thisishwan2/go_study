package dict

import "errors"

// 딕셔너리 선언
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")
var errCantUpdate = errors.New("Cant update word")

// 일반 응답과 에러 응답을 반환으로 선언
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	if err == errNotFound {
		d[key] = value
		return nil
	}
	return errWordExists
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	if err == nil {
		d[key] = value
		return nil
	}
	return errCantUpdate
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
