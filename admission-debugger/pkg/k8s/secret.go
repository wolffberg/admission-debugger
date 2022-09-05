package k8s

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var clientset *kubernetes.Clientset

func newClientSet() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	newClientSet()
}

func NewSecret(data map[string]string) error {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: "admission-debugger-tls",
		},
		Type:       "kubernetes.io/tls",
		StringData: data,
	}

	_, err := clientset.CoreV1().Secrets("").Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("Error creating secret: %v", err)
	}
	//secret, err := clientset.CoreV1().Secrets("").Get(context.TODO(), s, metav1.GetOptions{})
	//if k8sErrors.IsNotFound(err) {
	//	return nil, errors.New("Secret not found in default namespace\n")
	//} else if statusError, isStatus := err.(*k8sErrors.StatusError); isStatus {
	//	return nil, errors.New(fmt.Sprintf("Error getting secret %v\n", statusError.ErrStatus.Message))
	//} else if err != nil {
	//	panic(err.Error())
	//}

	return nil
}
