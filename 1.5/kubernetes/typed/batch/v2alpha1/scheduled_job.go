/*
Copyright 2016 The Kubernetes Authors.

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

package v2alpha1

import (
	api "k8s.io/client-go/1.5/pkg/api"
	"k8s.io/client-go/1.5/pkg/apis/batch/v2alpha1"
	watch "k8s.io/client-go/1.5/pkg/watch"
)

// JobsGetter has a method to return a ScheduledJobInterface.
// A group's client should implement this interface.
type ScheduledJobsGetter interface {
	ScheduledJobs(namespace string) ScheduledJobInterface
}

// ScheduledJobInterface has methods to work with Job resources.
type ScheduledJobInterface interface {
	Create(*v2alpha1.ScheduledJob) (*v2alpha1.ScheduledJob, error)
	Update(*v2alpha1.ScheduledJob) (*v2alpha1.ScheduledJob, error)
	UpdateStatus(*v2alpha1.ScheduledJob) (*v2alpha1.ScheduledJob, error)
	Delete(name string, options *api.DeleteOptions) error
	DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error
	Get(name string) (*v2alpha1.ScheduledJob, error)
	List(opts api.ListOptions) (*v2alpha1.ScheduledJobList, error)
	Watch(opts api.ListOptions) (watch.Interface, error)
	Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *v2alpha1.ScheduledJob, err error)
	JobExpansion
}

// scheduledJobs implements ScheduledJobInterface
type scheduledJobs struct {
	client *BatchClient
	ns     string
}

// newJobs returns a Jobs
func newScheduledJobs(c *BatchClient, namespace string) *scheduledJobs {
	return &scheduledJobs{
		client: c,
		ns:     namespace,
	}
}

// Create takes the representation of a job and creates it.  Returns the server's representation of the job, and an error, if there is any.
func (c *scheduledJobs) Create(job *v2alpha1.ScheduledJob) (result *v2alpha1.ScheduledJob, err error) {
	result = &v2alpha1.ScheduledJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("scheduledjobs").
		Body(job).
		Do().
		Into(result)
	return
}

// Update takes the representation of a job and updates it. Returns the server's representation of the job, and an error, if there is any.
func (c *scheduledJobs) Update(job *v2alpha1.ScheduledJob) (result *v2alpha1.ScheduledJob, err error) {
	result = &v2alpha1.ScheduledJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scheduledjobs").
		Name(job.Name).
		Body(job).
		Do().
		Into(result)
	return
}

func (c *scheduledJobs) UpdateStatus(job *v2alpha1.ScheduledJob) (result *v2alpha1.ScheduledJob, err error) {
	result = &v2alpha1.ScheduledJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scheduledjobs").
		Name(job.Name).
		SubResource("status").
		Body(job).
		Do().
		Into(result)
	return
}

// Delete takes name of the job and deletes it. Returns an error if one occurs.
func (c *scheduledJobs) Delete(name string, options *api.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scheduledjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scheduledJobs) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scheduledjobs").
		VersionedParams(&listOptions, api.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the job, and returns the corresponding job object, and an error if there is any.
func (c *scheduledJobs) Get(name string) (result *v2alpha1.ScheduledJob, err error) {
	result = &v2alpha1.ScheduledJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scheduledjobs").
		Name(name).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Jobs that match those selectors.
func (c *scheduledJobs) List(opts api.ListOptions) (result *v2alpha1.ScheduledJobList, err error) {
	result = &v2alpha1.ScheduledJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scheduledjobs").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested jobs.
func (c *scheduledJobs) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("scheduledjobs").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched job.
func (c *scheduledJobs) Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *v2alpha1.ScheduledJob, err error) {
	result = &v2alpha1.ScheduledJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("scheduledjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
