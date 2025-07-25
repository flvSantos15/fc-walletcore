package entity

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Client struct {
	ID string
	Name string
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		ID: uuid.NewV4().String(),
		Name: name,
		Email: email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := client.Validate()

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()

	err := c.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Email == "" {
		return errors.New("email is required")
	}

	return nil
}