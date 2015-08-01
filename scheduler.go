package main

import (
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	mesos "github.com/mesos/mesos-go/mesosproto"
	sched "github.com/mesos/mesos-go/scheduler"
)

type Scheduler struct {
}

func newScheduler() *Scheduler {
	return &Scheduler{}
}

func (sched *Scheduler) Registered(driver sched.SchedulerDriver, frameworkID *mesos.FrameworkID, masterInfo *mesos.MasterInfo) {
	glog.Infof("registered with master %s", masterInfo.GetId())
}

func (sched *Scheduler) Reregistered(driver sched.SchedulerDriver, masterInfo *mesos.MasterInfo) {
	glog.Infof("re-registered with master %s", masterInfo.GetId())
}

func (sched *Scheduler) Disconnected(driver sched.SchedulerDriver) {
	glog.Warningf("disconnected from master")
}

func (sched *Scheduler) ResourceOffers(driver sched.SchedulerDriver, offers []*mesos.Offer) {
	glog.Infof("received %d offer(s)", len(offers))
	for _, offer := range offers {
		driver.DeclineOffer(offer.GetId(), refuseFor(10*time.Second))
	}
}

func (sched *Scheduler) StatusUpdate(driver sched.SchedulerDriver, status *mesos.TaskStatus) {
	if glog.V(1) {
		glog.Infof("status update from task %s in state %s under executor %s on slave %s: %s",
			status.GetTaskId().GetValue(),
			status.GetState(),
			status.GetExecutorId().GetValue(),
			status.GetSlaveId().GetValue(),
			status.GetMessage(),
		)
	}
}

func (sched *Scheduler) OfferRescinded(driver sched.SchedulerDriver, offerID *mesos.OfferID) {
	glog.Infof("offer %s has been recinded before we could use it", offerID.GetValue())
}

func (sched *Scheduler) FrameworkMessage(driver sched.SchedulerDriver, executorID *mesos.ExecutorID, slaveID *mesos.SlaveID, data string) {
	glog.Errorf("got framework message from executor %s on slave %s: %q", executorID.GetValue(), slaveID.GetValue(), data)
}

func (sched *Scheduler) SlaveLost(driver sched.SchedulerDriver, slaveID *mesos.SlaveID) {
	glog.Errorf("lost slave %s", slaveID.GetValue())
}

func (sched *Scheduler) ExecutorLost(driver sched.SchedulerDriver, executorID *mesos.ExecutorID, slaveID *mesos.SlaveID, status int) {
	glog.Errorf("lost executor %s on slave %s with status %d", executorID.GetValue(), slaveID.GetValue(), status)
}

func (sched *Scheduler) Error(driver sched.SchedulerDriver, message string) {
	glog.Errorf("unrecoverable error in scheduler or driver: %s", message)
}

func refuseFor(d time.Duration) *mesos.Filters {
	return &mesos.Filters{RefuseSeconds: proto.Float64(d.Seconds())}
}

func runScheduler() error {
	config := sched.DriverConfig{
		Scheduler: newScheduler(),
		Framework: &mesos.FrameworkInfo{
			User: proto.String(""),
			Name: proto.String("Elevators"),
		},
		Master: "127.0.0.1:5050",
	}
	driver, err := sched.NewMesosSchedulerDriver(config)
	if err != nil {
		return err
	}
	status, err := driver.Run()
	if err != nil {
		return fmt.Errorf("framework stopped with %s: %s", status, err)
	}
	return nil
}
