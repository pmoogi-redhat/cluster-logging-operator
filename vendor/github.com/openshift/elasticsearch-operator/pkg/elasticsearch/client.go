package elasticsearch

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	api "github.com/openshift/elasticsearch-operator/pkg/apis/logging/v1"
	estypes "github.com/openshift/elasticsearch-operator/pkg/types/elasticsearch"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	certLocalPath = "/tmp/"
	k8sTokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
)

type Client interface {
	// Cluster Settings API
	GetClusterNodeVersions() ([]string, error)
	GetThresholdEnabled() (bool, error)
	GetDiskWatermarks() (interface{}, interface{}, error)
	GetMinMasterNodes() (int32, error)
	SetMinMasterNodes(numberMasters int32) (bool, error)
	DoSynchronizedFlush() (bool, error)

	// Cluster State API
	GetLowestClusterVersion() (string, error)
	IsNodeInCluster(nodeName string) (bool, error)

	// Health API
	GetClusterHealth() (api.ClusterHealth, error)
	GetClusterHealthStatus() (string, error)
	GetClusterNodeCount() (int32, error)

	// Index API
	GetIndex(name string) (*estypes.Index, error)
	CreateIndex(name string, index *estypes.Index) error
	ReIndex(src, dst, script, lang string) error
	GetAllIndices(name string) (estypes.CatIndicesResponses, error)

	// Index Alias API
	ListIndicesForAlias(aliasPattern string) ([]string, error)
	UpdateAlias(actions estypes.AliasActions) error
	AddAliasForOldIndices() bool

	// Index Settings API
	GetIndexSettings(name string) (*estypes.IndexSettings, error)
	UpdateIndexSettings(name string, settings *estypes.IndexSettings) error

	// Nodes API
	GetNodeDiskUsage(nodeName string) (string, float64, error)

	// Replicas
	UpdateReplicaCount(replicaCount int32) error
	GetIndexReplicaCounts() (map[string]interface{}, error)

	// Shards API
	ClearTransientShardAllocation() (bool, error)
	GetShardAllocation() (string, error)
	SetShardAllocation(state api.ShardAllocationState) (bool, error)

	// Index Templates API
	CreateIndexTemplate(name string, template *estypes.IndexTemplate) error
	DeleteIndexTemplate(name string) error
	ListTemplates() (sets.String, error)
	GetIndexTemplates() (map[string]interface{}, error)

	SetSendRequestFn(fn FnEsSendRequest)
}

type FnEsSendRequest func(cluster, namespace string, payload *EsRequest, client k8sclient.Client)

type esClient struct {
	cluster         string
	namespace       string
	k8sClient       k8sclient.Client
	fnSendEsRequest FnEsSendRequest
}

type EsRequest struct {
	Method          string // use net/http constants https://golang.org/pkg/net/http/#pkg-constants
	URI             string
	RequestBody     string
	StatusCode      int
	RawResponseBody string
	ResponseBody    map[string]interface{}
	Error           error
}

func NewClient(cluster, namespace string, client k8sclient.Client) Client {
	return &esClient{
		cluster:         cluster,
		namespace:       namespace,
		k8sClient:       client,
		fnSendEsRequest: sendEsRequest,
	}
}

func (ec *esClient) SetSendRequestFn(fn FnEsSendRequest) {
	ec.fnSendEsRequest = fn
}

