package tweet

import (
	"errors"
	"github.com/Edgarmontenegro123/X-Twitter/internal/user"
	"testing"
)

// MockDB es una implementación de TweetDB para pruebas
type MockDBService struct {
	tweets    []Tweet
	users     map[int]user.User
	followers map[int][]user.User
	db        user.UserDB
}

// SaveUser guarda un usuario en la simulación de la base de datos
func (m *MockDBService) SaveUser(u user.User) {
	m.users[u.ID] = u
}

func (m *MockDBService) SaveTweet(tweet Tweet) {
	m.tweets = append(m.tweets, tweet)
}

func (m *MockDBService) GetTweetsByUserID(userID int) []Tweet {
	return m.tweets
}

func (m *MockDBService) GetUserByID(userID int) user.User {
	return m.users[userID]
}

func (m *MockDBService) GetFollowers(userID int) ([]user.User, error) {
	followers, ok := m.followers[userID]
	if !ok {
		return nil, errors.New("usuario no encontrado")
	}
	return followers, nil
}

func (m *MockDBService) SetDB(db user.UserDB) {
	m.db = db
}

// Nuevo constructor para MockDBService que inicializa los mapas
func NewMockDBService(db user.UserDB) *MockDBService {
	return &MockDBService{
		users:     make(map[int]user.User),
		tweets:    make([]Tweet, 0),
		followers: make(map[int][]user.User),
		db:        db,
	}
}

func TestPublishTweet(t *testing.T) {
	// Configuración de prueba
	mockDBService := &MockDBService{}
	tweetService := NewTweetService(mockDBService)

	// Caso de prueba
	err := tweetService.PublishTweet(1, "Este es un tweet de prueba")

	// Verificación
	if err != nil {
		t.Errorf("Se esperaba un error nulo, pero se obtuvo: %v", err)
	}

	// Verifica si el tweet se guarda correctamente
	if len(mockDBService.tweets) != 1 {
		t.Errorf("Se esperaba que se guarde un tweet, pero se guardaron %d", len(mockDBService.tweets))
	}
}

/*func TestGetTweets(t *testing.T) {
	mockDBService := NewMockDBService()
	tweetService := NewTweetService(mockDBService)

	// Crear instancias de usuarios
	user1 := user.User{ID: 1, Username: "user1"}
	user2 := user.User{ID: 2, Username: "user2"}
	user3 := user.User{ID: 3, Username: "user3"}

	// Guardar usuarios en la base de datos
	mockDBService.SaveUser(user1)
	mockDBService.SaveUser(user2)
	mockDBService.SaveUser(user3)

	// Guardar tweets para los usuarios
	mockDBService.SaveTweet(Tweet{ID: 1, UserId: 2, Content: "Tweet1"})
	mockDBService.SaveTweet(Tweet{ID: 2, UserId: 1, Content: "Tweet2"})
	mockDBService.SaveTweet(Tweet{ID: 3, UserId: 3, Content: "Tweet3"})

	mockDBService.Follow(1, 2)
	mockDBService.Follow(1, 3)

	// Crear una instancia del servicio de tweets con la base de datos simulada
	tweetService = NewTweetService(mockDBService)

	// Obtener los tweets para el usuario 1 (sigue a user2 y user3)
	timelineTweets, err := tweetService.GetTweets(user1.ID)
	if err != nil {
		t.Fatalf("Error obteniendo tweets: %v", err)
	}

	// Verificar que se obtuvieron los tweets esperados
	expectedTweets := []TimeLineTweet{
		{Tweet: Tweet{ID: 1, UserId: 1, Content: "Tweet1"}, User: user1},
		{Tweet: Tweet{ID: 2, UserId: 2, Content: "Tweet2"}, User: user2},
		{Tweet: Tweet{ID: 3, UserId: 3, Content: "Tweet3"}, User: user3},
	}

	if !reflect.DeepEqual(timelineTweets, expectedTweets) {
		t.Errorf("Resultados esperados no coinciden. Esperado: %+v, Obtenido: %+v", expectedTweets, timelineTweets)
	}
}*/

/*func TestGetTweets(t *testing.T) {
	// Configuración de prueba
	mockDBService := NewMockDBService(nil)
	tweetService := NewTweetService(mockDBService)

	// Crear usuarios de prueba
	user1 := user.User{ID: 1}
	user2 := user.User{ID: 2}

	// Guardar usuarios en la base de datos simulada
	mockDBService.SetDB(mockDBService)
	mockDBService.SaveUser(user1)
	mockDBService.SaveUser(user2)

	// Publicar tweets para ambos usuarios
	tweetService.PublishTweet(user1.ID, "Tweet de usuario 1")
	tweetService.PublishTweet(user2.ID, "Tweet de usuario 2")

	fmt.Println("Usuarios en la base de datos:", mockDBService.users)

	// Obtener tweets para el usuario 1
	tweets, err := tweetService.GetTweets(user1.ID)

	// Verificación
	if err != nil {
		t.Errorf("Se esperaba un error nulo, pero se obtuvo: %v", err)
	}

	// Verificar la cantidad de tweets
	expectedTweets := []TimeLineTweet{
		{Tweet: Tweet{ID: 1, UserId: user1.ID, Content: "Tweet de usuario 1"}, User: user1},
	}
	fmt.Println("Tweets obtenidos:", tweets)
	fmt.Println("Tweets esperados:", expectedTweets)

	if !reflect.DeepEqual(tweets, expectedTweets) {
		t.Errorf("Los tweets obtenidos no coinciden con los esperados.\nObtenidos: %+v\nEsperados: %+v", tweets, expectedTweets)
	}
}*/
