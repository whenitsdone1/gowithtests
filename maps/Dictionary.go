package Dictionary

import "errors"

type Dictionary map[string]string

var alreadyPresentError = errors.New("Word already present in dictionary.")
var notPresentError = errors.New("Word is not yet in dictionary, adding...")

func (d Dictionary) Search(word string) string {
	val, exists := d[word]
	if exists {
		return val
	} else {
		return "Key not found"
	}
}

func (d Dictionary) Add(word, meaning string) error {
	_, exists := d[word]
	if exists {
		return alreadyPresentError
	}
	d[word] = meaning
	return nil
}
func (d Dictionary) Update(word, meaning string) error {
	_, exists := d[word]
	if !exists {
		d.Add(word, meaning)
		return notPresentError
	}
	d[word] = meaning
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, exists := d[word]
	if !exists {
		return notPresentError
	}
	delete(d, word)
	return nil
}