func sendEsRequest(cluster, namespace string, payload *EsRequest, client k8sclient.Client) {
	urlString := fmt.Sprintf("https://%s.%s.svc:9200/%s", cluster, namespace, payload.URI)
	urlURL, err := url.Parse(urlString)

	if err != nil {
		logrus.Warnf("Unable to parse URL %v: %v", urlString, err)
		return
	}

	request := &http.Request{
		Method: payload.Method,
		URL:    urlURL,
	}

	switch payload.Method {
	case http.MethodGet:
		// no more to do to request...
	case http.MethodPost:
		if payload.RequestBody != "" {
			// add to the request
			request.Header = map[string][]string{
				"Content-Type": {
					"application/json",
				},
			}
			request.Body = ioutil.NopCloser(bytes.NewReader([]byte(payload.RequestBody)))
		}

	case http.MethodPut:
		if payload.RequestBody != "" {
			// add to the request
			request.Header = map[string][]string{
				"Content-Type": {
					"application/json",
				},
			}
			request.Body = ioutil.NopCloser(bytes.NewReader([]byte(payload.RequestBody)))
		}

	default:
		// unsupported method -- do nothing
		return
	}

	request.Header = ensureTokenHeader(request.Header)
	httpClient := getClient(cluster, namespace, client)
	resp, err := httpClient.Do(request)

	if resp != nil {
		// TODO: eventually remove after all ES images have been updated to use SA token auth for EO?
		if resp.StatusCode == http.StatusForbidden ||
			resp.StatusCode == http.StatusUnauthorized {
			// if we get a 401 that means that we couldn't read from the token and provided
			// no header.
			// if we get a 403 that means the ES cluster doesn't allow us to use
			// our SA token.
			// in both cases, try the old way.

			// Not sure why, but just trying to reuse the request with the old client
			// resulted in a 400 every time. Doing it this way got a 200 response as expected.
			sendRequestWithOldClient(cluster, namespace, payload, client)
			return
		}

		payload.StatusCode = resp.StatusCode
		if payload.RawResponseBody, err = getRawBody(resp.Body); err != nil {
			logrus.Warnf("failed to get raw response body: %s", err)
		}
		if payload.ResponseBody, err = getMapFromBody(payload.RawResponseBody); err != nil {
			logrus.Warnf("getMapFromBody failed. E: %s\r\n", err.Error())
		}
	}

	payload.Error = err
}

func sendRequestWithOldClient(clusterName, namespace string, payload *EsRequest, client client.Client) {

	urlString := fmt.Sprintf("https://%s.%s.svc:9200/%s", clusterName, namespace, payload.URI)
	urlURL, err := url.Parse(urlString)

	if err != nil {
		logrus.Warnf("Unable to parse URL %v: %v", urlString, err)
		return
	}

	request := &http.Request{
		Method: payload.Method,
		URL:    urlURL,
	}

	switch payload.Method {
	case http.MethodGet:
		// no more to do to request...
	case http.MethodPost:
		if payload.RequestBody != "" {
			// add to the request
			request.Header = map[string][]string{
				"Content-Type": {
					"application/json",
				},
			}
			request.Body = ioutil.NopCloser(bytes.NewReader([]byte(payload.RequestBody)))
		}

	case http.MethodPut:
		if payload.RequestBody != "" {
			// add to the request
			request.Header = map[string][]string{
				"Content-Type": {
					"application/json",
				},
			}
			request.Body = ioutil.NopCloser(bytes.NewReader([]byte(payload.RequestBody)))
		}

	default:
		// unsupported method -- do nothing
		return
	}

	httpClient := getOldClient(clusterName, namespace, client)
	resp, err := httpClient.Do(request)

	if resp != nil {
		payload.StatusCode = resp.StatusCode
		if payload.RawResponseBody, err = getRawBody(resp.Body); err != nil {
			logrus.Warnf("failed to get raw response body: %s", err)
		}
		if payload.ResponseBody, err = getMapFromBody(payload.RawResponseBody); err != nil {
			logrus.Warnf("getMapFromBody failed. E: %s\r\n", err.Error())
		}
	}

	payload.Error = err
}

func ensureTokenHeader(header http.Header) http.Header {
	if header == nil {
		header = map[string][]string{}
	}

	if saToken, ok := readSAToken(k8sTokenFile); ok {
		header.Set("Authorization", fmt.Sprintf("Bearer %s", saToken))
	}

	return header
}

