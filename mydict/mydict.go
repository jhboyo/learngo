package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

// Dictionary errors
var (
	errNotFound = errors.New("not found")
	errCantUpdate = errors.New("can't update non-existing word")
	errWordExists = errors.New("that word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil	
	}
	return "", errNotFound
}


// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word in the dictionary
func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}



// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}