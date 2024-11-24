package utils

import (
	"CoreBaseGo/internal/interfaces/rest"
	messages "CoreBaseGo/internal/interfaces/rest/Messages"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ExtractRealIP extracts the real IP address from the request context
func ExtractRealIP(c *gin.Context) (string, error) {
	// Check for trusted headers in the preferred order
	trustedHeaders := []string{"X-Real-IP", "Cf-Connecting-Ip"}
	for _, header := range trustedHeaders {
		realIP := c.GetHeader(header)
		if realIP != "" {
			return realIP, nil
		}
	}

	// Fallback to X-Forwarded-For (cautiously)
	ipAddress := c.GetHeader("X-Forwarded-For")
	if ipAddress != "" {
		// Extract the first IP from comma-separated values (if present)
		parts := strings.Split(ipAddress, ",")
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0]), nil
		}
	}

	// Use the direct remote IP address if no headers are set
	if c.ClientIP() != "" {
		return c.ClientIP(), nil
	}

	// Return an error if no reliable IP found
	return "", fmt.Errorf("unable to extract real IP address")
}

// GenerateRandomOtoCode generate Otp code
func GenerateRandomOtoCode() string { return strconv.Itoa(1000 + rand.Intn(9000)) }

const (
	ipKeyPrefix = "ip_flood:"
)

// CheckFlood attempts to detect and potentially block flood attempts based on IP and phone number.
func CheckFlood(ip string, ipKeyPrefix string) (bool, error) {
	floodLimit := viper.GetInt64("FLOOD_LIMIT")
	floodTime := time.Duration(viper.GetInt64("FLOOD_TIME")) * time.Minute
	blockDuration := time.Duration(viper.GetInt64("FLOOD_BLOCK_DURATION")) * time.Minute

	client := GetRedisClient()
	ctx := context.Background()
	ipKey := fmt.Sprintf("%s%s", ipKeyPrefix, ip)

	// Check IP flood attempts
	val, err := client.Incr(ctx, ipKey).Result()
	if err != nil {
		return false, fmt.Errorf("error checking IP flood count: %w", err)
	}
	// Set expiration for the key after increment
	_, err = client.Expire(ctx, ipKey, floodTime).Result()
	if err != nil {
		return false, fmt.Errorf("error setting expiration for IP flood key: %w", err)
	}

	// Check if IP limit is exceeded
	if val > floodLimit {
		// Block IP for a period
		_, err = client.Set(ctx, ipKey, val, blockDuration).Result()
		if err != nil {
			return false, fmt.Errorf("error blocking IP: %w", err)
		}
		return true, fmt.Errorf("IP flood limit reached. Access blocked for %v", blockDuration)
	}

	// Not blocked
	return false, nil
}

func FloodControl(c *gin.Context, ipKeyPrefix string) bool {
	realIP, ExtractRealIPErr := ExtractRealIP(c)
	if ExtractRealIPErr != nil {
		log.Printf("Error checking flood: %v", ExtractRealIPErr.Error())
		rest.JSONOutput(c, http.StatusInternalServerError, nil, messages.InternalServerError, "Internal server error")
		return true
	}

	isBlocked, FloodErr := CheckFlood(realIP, ipKeyPrefix)
	if FloodErr != nil {
		rest.JSONOutput(c, http.StatusLocked, nil, messages.Locked, "Error checking flood:"+FloodErr.Error())
		return true
	}

	if isBlocked {
		rest.JSONOutput(c, http.StatusTooManyRequests, nil, messages.TooManyRequests, "Access blocked due to flood attempts")
		return true
	}
	return false
}

func HasPassword(PasswordHash string) bool {
	return PasswordHash != ""
}

func HashPassword(password string, key string) (string, error) {
	// Use a secure hashing algorithm like bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+key), bcrypt.DefaultCost)
	if err != nil {
		// Log the error with relevant details
		log.Printf("Error hashing password: %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(hashedPassword string, password string, key string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+key))
	if err != nil {
		// Handle potential errors during comparison
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			// Log a specific message for password mismatch
			log.Printf("Incorrect password entered")
			return false, nil
		}
		log.Printf("Error comparing password: %v", err)

		return false, err
	}
	return true, nil
}

// ClaimsJwtData extracts JWT claims from the context.
func ClaimsJwtData(c *gin.Context) (*Claims, bool) {
	userData, exists := c.Get("user")

	if !exists {
		rest.JSONOutput(c, http.StatusUnauthorized, nil, messages.NoUser, "User data not found in context")
		return nil, true
	}

	// Cast userData to *helper.Claims
	claims, ok := userData.(*Claims)
	if !ok {
		rest.JSONOutput(c, http.StatusUnauthorized, nil, messages.InvalidUserData, "Invalid user data")
		return nil, true
	}

	return claims, false
}
