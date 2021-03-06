/*Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	_ "time"

	"k8s.io/client-go/1.5/kubernetes"
	"k8s.io/client-go/1.5/pkg/api"
	"k8s.io/client-go/1.5/pkg/api/v1"
	"k8s.io/client-go/1.5/pkg/api/unversioned"
	// "k8s.io/client-go/1.5/pkg/apis/batch"
	"k8s.io/client-go/1.5/tools/clientcmd"
	v2alpha1batch "k8s.io/client-go/1.5/pkg/apis/batch/v2alpha1"
	"k8s.io/client-go/1.5/kubernetes/typed/batch/v2alpha1"
)

var (
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	task = flag.String("task", "", "Data collection task")
	jobName = flag.String("job_name", "", "Data collection task name")
)

func getJobSpec(jobName string, task string) *v2alpha1batch.JobSpec {
	args := []string{
		"--task","eyJzeXN0ZW1fY29udGV4dCI6ICJ7XCJzaW5rX2luZm9cIjoge1widHlwZVwiOiBcInNwbHVua1wiLCBcImNvbnRleHRcIjogW3tcImhlY191cmlcIjogXCJodHRwczovLzEwLjY2LjQuMTc3OjgwODhcIiwgXCJoZWNfdG9rZW5cIjogXCJBRjlGOEQ2MS1ERDcxLTQ1RUEtQTNEQy0xQzBFMkJEQTREQUZcIn1dfSxcInRhc2tfc2VydmljZV91cmlcIjogXCJodHRwOi8vMTAuNjYuNC4xNzc6ODA4MC9hcGkvbGF0ZXN0L3Rhc2tzXCIsXCJzdGF0ZV9zZXJ2aWNlX3VyaVwiOiBcImh0dHA6Ly81MTAuNjYuNC4xNzc6ODA4MC9hcGkvbGF0ZXN0L3N0YXRlc1wiLFwicGVyZl9zZXJ2aWNlX3VyaVwiOiBcImh0dHA6Ly8xMC42Ni40LjE3Nzo4MDgwL2FwaS9sYXRlc3Qvc3RhdGVzL3BlcmZcIn0iLCAiam9iX2lkIjogImpvYl9mNzBkNzIwMi05MGYxLTRkNzUtYmY5Yy02ODY3N2Y1YzI0ZGIiLCAiaWQiOiAidGFza19mNzBkNzIwMi05MGYxLTRkNzUtYmY5Yy02ODY3N2Y1YzI0ZGYiLCAiY29udGV4dCI6IHsiYWN0aW9uIjogIkNPTExFQ1QiLCAidHlwZSI6ICJhd3MuczMiLCAic2NoZWQiOiAiNjAwIiwgImNvbnRleHQiOiB7InNvdXJjZXR5cGUiOiAiYXdzOnMzIiwgImtleXMiOiBbIkFXU0xvZ3MvMDYzNjA1NzE1MjgwL0Nsb3VkVHJhaWwvYXAtc291dGhlYXN0LTEvMjAxNi8wMS8xNy8wNjM2MDU3MTUyODBfQ2xvdWRUcmFpbF9hcC1zb3V0aGVhc3QtMV8yMDE2MDExN1QwMjI1Wl9nVEd3STRZZnNSa3g2UEwyLmpzb24uZ3oiLCAiQVdTTG9ncy8wNjM2MDU3MTUyODAvQ2xvdWRUcmFpbC9hcC1zb3V0aGVhc3QtMS8yMDE2LzAxLzE3LzA2MzYwNTcxNTI4MF9DbG91ZFRyYWlsX2FwLXNvdXRoZWFzdC0xXzIwMTYwMTE3VDAyMzBaX1h1VjVOeTJGd05Sd25yM3ouanNvbi5neiIsICJBV1NMb2dzLzA2MzYwNTcxNTI4MC9DbG91ZFRyYWlsL2FwLXNvdXRoZWFzdC0xLzIwMTYvMDEvMTcvMDYzNjA1NzE1MjgwX0Nsb3VkVHJhaWxfYXAtc291dGhlYXN0LTFfMjAxNjAxMTdUMDIzNVpfYW5Rc2llSDJZc1FFV2RPWC5qc29uLmd6Il0sICJyZWdpb24iOiAiYXAtc291dGhlYXN0LTEiLCAiYnVja2V0IjogInNhYXNhcHBxYS1jbG91ZHRyYWlsLXMzLXNpbmdhcG9yZS1yemhhbmciLCAic2VjcmV0X2lkIjogIkFLSUFKUkwzWVM0M0QyR0RLTUJRIiwgInNlY3JldF9rZXkiOiAiaDZIMVlOZy9Vejc4S2pud1ZLN2RrdVhnVkhRZDVSRWVyVDBCY0pBZiJ9fSwgIm5hbWUiOiAiczNfZGF0YWlucHV0In0="}

	podSpec := v1.PodSpec {
		Containers: []v1.Container {
			v1.Container{
				Name: jobName,
				Image: "splunk/python/aws_s3:v3",
				Args: args,
			},
		},
		RestartPolicy: v1.RestartPolicyOnFailure,
	}

	template := v1.PodTemplateSpec {
		ObjectMeta: v1.ObjectMeta {
			Name: jobName,
			Namespace: "default",
		},
		Spec: podSpec,
	}

	var p int32 = 1
	var duration int64 = 30
	spec := v2alpha1batch.JobSpec {
		Parallelism: &p,
		Completions: &p,
		ActiveDeadlineSeconds: &duration,
		Template: template,
	}
	return &spec
}

func createScheduledJob(jobs v2alpha1.ScheduledJobInterface, jobName string, task string) {
	template := v2alpha1batch.JobTemplateSpec {
		Spec: *getJobSpec(jobName, task),
	}

	var startingDeadline int64 = 30
	spec := v2alpha1batch.ScheduledJobSpec {
		Schedule: "0/1 * * * ?",
		StartingDeadlineSeconds: &startingDeadline,
		ConcurrencyPolicy: v2alpha1batch.ForbidConcurrent,
		JobTemplate: template,
	}

	job := v2alpha1batch.ScheduledJob{
		TypeMeta: unversioned.TypeMeta {
			Kind: "ScheduledJob",
			APIVersion: "batch/v2alpha1",
		},
		ObjectMeta: v1.ObjectMeta {
			Name: jobName,
			Namespace: "default",
		},
		Spec: spec,
	}

	result, err := jobs.Create(&job)
	if err != nil {
		fmt.Printf("Failed to create job error=%s\n", err)
	} else {
		fmt.Printf("Successfully create job %s\n", result)
	}
}

func createJob(jobs v2alpha1.JobInterface, jobName string, task string) {
	if jobName == "" {
		fmt.Println("Missing jobName")
	}

	job := v2alpha1batch.Job{
		TypeMeta: unversioned.TypeMeta {
			Kind: "Job",
			APIVersion: "batch/v2alpha1",
		},
		ObjectMeta: v1.ObjectMeta {
			Name: jobName,
			Namespace: "default",
		},
		Spec: *getJobSpec(jobName, task),
	}

	result, err := jobs.Create(&job)
	if err != nil {
		fmt.Printf("Failed to create job error=%s\n", err)
	} else {
		fmt.Printf("Successfully create job %s\n", result)
	}
}

func main() {
	flag.Parse()
	// uses the current context in kubeconfig
	configfile := *kubeconfig
	if configfile == "" {
		configfile = "/Users/Ken/.kube/config"
	}

	config, err := clientcmd.BuildConfigFromFlags("", configfile)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.Core().Pods("").List(api.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	jobs := clientset.Batch().Jobs("default")

	res, err := jobs.List(api.ListOptions{})
	fmt.Printf("There are %d jobs in the cluster\n", len(res.Items))
	createJob(jobs, *jobName, *task)

	res, err = jobs.List(api.ListOptions{})
	fmt.Printf("There are %d jobs in the cluster\n", len(res.Items))

	/*err = jobs.Delete(*jobName, &api.DeleteOptions{})
	if err != nil {
		fmt.Printf("Failed to delete job=%s, error=%s\n", *jobName, err)
	}*/

	res, err = jobs.List(api.ListOptions{})
	fmt.Printf("There are %d jobs in the cluster\n", len(res.Items))

	_, err = jobs.Get(*jobName)
	if err != nil {
		fmt.Printf("Failed to get job=%s, error=%s\n", *jobName, err)
	}

	scheduled := clientset.Batch().ScheduledJobs("default")
	sres, err := scheduled.List(api.ListOptions{})
	// fmt.Printf("%+v, %s\n", res, err)
	fmt.Printf("There are %d scheduled jobs in the cluster\n", len(sres.Items))

	createScheduledJob(scheduled, *jobName, *task)
}
