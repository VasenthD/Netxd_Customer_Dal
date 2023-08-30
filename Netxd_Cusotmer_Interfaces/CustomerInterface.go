package interfaces

import models "Netxd_project/Netxd_Customer_Models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.CustomerResponse, error)
}
