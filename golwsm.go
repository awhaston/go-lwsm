package lwsm

import (
	"errors"
	"fmt"
)

var sessions map[string]interface{}

// Define the session struct
type session struct {
	hello string
}

// NewSessionManager initializes the sessions map (this can be expanded later if needed)
func NewSessionManager() {
	// Initialize or reset the global sessions map (if necessary)
	// Currently, it's initialized at the declaration, so no need to do anything here.
	sessions = make(map[string]interface{})
}

// GetSession retrieves a session by its key and asserts its type to T
func GetSession[T any](key string) (T, error) {
	// Retrieve the session from the sessions map
	value, exists := sessions[key]
	if !exists {
		return *new(T), errors.New("session not found")
	}

	// Type assertion: We assert the value to be of type T
	session, ok := value.(T)
	if !ok {
		// If the assertion fails, return an error
		return *new(T), fmt.Errorf("session type mismatch for key: %s", key)
	}

	return session, nil
}

func main() {
	// Initialize the session manager
	NewSessionManager()

	// Example session data
	id := "1234"
	sessionOpts := session{
		hello: "world",
	}

	// Store the session in the global sessions map
	sessions[id] = sessionOpts

	// Retrieve the session using generics
	sesh, err := GetSession[session](id)
	if err != nil {
		fmt.Println("Error retrieving session:", err)
		return
	}

	// Output the session
	fmt.Println("Session retrieved:", sesh)
}
