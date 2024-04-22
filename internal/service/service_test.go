package service_test

import (
	"testing"

	"github.com/GabriellGds/petshop-dti/internal/model"
	"github.com/GabriellGds/petshop-dti/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestListPetshops(t *testing.T) {
	testCases := []struct {
		name            string
		req             *model.Request
		expectedPrice   float64
		expectedPetshop string
		expectError     error
	}{
		{
			name: "success vai rex - weekend",
			req: &model.Request{
				Date:      "2024/04/21",
				SmallDogs: 25,
				BigDogs:   1,
			},
			expectedPrice: 555,
			expectedPetshop: "Vai Rex",
			expectError:   nil,
		},{
			name: "success vai rex chowChawgas - weekend",
			req: &model.Request{
				Date:      "2024/04/21",
				SmallDogs: 2,
				BigDogs:   20,
			},
			expectedPrice: 960,
			expectedPetshop: "ChowChawgas",
			expectError:   nil,
		},	
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			petshop, price, err := service.ListPetshops(tc.req)

			if tc.expectError != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedPetshop, petshop.Name)
				assert.Equal(t, tc.expectedPrice, price)
			}
		})
	}
}
