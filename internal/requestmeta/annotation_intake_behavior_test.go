package requestmeta

import (
	"context"
	"testing"

	"github.com/zhanglei10281852-gif/ai/internal/domain"
)

func TestRequestPrincipalsRemainIndependent(t *testing.T) {
	base := context.Background()
	ml := domain.Principal{UserID: "ml-request", Role: domain.RoleMLEngineer}
	reviewer := domain.Principal{UserID: "review-request", Role: domain.RoleRiskReviewer}
	mlContext := WithPrincipal(base, ml)
	reviewerContext := WithPrincipal(base, reviewer)

	gotML, ok := Principal(mlContext)
	if !ok || gotML.UserID != ml.UserID || gotML.Role != ml.Role {
		t.Fatalf("first request principal = %+v, ok=%v", gotML, ok)
	}
	gotReviewer, ok := Principal(reviewerContext)
	if !ok || gotReviewer.UserID != reviewer.UserID || gotReviewer.Role != reviewer.Role {
		t.Fatalf("second request principal = %+v, ok=%v", gotReviewer, ok)
	}
}
