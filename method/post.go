package method

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"smile-service/entities"
	"smile-service/models/cart"
	"smile-service/models/customer"
	"smile-service/models/login"
	"smile-service/models/need_approve"
	"smile-service/models/new_customer"
	"smile-service/models/smd"
	"smile-service/models/store_visit"
	"smile-service/models/transaction"
	usertoken "smile-service/models/user_token"

	"smile-service/models/notification"

	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
)

func POST(ctx context.Context, db *sql.DB, event events.APIGatewayProxyRequest, log *zap.Logger) *events.APIGatewayProxyResponse {
	var response *events.APIGatewayProxyResponse

	if event.Path == "/login" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataLogin := &entities.Login{}
		err := json.Unmarshal([]byte(event.Body), dataLogin)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		resultLogin, err := login.Login(ctx, db, dataLogin, log)
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
			Result:  resultLogin,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/transaction" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataTransaction := &entities.DataTransaction{}
		err := json.Unmarshal([]byte(event.Body), dataTransaction)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		qty_cart, err := transaction.InsertTransaction(ctx, db, dataTransaction, log)
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
			Result:  qty_cart,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/transaction/promo" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataTransactionPromo := &entities.DataTransactionPromo{}
		err := json.Unmarshal([]byte(event.Body), dataTransactionPromo)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = transaction.InsertTransactionPromo(ctx, db, dataTransactionPromo, log)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/transaction/update" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataTransaction := &entities.UpdateQtyCart{}
		err := json.Unmarshal([]byte(event.Body), dataTransaction)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = transaction.UpdateQtyTransaction(ctx, db, dataTransaction, log)
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
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataChecklistFreezer := &entities.DataSMDFreezerInsert{}
		err := json.Unmarshal([]byte(event.Body), dataChecklistFreezer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = smd.InsertChecklistFreezer(ctx, db, dataChecklistFreezer)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/smd/freezer/update" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataChecklistFreezer := &entities.DataSMDFreezerUpdate{}
		err := json.Unmarshal([]byte(event.Body), dataChecklistFreezer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = smd.UpdateChecklistFreezer(ctx, db, dataChecklistFreezer)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/startendday" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataStartEndDay := &entities.DataStartEndDay{}
		err := json.Unmarshal([]byte(event.Body), dataStartEndDay)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = store_visit.InsertStartEndDay(ctx, db, dataStartEndDay)

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
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCartProcess := &entities.CartProcess{}
		err := json.Unmarshal([]byte(event.Body), dataCartProcess)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		trans_no, err := cart.CartProcess(ctx, db, dataCartProcess, log)
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
			Result:  trans_no,
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
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCartProcess := &entities.CartProcessEdit{}
		err := json.Unmarshal([]byte(event.Body), dataCartProcess)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		trans_no, err := cart.CartProcessEdit(ctx, db, dataCartProcess, log)
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
			Result:  trans_no,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_request_insert" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCustomerRequest := &entities.InfoNewCustomerRequest{}
		err := json.Unmarshal([]byte(event.Body), dataCustomerRequest)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		customer_request_no, err := new_customer.DataNewCustomerRequestInsert(ctx, db, dataCustomerRequest, log)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  customer_request_no,
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
			Result:  customer_request_no,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve" {

		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		data_approve := &entities.DataNeedApprove{}
		err := json.Unmarshal([]byte(event.Body), data_approve)

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

		result, err := need_approve.GetDataNeedApproveProcess(
			ctx,
			db,
			log,
			data_approve,
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
	} else if event.Path == "/need_approve/reject" {

		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		data_reject := &entities.DataNeedApproveReject{}
		err := json.Unmarshal([]byte(event.Body), data_reject)

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

		err = need_approve.GetDataNeedApproveReject(
			ctx,
			db,
			log,
			data_reject,
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/checkinout" {

		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		data_check_inout := &entities.DataCheckInOut{}
		err := json.Unmarshal([]byte(event.Body), data_check_inout)

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

		rows, err := store_visit.InsertCheckInOut(
			ctx,
			db,
			data_check_inout,
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
			Result:  rows,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/insert_outlet" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCustomerOutlet := &entities.InfoNewCustomerOutlet{}
		err := json.Unmarshal([]byte(event.Body), dataCustomerOutlet)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		outlet_id, err := new_customer.DataNewCustomerOutletInsert(ctx, db, dataCustomerOutlet, log)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  outlet_id,
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
			Result:  outlet_id,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/insert_pic" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCustomerPic := &entities.InfoNewCustomerPic{}
		err := json.Unmarshal([]byte(event.Body), dataCustomerPic)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		pic_id, err := new_customer.DataNewCustomerPicInsert(ctx, db, dataCustomerPic, log)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  pic_id,
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
			Result:  pic_id,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/add_freezer" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCustomerFreezer := &entities.InfoNewCustomerFreezer{}
		err := json.Unmarshal([]byte(event.Body), dataCustomerFreezer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		freezer_id, err := new_customer.DataNewFreezerInsert(ctx, db, dataCustomerFreezer, log)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  freezer_id,
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
			Result:  freezer_id,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_request_update" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataNewCustomer := &entities.InfoNewCustomerRequestUpdate{}
		err := json.Unmarshal([]byte(event.Body), dataNewCustomer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = new_customer.CustomerReqestUpdate(ctx, db, dataNewCustomer)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_request_delete" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataNewCustomer := &entities.CustomerRequestNo{}
		err := json.Unmarshal([]byte(event.Body), dataNewCustomer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = new_customer.DeleteCustomerRequest(ctx, db, dataNewCustomer, log)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_outlet_delete" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataOutlet := &entities.OutletId{}
		err := json.Unmarshal([]byte(event.Body), dataOutlet)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = new_customer.DeleteCustomerOutlet(ctx, db, dataOutlet)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_outlet_update" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataOutlet := &entities.InfoNewCustomerOutletUpdate{}
		err := json.Unmarshal([]byte(event.Body), dataOutlet)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = new_customer.DataNewCustomerOutletUpdate(ctx, db, dataOutlet)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/outlet_freezer_delete" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataFreezer := &entities.FreezerId{}
		err := json.Unmarshal([]byte(event.Body), dataFreezer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = new_customer.DeleteOutletFreezer(ctx, db, dataFreezer)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_pic_update" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataPic := &entities.InfoNewCustomerPicUpdate{}
		err := json.Unmarshal([]byte(event.Body), dataPic)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = new_customer.UpdateCustomerPic(ctx, db, dataPic)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_pic_delete" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataPic := &entities.PicId{}
		err := json.Unmarshal([]byte(event.Body), dataPic)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = new_customer.DeleteCustomerPic(ctx, db, dataPic)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/checkfreezer/insert" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataChecklistFreezer := &entities.DataCheckFreezerInsert{}
		err := json.Unmarshal([]byte(event.Body), dataChecklistFreezer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		preSignURLS, err := store_visit.InsertCheckFreezer2(ctx, db, dataChecklistFreezer)

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
			Result:  preSignURLS,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/add_outlet_freezer_update" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataFreezer := &entities.InfoNewCustomerFreezerUpdate{}
		err := json.Unmarshal([]byte(event.Body), dataFreezer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = new_customer.DataOutletFreezerUpdate(ctx, db, dataFreezer)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/approve_customer_request_insert" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataApproveCustomer := &entities.DataNeedApprove{}
		err := json.Unmarshal([]byte(event.Body), dataApproveCustomer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		customer_no, err := need_approve.ApproveCustomerRequest(ctx, db, dataApproveCustomer, log)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  customer_no,
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
			Result:  customer_no,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/reject_customer_request_insert" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataApproveCustomer := &entities.DataNeedApproveReject{}
		err := json.Unmarshal([]byte(event.Body), dataApproveCustomer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = need_approve.RejectCustomerRequest(ctx, db, log, dataApproveCustomer)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/profile/change_password" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		changePassword := &entities.ChangePassword{}
		err := json.Unmarshal([]byte(event.Body), changePassword)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		resultChangePassword, err := login.ChangePassword(ctx, db, changePassword, log)
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
			Result:  resultChangePassword,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/checkstock/insert" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCheckStock := &entities.DataCheckStockInsert{}
		err := json.Unmarshal([]byte(event.Body), dataCheckStock)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		preSignURLS, err := store_visit.InsertCheckStock2(ctx, db, dataCheckStock)

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
			Result:  preSignURLS,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/finding/insert" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataFinding := &entities.DataFindingInsert{}
		err := json.Unmarshal([]byte(event.Body), dataFinding)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = store_visit.InsertFinding(ctx, db, dataFinding)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/approve_customer_address_insert" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCustomerAddress := &entities.DataNeedApprove{}
		err := json.Unmarshal([]byte(event.Body), dataCustomerAddress)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = need_approve.ApproveCustomerAddress(ctx, db, dataCustomerAddress, log)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/token/login" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataUserToken := &entities.UserToken{}
		err := json.Unmarshal([]byte(event.Body), dataUserToken)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = usertoken.UserTokenInsert(ctx, db, dataUserToken, log)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/token/logout" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataUserToken := &entities.UserTokenLogout{}
		err := json.Unmarshal([]byte(event.Body), dataUserToken)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = usertoken.UserTokenLogout(ctx, db, dataUserToken, log)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/notification" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataNotificationSend := &entities.NotificationSend{}
		err := json.Unmarshal([]byte(event.Body), dataNotificationSend)

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

		err = notification.SendNotificationSO(ctx, db, dataNotificationSend, log)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/notification/delete" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataUserNotificationDelete := &entities.NotificationDelete{}
		err := json.Unmarshal([]byte(event.Body), dataUserNotificationDelete)

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

		err = notification.UserNotificationDelete(ctx, db, dataUserNotificationDelete, log)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/notification/read" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataUserNotificationRead := &entities.NotificationRead{}
		err := json.Unmarshal([]byte(event.Body), dataUserNotificationRead)
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

		err = notification.UserNotificationRead(ctx, db, dataUserNotificationRead, log)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/add_freezer" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCustomerFreezer := &entities.InfoCustomerFreezer{}
		err := json.Unmarshal([]byte(event.Body), dataCustomerFreezer)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		freezer_id, err := customer.DataCustomerFreezerInsert(ctx, db, dataCustomerFreezer, log)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  freezer_id,
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
			Result:  freezer_id,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/add_document" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		DataNewCustomerDocument := &entities.InfoNewCustomerDocument{}
		err := json.Unmarshal([]byte(event.Body), DataNewCustomerDocument)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		doc_id, err := new_customer.DataNewCustomerDocumentInsert(ctx, db, DataNewCustomerDocument, log)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
				Result:  doc_id,
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
			Result:  doc_id,
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/new_customer/new_customer_document_update" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataNewDoc := &entities.InfoNewCustomerDocumentUpdate{}
		err := json.Unmarshal([]byte(event.Body), dataNewDoc)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = new_customer.DataNewCustomerDocumentUpdate(ctx, db, dataNewDoc)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/storevisit/update_store_coordinate" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCoordinate := &entities.DataStoreVisitUpdateCoordinate{}
		err := json.Unmarshal([]byte(event.Body), dataCoordinate)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = store_visit.UpdateStoreCoordinate(ctx, db, dataCoordinate)
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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/approve_document" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataDoc := &entities.InfoCustomerDocument{}
		err := json.Unmarshal([]byte(event.Body), dataDoc)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = need_approve.DataCustomerDocInsert(ctx, db, dataDoc, log)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/customer/customer_asset_insert" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataCustomerAsset := &entities.InfoCustomerAsset{}
		err := json.Unmarshal([]byte(event.Body), dataCustomerAsset)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		err = customer.DataCustomerAssetInsert(ctx, db, dataCustomerAsset, log)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/send_email" {
		// event.Headers = map[string]string{
		// 	"Access-Control-Allow-Origin": "*",
		// 	"Content-Type":                "application/json",
		// }

		dataMail := &entities.DataSendEmail{}
		err := json.Unmarshal([]byte(event.Body), dataMail)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		need_approve.SendEmailTest()

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}

	} else if event.Path == "/need_approve/copy_folder_document" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataDoc := &entities.DataCustomerNumbers{}
		err := json.Unmarshal([]byte(event.Body), dataDoc)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		need_approve.FolderDocumentCopy(dataDoc)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/customer_doc_insert" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataDoc := &entities.InfoCustomerDocument{}
		err := json.Unmarshal([]byte(event.Body), dataDoc)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		need_approve.DataCustomerDocumentInsert(ctx, db, dataDoc, log)

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
		})

		response = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
				"Content-Type":                "application/json",
			},
			Body: string(body),
		}
	} else if event.Path == "/need_approve/send_email_benfarm" {
		event.Headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}

		dataFreezerBen := &entities.DataSendEmailBenfarm{}
		err := json.Unmarshal([]byte(event.Body), dataFreezerBen)

		if err != nil {
			body, _ := json.Marshal(&entities.DefaultResponse{
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			})

			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
					"Content-Type":                "application/json",
				},
				Body: string(body),
			}
		}

		need_approve.SendEmailHashNode(dataFreezerBen)

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
