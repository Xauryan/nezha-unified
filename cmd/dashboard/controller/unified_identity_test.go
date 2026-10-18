package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/gin-gonic/gin"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/service/singleton"
	"github.com/nezhahq/nezha/service/unified"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUnifiedScopedBearerRoutingRejectsLegacyAndManagementEscalation(t *testing.T) {
	dir := t.TempDir()
	engine, err := embedded.Open(filepath.Join(dir, "certimate"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = engine.Close() })
	db, err := gorm.Open(sqlite.Open(filepath.Join(dir, "nezha.db")), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = sqlDB.Close() })
	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatal(err)
	}
	user := model.User{Common: model.Common{ID: 1}, Username: "synthetic-admin", Role: model.RoleAdmin}
	if err := db.Create(&user).Error; err != nil {
		t.Fatal(err)
	}
	s, err := unified.New(db, engine)
	if err != nil {
		t.Fatal(err)
	}
	prior := singleton.Unified
	singleton.Unified = s
	t.Cleanup(func() { singleton.Unified = prior })
	admin := unified.Principal{UserID: 1, Admin: true}
	issued, err := s.CreateScopedToken(context.Background(), admin, unified.ScopedTokenInput{Name: "synthetic-http-test", ExpiresAt: time.Now().Add(time.Hour), Scopes: []unified.Scope{{Kind: "certificate", ResourceID: "permitted", Action: "read"}}})
	if err != nil {
		t.Fatal(err)
	}
	r := gin.New()
	api := r.Group("/api/v1/unified-api", unifiedScopedTokenAuthMiddleware())
	bindUnifiedIdentityRoutes(api)
	api.GET("/certificates/:id", func(c *gin.Context) {
		p := unifiedPrincipal(c)
		if _, ok := c.Get(model.CtxKeyAPIToken); ok {
			t.Error("unified token gained a legacy PAT context")
		}
		if !s.Can(c.Request.Context(), p, "certificate", c.Param("id"), "read") {
			unifiedResult(c, nil, unified.ErrDenied)
			return
		}
		unifiedResult(c, gin.H{"id": c.Param("id")}, nil)
	})
	for _, tc := range []struct {
		name, path, bearer string
		want               int
	}{
		{"missing", "/certificates/permitted", "", http.StatusUnauthorized},
		{"legacy", "/certificates/permitted", "nzp_synthetic-legacy", http.StatusUnauthorized},
		{"valid", "/certificates/permitted", issued.Secret, http.StatusOK},
		{"scope", "/certificates/other", issued.Secret, http.StatusForbidden},
		{"tokens", "/tokens", issued.Secret, http.StatusForbidden},
		{"identities", "/identities", issued.Secret, http.StatusForbidden},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/unified-api"+tc.path, nil)
			req.AddCookie(&http.Cookie{Name: "nz_token", Value: "synthetic-cookie-cannot-authenticate"})
			if tc.bearer != "" {
				req.Header.Set("Authorization", "Bearer "+tc.bearer)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != tc.want {
				t.Fatalf("status = %d, want %d", w.Code, tc.want)
			}
			if strings.Contains(w.Body.String(), issued.Secret) {
				t.Fatal("bearer appeared in a response")
			}
			if w.Header().Get("Cache-Control") != "no-store" {
				t.Fatal("identity response missing no-store")
			}
		})
	}
	req := httptest.NewRequest(http.MethodGet, "/api/v1/unified-api/certificates/permitted", nil)
	req.Header.Add("Authorization", "Bearer "+issued.Secret)
	req.Header.Add("Authorization", "Bearer "+issued.Secret)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatal("ambiguous bearer headers accepted")
	}
	if err := s.RevokeScopedToken(context.Background(), admin, issued.Token.ID); err != nil {
		t.Fatal(err)
	}
	req = httptest.NewRequest(http.MethodGet, "/api/v1/unified-api/certificates/permitted", nil)
	req.Header.Set("Authorization", "Bearer "+issued.Secret)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatal("revoked token passed HTTP authentication")
	}
}
