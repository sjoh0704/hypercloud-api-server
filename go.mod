module github.com/tmax-cloud/hypercloud-api-server

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.1
	github.com/oklog/ulid v1.3.1
	github.com/robfig/cron v1.2.0
	github.com/tmax-cloud/claim-operator v0.0.0-20210114141758-083187ba4fc3
	github.com/tmax-cloud/cluster-manager-operator v0.0.0-20210112070657-53ca50ba8160
	github.com/tmax-cloud/efk-operator v0.0.0-20201207030412-fd9c02a3e1c2
	github.com/tmax-cloud/hypercloud-go-operator v0.0.0-20201125074013-0e686fd12999
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.19.4
	k8s.io/apimachinery v0.19.4
	k8s.io/apiserver v0.19.4
	k8s.io/client-go v0.19.4
	k8s.io/klog v1.0.0
	k8s.io/kubectl v0.19.3
	k8s.io/utils v0.0.0-20201005171033-6301aaf42dc7
	sigs.k8s.io/controller-runtime v0.7.0

)
