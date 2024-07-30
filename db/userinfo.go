import (
    "context"
    "log"
    "os"
    "time"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGO_URI")
    clientOptions := options.Client().ApplyURI(mongoURI)

    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }

    // Test connection
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("Error pinging MongoDB: %v", err)
    }

    log.Println("Connected to MongoDB!")
}
