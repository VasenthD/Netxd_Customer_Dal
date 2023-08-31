package interfaces

import models "github.com/VasenthD/Netxd_Customer_Dal/Netxd_Customer_Models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.CustomerResponse, error)
}
