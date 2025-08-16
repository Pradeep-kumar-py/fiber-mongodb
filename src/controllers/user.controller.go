package controllers

import (
    "github.com/gofiber/fiber/v2"
)

// RegisterUser handles user registration
func RegisterUser(c *fiber.Ctx) error {
    var data struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    
    if err := c.BodyParser(&data); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }
    
    // Your registration logic here
    return c.JSON(fiber.Map{
        "message": "User registered successfully",
        "user": fiber.Map{
            "name":  data.Name,
            "email": data.Email,
        },
    })
}

// LoginUser handles user login
func LoginUser(c *fiber.Ctx) error {
    // Your login logic here
    return c.JSON(fiber.Map{
        "message": "User logged in",
    })
}

// LogoutUser handles user logout
func LogoutUser(c *fiber.Ctx) error {
    // Your logout logic here
    return c.JSON(fiber.Map{
        "message": "User logged out",
    })
}

// GetAllUsers gets all users
func GetAllUsers(c *fiber.Ctx) error {
    // Your get all users logic here
    return c.JSON(fiber.Map{
        "message": "All users",
        "users":   []string{"user1", "user2"},
    })
}