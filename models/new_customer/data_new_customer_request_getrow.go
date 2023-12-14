package new_customer

import (
	"context"
	"database/sql"
	"smile-service/entities"

	"go.uber.org/zap"
)

const (
	execSPCustomerRequestRow = `exec [sp_smile_new_customer_request_getrow] $1`
)

func GetDataPreviewCustomer(ctx context.Context, db *sql.DB, customerId string, log *zap.Logger) ([]*entities.DataPreviewCustomer, error) {
	rows, err := db.QueryContext(
		ctx,
		execSPCustomerRequestRow,
		customerId,
	)

	if err != nil {
		log.Info(
			"Exec DB",
			zap.Any("Exec DB", err),
		)
		return nil, err
	}
	defer rows.Close()

	data_preview_customers := make([]*entities.DataPreviewCustomer, 0)
	for rows.Next() {
		data_preview_customer := &entities.DataPreviewCustomer{}

		if err := rows.Scan(
			&data_preview_customer.CustomerRequestNo,
			&data_preview_customer.CustomerName,
			&data_preview_customer.CustomerType,
			&data_preview_customer.Type,
			&data_preview_customer.EstablishedIn,
			&data_preview_customer.NoOfEmployee,
			&data_preview_customer.AnnualSales,
			&data_preview_customer.PhoneNo,
			&data_preview_customer.Fax,
			&data_preview_customer.Email,
			&data_preview_customer.Website,
			&data_preview_customer.StreetAddress,
			&data_preview_customer.ProvinceId,
			&data_preview_customer.ProvinceName,
			&data_preview_customer.CityId,
			&data_preview_customer.CityName,
			&data_preview_customer.DistrictId,
			&data_preview_customer.DistrictName,
			&data_preview_customer.SubdistrictId,
			&data_preview_customer.SubdistrictName,
			&data_preview_customer.PostalCode,
			&data_preview_customer.TaxStatus,
			&data_preview_customer.TaxMandatory,
			&data_preview_customer.VatRegNo,
			&data_preview_customer.NpwpName,
			&data_preview_customer.VatAddress,
			&data_preview_customer.VatImg,
			&data_preview_customer.BillingCode,
			&data_preview_customer.BillingName,
			&data_preview_customer.BillingAddress,
			&data_preview_customer.BillingProvince,
			&data_preview_customer.BillingProvinceName,
			&data_preview_customer.BillingCity,
			&data_preview_customer.BillingCityName,
			&data_preview_customer.BillingDistrict,
			&data_preview_customer.BillingDistrictName,
			&data_preview_customer.BillingSubdistrict,
			&data_preview_customer.BillingSubdistrictName,
			&data_preview_customer.BillingPostal,
			&data_preview_customer.BillingMethod,
			&data_preview_customer.PaymentTermsId,
			&data_preview_customer.BillingDocuments,
			&data_preview_customer.ReturAvailable,
			&data_preview_customer.FactureSchedule,
			&data_preview_customer.Status,
			&data_preview_customer.ApprovalStatus,
		); err != nil {
			log.Info("scan rows",
				zap.Any("rows", err),
			)
			return nil, err
		}

		data_preview_customers = append(data_preview_customers, data_preview_customer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data_preview_customers, nil
}
