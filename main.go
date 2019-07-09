package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"log"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
  "github.com/davecgh/go-spew/spew"
	_ "time"
	_ "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

/* I'm just returning a HTTP 200 OK */
/* This is a request done by docker-cli before doing anything. */
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "OK")
}


/* This method is called by `docker ps` */
/* It will get all the pods running in the Kubernetes namespace, then craft an imitation response of the Docker daemon API */ 
func containerHandler(w http.ResponseWriter, r *http.Request, kubeconfig *string) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("staging").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// We'll initiate an empty slice of Containers
	var containers []Container
	for _, pod := range pods.Items {
		// For debugging purposes
		spew.Dump(pod)
		// A pod is a collection of containers. Therefore, let's build a slice of the images used by each containers
		var images []string
		for _, c := range pod.Spec.Containers {
			images = append(images, c.Image)
		}
		container := Container {
				ID: string(pod.UID),
				Image: images[0],
				ImageID: images[0],
				Names : []string { pod.Name },
				Status: string(pod.Status.Phase),
				Created: int(pod.CreationTimestamp.Unix()),
		}
		// This declaration is more convenient for nested structures
		container.NetworkSettings.Networks.Bridge.IPAddress = pod.Status.PodIP
		containers = append(containers, container)
	}

	// We now have a nice JSON containing our response.
	js, err := json.Marshal(containers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	  }
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

/* Used for getting the kubeconfig file */
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}


func main() {
		var kubeconfig *string
		if home := homeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		// Our API proxy will implement a few methods to imitate the Docker Daemon API
		// This one is for listing all containers
		http.HandleFunc("/v1.24/containers/json", func(w http.ResponseWriter, r *http.Request) {
			containerHandler(w, r, kubeconfig)
		})

		// This one is to check weither our "Docker Daemon" is alive
		http.HandleFunc("/_ping", func(w http.ResponseWriter, r *http.Request) {
			pingHandler(w, r)
		})

		// Catchall for all other undefined routes.
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			dump, err := httputil.DumpRequest(r, true)
			if err != nil {
				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
				return
			}
			// Let's dump the HTTP request in the console
			// Easier to reverse engineer which HTTP routes are called by the Docker CLI.
			spew.Dump(dump)
		})

    log.Fatal(http.ListenAndServe(":3000", nil))
}