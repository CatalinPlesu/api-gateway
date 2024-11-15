package application

import (
	"os"
	"strconv"
)


type Config struct {
	RedisAddress           string
	ServerPort             uint16
	UserServiceAddress     string
	ChannelServiceAddress  string
	MessageServiceAddress  string
	LiveTypingServiceAddress string
}

func LoadConfig() Config {
	cfg := Config{
		ServerPort:             2020,
		RedisAddress:           "localhost:6379",
		UserServiceAddress:     "http://localhost:3000",  
		ChannelServiceAddress:  "http://localhost:3001",  
		MessageServiceAddress:  "http://localhost:3002",  
		LiveTypingServiceAddress: "http://localhost:3003",
	}

	// Load Redis address from environment variable
	if redisAddr, exists := os.LookupEnv("REDIS_ADDR"); exists {
		cfg.RedisAddress = redisAddr
	}

	// Load server port from environment variable
	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	// Load service addresses from environment variables
	if userServiceAddr, exists := os.LookupEnv("USER_SERVICE_ADDR"); exists {
		cfg.UserServiceAddress = userServiceAddr
	}
	if channelServiceAddr, exists := os.LookupEnv("CHANNEL_SERVICE_ADDR"); exists {
		cfg.ChannelServiceAddress = channelServiceAddr
	}
	if messageServiceAddr, exists := os.LookupEnv("MESSAGE_SERVICE_ADDR"); exists {
		cfg.MessageServiceAddress = messageServiceAddr
	}
	if liveTypingServiceAddr, exists := os.LookupEnv("LIVE_TYPING_SERVICE_ADDR"); exists {
		cfg.LiveTypingServiceAddress = liveTypingServiceAddr
	}

	return cfg
}
