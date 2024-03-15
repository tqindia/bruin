package helpers

import (
	"github.com/bruin-data/bruin/pkg/pipeline"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetIngestrDestinationType(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name    string
		asset   *pipeline.Asset
		want    pipeline.AssetType
		wantErr bool
	}{
		{
			name: "postgres",
			asset: &pipeline.Asset{
				Parameters: map[string]string{
					"destination": "postgres",
				},
			},
			want: pipeline.AssetTypePostgresQuery,
		},
		{
			name: "gcp",
			asset: &pipeline.Asset{
				Parameters: map[string]string{
					"destination": "postgres",
				},
			},
			want: pipeline.AssetTypeBigqueryQuery,
		},
		{
			name: "not found",
			asset: &pipeline.Asset{
				Parameters: map[string]string{
					"destination": "sqlite",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assetType, error := GetIngestrDestinationType(tc.asset)
			if tc.wantErr {
				require.Error(t, error)
			} else {
				assert.Equal(t, tc.want, assetType)
			}
		})
	}
}
