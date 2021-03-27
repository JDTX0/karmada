package detector

import (
	"github.com/karmada-io/karmada/pkg/util"
	"github.com/karmada-io/karmada/pkg/util/informermanager/keys"
)

// ClusterWideKeyFunc generates a ClusterWideKey for object.
func ClusterWideKeyFunc(obj interface{}) (util.QueueKey, error) {
	return keys.ClusterWideKeyFunc(obj)
}
