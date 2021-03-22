package model

import (
	kc "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func KeycloakRealmSecret(cr *kc.KeycloakRealm) *v1.Secret {
	return &v1.Secret{
		ObjectMeta: v12.ObjectMeta{
			Name:      "keycloak-" + cr.Name,
			Namespace: cr.Namespace,
			Labels: map[string]string{
				"app":           ApplicationName,
				ApplicationName: cr.Name,
			},
		},
		Data: map[string][]byte{
			// This will be overwritten in the runner because we do not have
			// access to the public key here
			RealmSecretPublicKey: []byte("dummy"),
		},
		Type: "Opaque",
	}
}

func KeycloakRealmSecretSelector(cr *kc.KeycloakRealm) client.ObjectKey {
	return client.ObjectKey{
		Name:      "keycloak-" + cr.Name,
		Namespace: cr.Namespace,
	}
}

func KeycloakRealmSecretReconciled(cr *kc.KeycloakRealm, currentState *v1.Secret) *v1.Secret {
	reconciled := currentState.DeepCopy()
	// Secret will not be updated here
	return reconciled
}
