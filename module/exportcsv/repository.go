package exportcsv

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// dumy(
type Transaction struct {
	Id         int
	Oda_number string
	Status     string
	Price      float64
	Total_data int
	Created_at string
}

func (r Repository) GetAllTransaction() ([]Transaction, error) {
	//dumy(
	data := []Transaction{
		Transaction{Id: 1,
			Oda_number: "123412341234",
			Status:     "success",
			Price:      17000,
			Total_data: 12,
			Created_at: "2023-06-09 07:51:14",
		},
		
		{Id: 2,
			Oda_number: "2023-06-09 07:51:00",
			Status:     "pending",
			Price:      14000,
			Total_data: 12,
			Created_at: "2023-06-09 07:50:14",
		},
		Transaction{Id: 3,
			Oda_number: "123412341234",
			Status:     "success",
			Price:      16000,
			Total_data: 12,
			Created_at: "2023-06-09 06:51:14",
		},
		Transaction{Id: 4,
			Oda_number: "123412341234",
			Status:     "pending",
			Price:      19000,
			Total_data: 12,
			Created_at: "2023-06-09 05:51:14",
		},
	}
	//)
	return data, nil
}
func (r Repository) GetTransactionByStatus(req *ExportCSVRequest) ([]Transaction, error) {
	//dumy(
	dataSucces := []Transaction{
		Transaction{Id: 1,
			Oda_number: "123412341234",
			Status:     "success",
			Price:      17000,
			Total_data: 12,
			Created_at: "2023-06-09 07:51:14",
		},
		Transaction{Id: 2,
			Oda_number: "2023-06-09 07:51:00",
			Status:     "success",
			Price:      14000,
			Total_data: 12,
			Created_at: "2023-06-09 07:50:14",
		},
		Transaction{Id: 3,
			Oda_number: "123412341234",
			Status:     "success",
			Price:      16000,
			Total_data: 12,
			Created_at: "2023-06-09 06:51:14",
		},
		Transaction{Id: 4,
			Oda_number: "123412341234",
			Status:     "success",
			Price:      19000,
			Total_data: 12,
			Created_at: "2023-06-09 05:51:14",
		},
	}

	dataPending := []Transaction{
		Transaction{Id: 1,
			Oda_number: "123412341234",
			Status:     "pending",
			Price:      17000,
			Total_data: 12,
			Created_at: "2023-06-09 07:51:14",
		},
		Transaction{Id: 2,
			Oda_number: "2023-06-09 07:51:00",
			Status:     "pending",
			Price:      14000,
			Total_data: 12,
			Created_at: "2023-06-09 07:50:14",
		},
		Transaction{Id: 3,
			Oda_number: "123412341234",
			Status:     "pending",
			Price:      16000,
			Total_data: 12,
			Created_at: "2023-06-09 06:51:14",
		},
		Transaction{Id: 4,
			Oda_number: "123412341234",
			Status:     "pending",
			Price:      19000,
			Total_data: 12,
			Created_at: "2023-06-09 05:51:14",
		},
	}
	dataReject := []Transaction{
		Transaction{Id: 1,
			Oda_number: "123412341234",
			Status:     "rejected",
			Price:      17000,
			Total_data: 12,
			Created_at: "2023-06-09 07:51:14",
		},
		Transaction{Id: 2,
			Oda_number: "2023-06-09 07:51:00",
			Status:     "rejected",
			Price:      14000,
			Total_data: 12,
			Created_at: "2023-06-09 07:50:14",
		},
		Transaction{Id: 3,
			Oda_number: "123412341234",
			Status:     "rejected",
			Price:      16000,
			Total_data: 12,
			Created_at: "2023-06-09 06:51:14",
		},
		Transaction{Id: 4,
			Oda_number: "123412341234",
			Status:     "rejected",
			Price:      19000,
			Total_data: 12,
			Created_at: "2023-06-09 05:51:14",
		},
	}
	//)
	switch req.Status {
	case "success":
		return dataSucces, nil

	case "pending":
		return dataPending, nil
	case "reject":
		return dataReject, nil
	}
	return nil, nil

}
