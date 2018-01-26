package setup

import (
	"github.com/rancher/types/config"
	// "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AddDate(logCtx config.LoggingContext) error {
	// initClusterLogging := v3.Logging{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name: "clusterLogging",
	// 		Labels: map[string]string{
	// 			"cluster": "logging",
	// 		},
	// 	},
	// 	Spec: v3.LoggingSpec{},
	// }

	_, err := logCtx.Logging.ProjectLoggings("").List(metav1.ListOptions{})
	// Create(&initClusterLogging)
	return err
}
