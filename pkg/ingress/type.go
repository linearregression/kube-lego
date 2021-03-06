package ingress

import (
	"github.com/jetstack/kube-lego/pkg/kubelego_const"

	k8sExtensions "k8s.io/kubernetes/pkg/apis/extensions"
)

type Ingress struct {
	kubelego.Ingress

	IngressApi *k8sExtensions.Ingress
	exists     bool
	kubelego   kubelego.KubeLego
}

type Tls struct {
	*k8sExtensions.IngressTLS
	ingress *Ingress
}
