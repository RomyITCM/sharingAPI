package method

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"smile-service/entities"
	"smile-service/models/ar_overdue"
	"smile-service/models/area"
	"smile-service/models/cart"
	"smile-service/models/customer"
	transaction_hold "smile-service/models/drafts"
	"smile-service/models/history_transaction"
	"smile-service/models/need_approve"
	"smile-service/models/new_customer"
	"smile-service/models/notification"
	"smile-service/models/pre_signed_url"
	"smile-service/models/price_zone"
	"smile-service/models/product"
	"smile-service/models/site"
	"smile-service/models/smd"
	"smile-service/models/store_visit"
	"smile-service/models/transaction_success"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.uber.org/zap"

	// "github.com/aws/aws-sdk-go-v2/aws"

	// "github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GET(ctx context.Context, db *sql.DB, event events.APIGatewayProxyRequest, log *zap.Logger) *events.APIGatewayProxyResponse {
	var response *events.APIGatewayProxyResponse

	if event.Path == "/customer" {
		data_customers, err := customer.GetDataCustomer(
			ctx,
			db,
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  db,
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_customers,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/bill_to" {
		data_customers_bill_to, err := customer.GetDataCustomerBillTo(
			ctx,
			db,
			event.QueryStringParameters["customer"],
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_customers_bill_to,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/ship_to" {
		data_customers_ship_to, err := customer.GetDataCustomerShipTo(
			ctx,
			db,
			event.QueryStringParameters["customer"],
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_customers_ship_to,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/promo" {
		data_customers_promo_active, err := customer.GetDataCustomerPromoActive(
			ctx,
			db,
			event.QueryStringParameters["customer_no"],
			event.QueryStringParameters["customer_type"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  db,
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_customers_promo_active,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/site" {
		data_sites, err := site.GetDataSite(
			ctx,
			db,
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_sites,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/products" {
		data_products, err := product.GetDataProduct(
			ctx,
			db,
			event.QueryStringParameters["site_id"],
			event.QueryStringParameters["ship_to"],
			event.QueryStringParameters["sales_man"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_products,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/products/edit" {
		data_products, err := product.GetDataProductEdit(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_products,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/cart" {
		cart_products, err := cart.GetDataCartProduct(
			ctx,
			db,
			event.QueryStringParameters["ship_to"],
			event.QueryStringParameters["sales_man"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		cart_delivery, err := cart.GetDataCartDelivery(
			ctx,
			db,
			event.QueryStringParameters["ship_to"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		cart_bills, err := cart.GetDataCartBill(
			ctx,
			db,
			event.QueryStringParameters["ship_to"],
			event.QueryStringParameters["sales_man"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		cart_promo, err := cart.GetDataCartPromo(
			ctx,
			db,
			event.QueryStringParameters["ship_to"],
			event.QueryStringParameters["sales_man"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		result := &entities.DataCart{
			CartProducts:   cart_products,
			CartDeliveries: cart_delivery,
			CartBills:      cart_bills,
			CartPromo:      cart_promo,
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  result,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/cart/payment_term" {
		cart_payment_term, err := cart.GetDataCartPaymentTerm(
			ctx,
			db,
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  cart_payment_term,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/smd" {
		smd_list, err := smd.GetDataSMD(
			ctx,
			db,
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  smd_list,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/smd/freezers" {
		smd_freezer_list, err := smd.GetDataSMDFreezerList(
			ctx,
			db,
			event.QueryStringParameters["customer_no"],
			event.QueryStringParameters["bill_to"],
			event.QueryStringParameters["ship_to"],
			event.QueryStringParameters["created_by"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  smd_freezer_list,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/smd/freezer" {
		smd_freezer, err := smd.GetDataSMDFreezer(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			event.QueryStringParameters["serial_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  smd_freezer,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/transaction/success" {
		data_transaction_success_header, err := transaction_success.GetDataTransactionSuccesHeader(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  db,
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		data_transaction_success_detail, err := transaction_success.GetDataTransactionSuccesDetail(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  db,
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		data_transaction_success_promo, err := transaction_success.GetDataTransactionSuccesPromo(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  db,
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		result := &entities.DataTransactionSuccess{
			DataTransactionSuccessHeader: data_transaction_success_header,
			DataTransactionSuccessDetail: data_transaction_success_detail,
			DataTransactionSuccessPromo:  data_transaction_success_promo,
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  result,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}

		// _ = notification.SendNotificationSO(ctx, db, event.QueryStringParameters["trans_no"], log)

	} else if event.Path == "/transaction/drafts" {
		data_transaction_hold, err := transaction_hold.GetDataTransactionHold(
			ctx,
			db,
			event.QueryStringParameters["salesman"],
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_transaction_hold,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/history_transaction/list" {
		data_history_transaction, err := history_transaction.GetDataHistoryTransaction(
			ctx,
			db,
			event.QueryStringParameters["page"],
			event.QueryStringParameters["time_stamp"],
			event.QueryStringParameters["search"],
			event.QueryStringParameters["user_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_history_transaction,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/history_transaction/detail" {
		history_products, err := history_transaction.GetDataHistoryTransactionProduct(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		history_detail, err := history_transaction.GetDataHistoryTransactionDetail(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		result := &entities.DataHistoryTransactionDetail{
			DataHistoryTransactionProducts: history_products,
			DataHistoryTransactionDelivery: history_detail,
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  result,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_request" {
		new_customer_request, err := new_customer.GetDataNewCustomerRequest(
			ctx,
			db,
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  new_customer_request,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/list" {
		data_need_approve, err := need_approve.GetDataNeedApprove(
			ctx,
			db,
			event.QueryStringParameters["user_id"],
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_need_approve,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/getrows" {
		new_customer_outlet, err := store_visit.GetDataStoreVisit(
			ctx,
			db,
			event.QueryStringParameters["pic"],
			event.QueryStringParameters["visit_date"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  new_customer_outlet,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/startday/getrow" {
		data_start_day, err := store_visit.GetDataStartDay(
			ctx,
			db,
			event.QueryStringParameters["id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_start_day,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/checkinout/getrow" {
		data_check_in_out, err := store_visit.GetDataCheckInOut(
			ctx,
			db,
			event.QueryStringParameters["id"],
			event.QueryStringParameters["type"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_check_in_out,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/storelist" {
		store_list, err := store_visit.GetDataStoreList(
			ctx,
			db,
			event.QueryStringParameters["search"],
			event.QueryStringParameters["dept_code"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  store_list,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/checkfreezer/getrows" {
		freezer_list, err := store_visit.GetDataCheckFreezerList(
			ctx,
			db,
			event.QueryStringParameters["ship_to"],
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  freezer_list,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_outlet" {
		new_customer_outlet, err := new_customer.GetDataNewCustomerOutlet(
			ctx,
			db,
			event.QueryStringParameters["customer_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  new_customer_outlet,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_outlet_detail" {
		data_new_customer_outlet_detail, err := new_customer.GetDataNewCustomerOutletDetail(
			ctx,
			db,
			event.QueryStringParameters["outlet_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_new_customer_outlet_detail,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_pic" {
		new_customer_pic, err := new_customer.GetDataNewCustomerPic(
			ctx,
			db,
			event.QueryStringParameters["customer_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  new_customer_pic,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_preview" { //getrow preview
		new_customer_preview, err := new_customer.GetDataPreviewCustomer(
			ctx,
			db,
			event.QueryStringParameters["customer_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  new_customer_preview,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/outlet_freezers" {
		data_new_customer_freezer, err := new_customer.GetDataNewCustomerFreezer(
			ctx,
			db,
			event.QueryStringParameters["outlet_id"],
			event.QueryStringParameters["freezer_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_new_customer_freezer,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_pic_detail" {
		row_data_customer_pic, err := new_customer.GetRowNewCustomerPic(
			ctx,
			db,
			event.QueryStringParameters["pic_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  row_data_customer_pic,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/approve_customer_getrows" {
		data_need_approve_customer, err := need_approve.GetNeedApproveCustomer(
			ctx,
			db,
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_need_approve_customer,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/approve_outlet_getrows" {
		data_need_approve_outlet, err := need_approve.GetNeedApproveOutlet(
			ctx,
			db,
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_need_approve_outlet,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/checkstock/getrows" {
		data_check_stock_list, err := store_visit.GetDataCheckStockList(
			ctx,
			db,
			event.QueryStringParameters["ship_to"],
			event.QueryStringParameters["search"],
			event.QueryStringParameters["created_by"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_check_stock_list,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/checkstock/competitor/getrows" {
		data_check_stock_list, err := store_visit.GetDataCheckStockCompetitorList(
			ctx,
			db,
			event.QueryStringParameters["ship_to"],
			event.QueryStringParameters["sku_article"],
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_check_stock_list,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/active_customer_getrows" {
		data_customer, err := customer.GetRowsDataCustomer(
			ctx,
			db,
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_customer,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/active_customer_getrow" {
		data_customer_detail, err := customer.GetRowDataCustomer(
			ctx,
			db,
			event.QueryStringParameters["customer_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_customer_detail,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/customer_address_getrows" {
		customerAddress, err := customer.GetRowsCustomerAddress(
			ctx,
			db,
			event.QueryStringParameters["customer_no"],
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  customerAddress,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/customer_address_getrow" {
		customerAddress, err := customer.GetRowCustomerAddress(
			ctx,
			db,
			event.QueryStringParameters["customer_address_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  customerAddress,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/customer_pic_getrows" {
		data_customer_pic, err := customer.GetRowsDataCustomerPic(
			ctx,
			db,
			event.QueryStringParameters["customer_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_customer_pic,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/customer_freezers" {
		data_customer_freezer, err := customer.GetRowsCustomerFreezer(
			ctx,
			db,
			event.QueryStringParameters["customer_address_id"],
			event.QueryStringParameters["freezer_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_customer_freezer,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/customer_documents" {
		data_new_customer_document, err := new_customer.GetRowsNewCustomerDocument(
			ctx,
			db,
			event.QueryStringParameters["customer_request_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_new_customer_document,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/checkfreezer/status" {
		data_check_freezer_status, err := store_visit.GetDataCheckFreezerStatus(
			ctx,
			db,
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_check_freezer_status,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/user_notifications" {
		data_user_notificatios, err := notification.GetDataUserNotification(
			ctx,
			db,
			event.QueryStringParameters["user_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_user_notificatios,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/customer_banks" {
		data_new_customer_banks, err := new_customer.GetRowsCustomerBanks(
			ctx,
			db,
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_new_customer_banks,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/customer_documents_by_type" {
		data_new_customer_document, err := new_customer.GetRowNewCustomerDocumentByType(
			ctx,
			db,
			event.QueryStringParameters["customer_request_no"],
			event.QueryStringParameters["doc_type"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_new_customer_document,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/area/data_province" {
		data_province, err := area.GetRowsDataProvince(
			ctx,
			db,
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_province,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/area/data_city" {
		data_city, err := area.GetRowsDataCity(
			ctx,
			db,
			event.QueryStringParameters["province_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_city,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/area/data_district" {
		data_district, err := area.GetRowsDataDistrict(
			ctx,
			db,
			event.QueryStringParameters["city_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_district,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/area/data_subdistrict" {
		data_subdistrict, err := area.GetRowsDataSubdistrict(
			ctx,
			db,
			event.QueryStringParameters["district_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_subdistrict,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/checkfreezer/getrow" {
		data_new_customer_document, err := store_visit.GetDataCheckFreezerGetrow(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			event.QueryStringParameters["serial_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_new_customer_document,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/checkstock/getrow" {
		data_new_customer_document, err := store_visit.GetDataCheckStockGetrow(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			event.QueryStringParameters["article_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_new_customer_document,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/checkstock/detail/getrow" {
		data_new_customer_document, err := store_visit.GetDataCheckStockDetailGetrow(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			event.QueryStringParameters["article_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_new_customer_document,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}

	} else if event.Path == "/storevisit/menu/validation" {
		data_menu_validation, err := store_visit.GetDataStoreVisitMenuValidation(
			ctx,
			db,
			event.QueryStringParameters["ship_to"],
			event.QueryStringParameters["created_by"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_menu_validation,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/status_user_approve" {
		user_approve, err := need_approve.GetStatusUserApprove(
			ctx,
			db,
			event.QueryStringParameters["user_no"],
			event.QueryStringParameters["customer_request_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  user_approve,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/benfarm_freezer" {
		benFreezer, err := new_customer.GetRowsBenfarmFreezer(
			ctx,
			db,
			event.QueryStringParameters["customer_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  benFreezer,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}

	} else if event.Path == "/need_approve/status_approval" {
		status_approval, err := need_approve.GetDataStatusApproved(
			ctx,
			db,
			event.QueryStringParameters["user_id"],
			event.QueryStringParameters["customer_request_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  status_approval,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/ar_overdue/detail" {
		ar_overdue_header, err := ar_overdue.GetDataCustArOverdueGetrow(
			ctx,
			db,
			event.QueryStringParameters["customer_no"],
			event.QueryStringParameters["ship_to"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		ar_overdue_detail, err := ar_overdue.GetDataCustArOverdueGetrows(
			ctx,
			db,
			event.QueryStringParameters["customer_no"],
			event.QueryStringParameters["ship_to"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		result := &entities.DataCustArOverdue{
			AROverdueHeader: ar_overdue_header,
			AROverdueDetail: ar_overdue_detail,
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  result,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/area/master_region" {
		regions, err := area.GetRowsDataRegion(
			ctx,
			db,
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  regions,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/customer_address_bill_to_by_customer_getrows" {
		billTo, err := customer.GetRowsCustomerAddressBillToByCustomer(
			ctx,
			db,
			event.QueryStringParameters["customer_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  billTo,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/price_zone/price_zone_list" {
		priceZone, err := price_zone.GetRowsPriceZone(
			ctx,
			db,
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  priceZone,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/price_zone/article_price" {
		articles, err := price_zone.GetRowsArticlePrice(
			ctx,
			db,
			event.QueryStringParameters["zone_id"],
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  articles,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/area/master_area" {
		areas, err := area.GetRowsArea(
			ctx,
			db,
			event.QueryStringParameters["region_id"],
			event.QueryStringParameters["search"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  areas,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/pic_responsibility" {
		responsibilities, err := new_customer.GetRowsPicResponsibility(
			ctx,
			db,
			event.QueryStringParameters["customer_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  responsibilities,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/s3/getPresignedURL" {

		data_menu_validation, err := getPresignedURL2()

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  data_menu_validation,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/cart/edit" {
		cart_products, err := cart.GetDataCartProductEdit(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		cart_delivery, err := cart.GetDataCartDeliveryEdit(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		cart_bills, err := cart.GetDataCartBillEdit(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		cart_promo, err := cart.GetDataCartPromoEdit(
			ctx,
			db,
			event.QueryStringParameters["trans_no"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		result := &entities.DataCart{
			CartProducts:   cart_products,
			CartDeliveries: cart_delivery,
			CartBills:      cart_bills,
			CartPromo:      cart_promo,
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  result,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/benfarm_freezer_outlet" {
		benFreezer, err := need_approve.GetRowsBenfarmFreezerOutlet(
			ctx,
			db,
			event.QueryStringParameters["outlet_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  benFreezer,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}

	} else if event.Path == "/s3/pre_signed_url" {
		pre_signed_url, err := pre_signed_url.GetDataPreSignedURL(
			ctx,
			event.QueryStringParameters["file_path"],
			event.QueryStringParameters["data"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  pre_signed_url,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/status_approval_outlet" {
		status_approval, err := need_approve.GetDataOutletStatusApproved(
			ctx,
			db,
			event.QueryStringParameters["user_no"],
			event.QueryStringParameters["outlet_id"],
			log,
		)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		body, _ := json.Marshal(&entities.DefaultResponse{
			Status:  http.StatusText(http.StatusOK),
			Message: "Success",
			Result:  status_approval,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	}

	return response

}

////////////////////////////////////////////////////////////////////

func getPresignedURL2() (string, error) {

	// Load env vars
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	// Load the bucket name
	s3Bucket := os.Getenv("S3_BUCKET")
	if s3Bucket == "" {
		log.Fatal("an s3 bucket was unable to be loaded from env vars")
	}

	// Prepare the S3 request so a signature can be generated
	// svc := s3.New(session.NewSession(&aws.Config{
	// 	Region: aws.String("ap-southeast-1"),
	// 	Credentials: credentials.NewStaticCredentials(
	// 		*aws.String("AKIA3WB2X4HJGJTDWPGL"),
	// 		*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
	// 		"", // a token will be created when the session it's used.
	// 	),
	// }))

	// snippet-start:[s3.go.generate_presigned_url.session]
	sess, _ := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				*aws.String("AKIA3WB2X4HJGJTDWPGL"),
				*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
				"", // a token will be created when the session it's used.
			),
		})
	// snippet-end:[s3.go.generate_presigned_url.session]

	// snippet-start:[s3.go.generate_presigned_url.call]
	svc := s3.New(sess)

	r, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String("test-file.png"),
		ACL:    aws.String("public-read"),
	})

	// Create the pre-signed url with an expiry
	url, err := r.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("Failed to generate a pre-signed url: ", err)
		return "", err
	}

	// Display the pre-signed url
	fmt.Println("Pre-signed URL", url)
	return url, nil
}

// /////////////////////////////////////////////////////////////////

// type S3 struct {
// 	client *s3.Client
// 	signer *s3.PresignClient
// }

// type PresignedURLArgs struct {
// 	Bucket string
// 	Key    string
// 	Expiry time.Duration
// }

// func NewS3(cfg aws.Config) S3 {
// 	client := s3.NewFromConfig(cfg)

// 	return S3{
// 		client: client,
// 		signer: s3.NewPresignClient(client),
// 	}
// }

// // PresignedUploadURL creates a presigned request URL that can be used to upload
// // an object in a bucket. The URL is valid for the specified number of seconds.
// func (s S3) PresignedUploadURL(ctx context.Context, args PresignedURLArgs) (string, error) {
// 	input := s3.PutObjectInput{
// 		Bucket: aws.String(args.Bucket),
// 		Key:    aws.String(args.Key),
// 	}

// 	expiry := func(opts *s3.PresignOptions) {
// 		opts.Expires = args.Expiry
// 	}

// 	req, err := s.signer.PresignPutObject(ctx, &input, expiry)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create request: %w", err)
// 	}

// 	return req.URL, nil
// }

// func getPresignedURL() (string, error) {
// 	ctx := context.Background()

// 	awsConfig, err := NewCrossAccountConfigWithRole(ctx, "arn:aws:iam::803282084306:group/s3.admin")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	s3 := NewS3(awsConfig)

// 	// UPLOAD ------------------------------------------------------------------
// 	uploadArgs := PresignedURLArgs{
// 		Bucket: "upload.file",
// 		Key:    "picture.png",
// 		Expiry: time.Minute * 15,
// 	}

// 	uploadURL, err := s3.PresignedUploadURL(ctx, uploadArgs)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return uploadURL, err

// 	// if err := Upload(ctx, uploadURL); err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// }
// func NewCrossAccountConfigWithRole(ctx context.Context, roleARN string) (aws.Config, error) {
// 	cfg, err := config.LoadDefaultConfig(ctx)
// 	if err != nil {
// 		return aws.Config{}, err
// 	}

// 	stsClient := sts.NewFromConfig(cfg)
// 	stsCreds := stscreds.NewAssumeRoleProvider(stsClient, roleARN)

// 	cfg.Credentials = aws.NewCredentialsCache(stsCreds)

// 	return cfg, nil
// }

// func Upload(ctx context.Context, url string) error {
// 	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, strings.NewReader(`{"users":["a,b,c"]}`))
// 	if err != nil {
// 		return fmt.Errorf("failed to create request: %w", err)
// 	}

// 	res, err := http.DefaultTransport.RoundTrip(req)
// 	if err != nil {
// 		return fmt.Errorf("failed to send request: %w", err)
// 	}

// 	if res.StatusCode != http.StatusOK {
// 		return fmt.Errorf("unexpected response: %s", res.Status)
// 	}

// 	return nil
// }

///////////////////////////////////////////////////////////////////////////////////////////////////////

// func newConfig() (string, error) {
// 	// cfg, err := config.LoadDefaultConfig(context.Background(),
// 	// 	config.WithRegion("ap-southeast-1"),
// 	// 	config.WithSharedConfigProfile(profileName))
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// snippet-start:[s3.go.generate_presigned_url.args]
// 	bucket := flag.String("b", "upload.file", "The bucket")
// 	key := flag.String("k", "picture.png", "The object key")
// 	flag.Parse()

// 	if *bucket == "" || *key == "" {
// 		fmt.Println("You must supply a bucket name (-b BUCKET) and object key (-k KEY)")
// 		return "", nil
// 	}
// 	// snippet-end:[s3.go.generate_presigned_url.args]

// 	// snippet-start:[s3.go.generate_presigned_url.session]
// 	sess, err := session.NewSession(
// 		&aws.Config{
// 			Region: aws.String("ap-southeast-1"),
// 			Credentials: credentials.NewStaticCredentials(
// 				*aws.String("AKIA3WB2X4HJGJTDWPGL"),
// 				*aws.String("VVmHtwoOR6xUosj9fi0NW4hRY9i2KVIvMFb7KX+j"),
// 				"", // a token will be created when the session it's used.
// 			),
// 		})
// 	// snippet-end:[s3.go.generate_presigned_url.session]

// 	// snippet-start:[s3.go.generate_presigned_url.call]
// 	svc := s3.New(sess)

// 	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
// 		Bucket: bucket,
// 		Key:    key,
// 	})
// 	urlStr, err := req.Presign(15 * time.Minute)
// 	// snippet-end:[s3.go.generate_presigned_url.call]
// 	if err != nil {
// 		return "", err
// 	}

// 	// return urlStr, nil

// 	// urlStr, err := GetPresignedURL(sess, bucket, key)
// 	// if err != nil {
// 	// 	fmt.Println("Got an error retrieving a presigned URL:")
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }

// 	// snippet-start:[s3.go.generate_presigned_url.print]
// 	fmt.Println("The presigned URL: " + urlStr + " is valid for 15 minutes")
// 	// snippet-end:[s3.go.generate_presigned_url.print]

// 	return urlStr, nil
// }
