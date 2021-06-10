package cluster

import (
	"net/http"

	gmux "github.com/gorilla/mux"
	util "github.com/tmax-cloud/hypercloud-api-server/util"
	caller "github.com/tmax-cloud/hypercloud-api-server/util/caller"
	clusterDataFactory "github.com/tmax-cloud/hypercloud-api-server/util/dataFactory/cluster"
	"k8s.io/klog"
	// "encoding/json"
)

const (
	// QUERY_PARAMETER_USER_NAME   = "Name"
	QUERY_PARAMETER_USER_ID     = "userId"
	QUERY_PARAMETER_USER_NAME   = "userName"
	QUERY_PARAMETER_LIMIT       = "limit"
	QUERY_PARAMETER_OFFSET      = "offset"
	QUERY_PARAMETER_CLUSTER     = "cluster"
	QUERY_PARAMETER_ACCESS_ONLY = "accessOnly"
	QUERY_PARAMETER_REMOTE_ROLE = "remoteRole"
	QUERY_PARAMETER_MEMBER_NAME = "memberName"
)

func ListPage(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	userId := queryParams.Get(QUERY_PARAMETER_USER_ID)
	userGroups := queryParams[util.QUERY_PARAMETER_USER_GROUP]
	// accessOnly := queryParams.Get(QUERY_PARAMETER_ACCESS_ONLY)
	vars := gmux.Vars(req)
	namespace := vars["namespace"]

	if err := util.StringParameterException(userGroups, userId); err != nil {
		klog.Errorln(err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	if namespace == "" {
		clmList, msg, status := caller.ListAllCluster(userId, userGroups)
		util.SetResponse(res, msg, clmList, status)
		return
	} else {
		clusterClaimList, msg, status := caller.ListClusterInNamespace(userId, userGroups, namespace)
		util.SetResponse(res, msg, clusterClaimList, status)
		return
	}
}

func ListLNB(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	userId := queryParams.Get(QUERY_PARAMETER_USER_ID)
	userGroups := queryParams[util.QUERY_PARAMETER_USER_GROUP]

	if err := util.StringParameterException(userGroups, userId); err != nil {
		klog.Errorln(err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	clmList, msg, status := caller.ListAccesibleCluster(userId, userGroups)
	util.SetResponse(res, msg, clmList, status)
	return

}

func DeleteCLM(res http.ResponseWriter, req *http.Request) {
	// queryParams := req.URL.Query()
	vars := gmux.Vars(req)
	namespace := vars["namespace"]
	clustermanager := vars["clustermanager"]

	if err := util.StringParameterException([]string{}, namespace, clustermanager); err != nil {
		klog.Errorln(err)
		util.SetResponse(res, err.Error(), nil, http.StatusBadRequest)
		return
	}

	if err := clusterDataFactory.DeleteALL(namespace, clustermanager); err != nil {
		msg := "Failed to delete cluster info from db"
		klog.Infoln(msg)
		util.SetResponse(res, msg, nil, http.StatusInternalServerError)
		return
	}
	msg := "Success to delete cluster info from db"
	klog.Infoln(msg)
	util.SetResponse(res, msg, nil, http.StatusOK)
	return
}
