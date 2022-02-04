// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

package repository

import (
	"context"
	"fmt"
	promApiV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	promModel "github.com/prometheus/common/model"
	"time"
)

// AnalyticsRepository -
type AnalyticsRepository struct {
	promAPIV1 promApiV1.API
}

// NewAnalyticsRepository creates a new analytics repository
func NewAnalyticsRepository(promV1 promApiV1.API) AnalyticsRepository {
	return AnalyticsRepository{
		promAPIV1: promV1,
	}
}

// QueryMetrics executes a query
func (a *AnalyticsRepository) QueryMetrics(query string) (promModel.Value, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, warnings, err := a.promAPIV1.Query(ctx, query, time.Now())
	if err != nil {
		return nil, fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}

	// result is a Value, which is an interface to ValueType and a String() method
	// Can cast to:
	//    Matrix, Vector, *Scalar, *String
	return result, nil
}
