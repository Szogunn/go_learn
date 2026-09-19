package maps

const (
	ErrNotFound        = DictionaryErr("could not find the word you were looking for")
	ErrWordExists      = DictionaryErr("cannot add word because it already exists")
	ErrKeyDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", DictionaryErr(ErrNotFound)
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	if _, ok := d[key]; ok {
		return DictionaryErr(ErrWordExists)
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	if _, ok := d[key]; !ok {
		return DictionaryErr(ErrKeyDoesNotExist)
	}

	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return DictionaryErr(ErrKeyDoesNotExist)
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}
