package order_service

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/OybekAbduvosiqov/tovar/models"
	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

var s int64

func TestProduct(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createProduct(t)
			updateProduct(t, id)
			deleteProduct(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createProduct(t *testing.T) string {
	response := &models.Product{}

	request := &models.CreateProduct{
		Name:        faker.Name(),
		Price:       float64(faker.RandomUnixTime()),
		Description: faker.Paragraph(),
		Photo:       faker.Paragraph(),
		CategoryId:  "",
	}

	resp, err := PerformRequest(http.MethodPost, "/product", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Id
}

func updateProduct(t *testing.T, id string) string {
	response := &models.Product{}
	request := &models.UpdateProduct{
		Id:          id,
		Name:        faker.Name(),
		Price:       float64(faker.RandomUnixTime()),
		Description: faker.Paragraph(),
		Photo:       faker.Paragraph(),
		CategoryId:  id,
	}

	resp, err := PerformRequest(http.MethodPut, "/product/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.Id
}

func deleteProduct(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/product/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
