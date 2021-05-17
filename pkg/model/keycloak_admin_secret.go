package model

import (
	"os"

	"github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func KeycloakAdminSecret(cr *v1alpha1.Keycloak) *v1.Secret {
	return &v1.Secret{
		ObjectMeta: v12.ObjectMeta{
			Name:      "credential-" + cr.Name,
			Namespace: cr.Namespace,
			Labels: map[string]string{
				"app":           ApplicationName,
				ApplicationName: cr.Name,
			},
		},
		Data: map[string][]byte{
			AdminUsernameProperty: []byte("admin"),
			AdminPasswordProperty: []byte(GetPassword()),
		},
		Type: "Opaque",
	}
}

func KeycloakAdminSecretSelector(cr *v1alpha1.Keycloak) client.ObjectKey {
	return client.ObjectKey{
		Name:      "credential-" + cr.Name,
		Namespace: cr.Namespace,
	}
}

func KeycloakAdminSecretReconciled(cr *v1alpha1.Keycloak, currentState *v1.Secret) *v1.Secret {
	reconciled := currentState.DeepCopy()
	if val, ok := reconciled.Data[AdminUsernameProperty]; !ok || len(val) == 0 {
		reconciled.Data[AdminUsernameProperty] = []byte("admin")
	}
	adminPassword, ok := reconciled.Data[AdminPasswordProperty]
	envAdminPassword := []byte(GetEnvPassword())
	if !ok || len(adminPassword) == 0 {
		reconciled.Data[AdminPasswordProperty] = []byte(GetPassword())
	} else if len(envAdminPassword) != 0 && string(adminPassword) != string(envAdminPassword) {
		reconciled.Data[AdminPasswordProperty] = envAdminPassword
	}
	return reconciled
}

func GetEnvPassword() string {
	return os.Getenv("KEYCLOAK_PASSWORD")
}

func GetPassword() string {
	password := GetEnvPassword()

	if password == "" {
		return GenerateRandomString(10)
	}
	return password
}
