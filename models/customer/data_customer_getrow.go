package customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const execSPDataCustomerDetail = `exec [sp_smile_approved_customer_getrow] $1`

func GetRowDataCustomer(ctx context.Context, db *sql.DB,
	customerNo string, log *zap.Logger) ([]*entities.DataApprovedCustomerDetail, error) {

	rows, err := db.QueryContext(
		ctx,
		execSPDataCustomerDetail,
		customerNo,
	)

	if err != nil {
		log.Info("Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_customers := make([]*entities.DataApprovedCustomerDetail, 0)
	for rows.Next() {
		data_customer := &entities.DataApprovedCustomerDetail{}

		if err := rows.Scan(
			&data_customer.CustomerNo,
			&data_customer.CustomerName,
			&data_customer.CustomerType,
			&data_customer.Type,
			&data_customer.EstablishedIn,
			&data_customer.NoOfEmployee,
			&data_customer.AnnualSales,
			&data_customer.PhoneNo,
			&data_customer.Fax,
			&data_customer.Email,
			&data_customer.Website,
			&data_customer.StreetAddress,
			&data_customer.ProvinceId,
			&data_customer.ProvinceName,
			&data_customer.CityId,
			&data_customer.CityName,
			&data_customer.DistrictId,
			&data_customer.DistrictName,
			&data_customer.SubdistrictId,
			&data_customer.SubdistrictName,
			&data_customer.PostalCode,
			&data_customer.TaxStatus,
			&data_customer.TaxMandatory,
			&data_customer.VatRegNo,
			&data_customer.NpwpName,
			&data_customer.VatAddress,
			&data_customer.VatImg,
			&data_customer.AliasName,
			&data_customer.AliasNameFull,
			&data_customer.BillingAddress,
			&data_customer.BillingProvince,
			&data_customer.BillingProvinceName,
			&data_customer.BillingCity,
			&data_customer.BillingCityName,
			&data_customer.BillingDistrict,
			&data_customer.BillingDistrctName,
			&data_customer.BillingSubdistrict,
			&data_customer.BillingSubdistrictName,
			&data_customer.BillingPostal,
			&data_customer.InvoiceMethod,
			&data_customer.PaymentTermsId,
			&data_customer.InvoiceDocument,
			&data_customer.ReturnStatus,
			&data_customer.InvoiceSchedule,
			&data_customer.Status,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_customers = append(data_customers, data_customer)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_customers, nil
}