// we want to read each time so that we can be sure to have the most up to date
// token in the case where our perms change and a new token is mounted
func readSAToken(tokenFile string) (string, bool) {
	// read from /var/run/secrets/kubernetes.io/serviceaccount/token
	token, err := ioutil.ReadFile(tokenFile)

	if err != nil {
		logrus.Errorf("Unable to read auth token from file [%s]: %v", tokenFile, err)
		return "", false
	}

	if len(token) == 0 {
		logrus.Errorf("Unable to read auth token from file [%s]: empty token", tokenFile)
		return "", false
	}

	return string(token), true
}

func getClient(clusterName, namespace string, client client.Client) *http.Client {

	// get the contents of the secret
	extractSecret(clusterName, namespace, client)

	// http.Transport sourced from go 1.10.7
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
				RootCAs:            getRootCA(clusterName, namespace),
				Certificates:       getClientCertificates(clusterName, namespace),
			},
		},
	}
}

func getOldClient(clusterName, namespace string, client client.Client) *http.Client {

	// get the contents of the secret
	extractSecret(clusterName, namespace, client)

	// http.Transport sourced from go 1.10.7
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
				RootCAs:            getRootCA(clusterName, namespace),
				Certificates:       getClientCertificates(clusterName, namespace),
			},
		},
	}
}

func getRootCA(clusterName, namespace string) *x509.CertPool {
	certPool := x509.NewCertPool()

	// load cert into []byte
	caPem, err := ioutil.ReadFile(path.Join(certLocalPath, clusterName, "admin-ca"))
	if err != nil {
		logrus.Errorf("Unable to read file to get contents: %v", err)
		return nil
	}

	certPool.AppendCertsFromPEM(caPem)

	return certPool
}

func getClientCertificates(clusterName, namespace string) []tls.Certificate {
	certificate, err := tls.LoadX509KeyPair(
		path.Join(certLocalPath, clusterName, "admin-cert"),
		path.Join(certLocalPath, clusterName, "admin-key"),
	)
	if err != nil {
		return []tls.Certificate{}
	}

	return []tls.Certificate{
		certificate,
	}
}

func getRawBody(body io.ReadCloser) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(body); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func getMapFromBody(rawBody string) (map[string]interface{}, error) {
	if rawBody == "" {
		return make(map[string]interface{}), nil
	}
	var results map[string]interface{}
	err := json.Unmarshal([]byte(rawBody), &results)
	if err != nil {
		results = make(map[string]interface{})
		results["results"] = rawBody
	}

	return results, nil
}

func extractSecret(secretName, namespace string, client client.Client) {
	secret := &v1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: v1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: namespace,
		},
	}
	if err := client.Get(context.TODO(), types.NamespacedName{Name: secret.Name, Namespace: secret.Namespace}, secret); err != nil {
		if errors.IsNotFound(err) {
			//return err
			logrus.Errorf("Unable to find secret %v: %v", secretName, err)
		}

		logrus.Errorf("Error reading secret %v: %v", secretName, err)
		//return fmt.Errorf("Unable to extract secret to file: %v", secretName, err)
	}

	// make sure that the dir === secretName exists
	if _, err := os.Stat(path.Join(certLocalPath, secretName)); os.IsNotExist(err) {
		err = os.MkdirAll(path.Join(certLocalPath, secretName), 0755)
		if err != nil {
			logrus.Errorf("Error creating dir %v: %v", path.Join(certLocalPath, secretName), err)
		}
	}

	for _, key := range []string{"admin-ca", "admin-cert", "admin-key"} {

		value, ok := secret.Data[key]

		// check to see if the map value exists
		if !ok {
			logrus.Errorf("Error secret key %v not found", key)
			//return fmt.Errorf("No secret data \"%s\" found", key)
		}

		if err := ioutil.WriteFile(path.Join(certLocalPath, secretName, key), value, 0644); err != nil {
			//return fmt.Errorf("Unable to write to working dir: %v", err)
			logrus.Errorf("Error writing %v to %v: %v", value, path.Join(certLocalPath, secretName, key), err)
		}
	}
}
