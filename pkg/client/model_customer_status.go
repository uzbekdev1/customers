/*
 * Customers API
 *
 * Customers focuses on solving authentic identification of humans who are legally able to hold and transfer currency within the US. Primarily this project solves [Know Your Customer](https://en.wikipedia.org/wiki/Know_your_customer) (KYC), [Customer Identification Program](https://en.wikipedia.org/wiki/Customer_Identification_Program) (CIP), [Office of Foreign Asset Control](https://www.treasury.gov/about/organizational-structure/offices/Pages/Office-of-Foreign-Assets-Control.aspx) (OFAC) checks and verification workflows to comply with US federal law and ensure authentic transfers. Customers has an objective to be a service for detailed due diligence on individuals and companies for Financial Institutions and services in a modernized and extensible way.  Customer phone numbers and addresses are stored and partially used in KYC/OFAC validation. Arbitrary key/value pairs can be stored for a Customer. Documents and Disclaimers, and their acknowledgement are also stored under a Customer as they're accepted. Bank Accounts, which can be validated with micro-deposits currently, are stored under each Customer.  ![](https://raw.githubusercontent.com/adamdecaf/customers/create-accounts/docs/images/customer.png)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// CustomerStatus State of the customer
type CustomerStatus string

// List of CustomerStatus
const (
	CUSTOMERSTATUS_DECEASED     CustomerStatus = "Deceased"
	CUSTOMERSTATUS_REJECTED     CustomerStatus = "Rejected"
	CUSTOMERSTATUS_UNKNOWN      CustomerStatus = "Unknown"
	CUSTOMERSTATUS_RECEIVE_ONLY CustomerStatus = "ReceiveOnly"
	CUSTOMERSTATUS_VERIFIED     CustomerStatus = "Verified"
	CUSTOMERSTATUS_FROZEN       CustomerStatus = "Frozen"
)
