package Services

import (
	"context"
	"time"

	interfaces "github.com/VasenthD/Netxd_Customer_Dal/Netxd_Cusotmer_Interfaces"
	models "github.com/VasenthD/Netxd_Customer_Dal/Netxd_Customer_Models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomer(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (c *CustomerService) CreateCustomer(user *models.Customer) (*models.CustomerResponse, error) {

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.IsActive = true

	_, err := c.CustomerCollection.InsertOne(c.ctx, &user)
	if err != nil {
		return nil, err
	}
	response := &models.CustomerResponse{
		CustomerId: user.Customer_Id,
		CreatedAt:  user.CreatedAt.UTC(),
	}

	return response, nil
}
